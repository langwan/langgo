package download

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/langwan/langgo/helpers/os"
	"github.com/langwan/langgo/helpers/progress"
	helper_size "github.com/langwan/langgo/helpers/size"
	"github.com/panjf2000/ants"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var pool *ants.PoolWithFunc

type Instance struct {
	Workers  int    `yaml:"workers"`
	PartSize string `yaml:"partSize"`
	partSize int64
	BufSize  string `yaml:"bufSize"`
	bufSize  int64
}

const name = "download"

var instance *Instance

func (d *Instance) Run() error {
	instance = d
	d.bufSize, _ = helper_size.RAMInBytes(d.BufSize)
	d.partSize, _ = helper_size.RAMInBytes(d.PartSize)
	pool, _ = ants.NewPoolWithFunc(instance.Workers, func(i interface{}) {
		params := i.(*invokeParams)
		downloadPartToFile(params)
	})
	return nil
}

func (d *Instance) GetName() string {
	return name
}

func Get() *Instance {
	return instance
}

const (
	suffixDb = ".db"
	suffixDp = ".dp"
)

type part struct {
	id          int   `json:"id"`
	offset      int64 `json:"offset"`
	size        int64 `json:"size"`
	isCompleted bool  `json:"is_completed"`
}

type invokeParams struct {
	ctx       context.Context
	completed chan *part
	failed    chan error
	url       string
	buf       []byte
	dst       string
	part      *part
	listener  helper_progress.ProgressListener
	fileSize  int64
}

func (d *Instance) Download(ctx context.Context, url, dst string, listener helper_progress.ProgressListener) (err error) {
	db := dbPath(dst)
	var parts []*part
	if helper_os.FileExists(db) {
		parts, err = loadDb(db)
	} else {
		parts, err = d.genParts(url)
	}
	var fileSize int64 = 0
	completedCount := 0
	for _, part := range parts {
		if part.isCompleted {
			completedCount++
		}
		fileSize += part.size
	}

	partCount := len(parts)
	completed := make(chan *part, partCount)
	failed := make(chan error)

	for _, part := range parts {
		if !part.isCompleted {
			pool.Invoke(&invokeParams{
				ctx:       ctx,
				completed: completed,
				failed:    failed,
				url:       url,
				buf:       make([]byte, d.bufSize),
				dst:       dpPath(part.id, dst),
				part:      part,
				fileSize:  fileSize,
			})
		}
	}
	var wm int64 = 0
	for completedCount < partCount {
		select {
		case rp := <-completed:
			completedCount++
			rp.isCompleted = true
			marshal, err := json.Marshal(parts)
			if err != nil {
				return err
			}
			os.WriteFile(db, marshal, os.ModePerm)
			wm += rp.size
			listener.ProgressChanged(&helper_progress.ProgressEvent{
				ConsumedBytes: wm,
				TotalBytes:    fileSize,
				RwBytes:       wm,
				EventType:     helper_progress.TransferDataEvent,
			})
		case err = <-failed:
			return err
		}

		if completedCount >= partCount {
			break
		}
	}
	return merge(dst, parts)
}

func (d *Instance) Tune(size int) {
	pool.Tune(size)
}

func dbPath(dst string) string {
	dstDir := filepath.Dir(dst)
	dstBase := filepath.Base(dst)
	return filepath.Join(dstDir, fmt.Sprintf("%s%s", dstBase, suffixDb))
}

func merge(dst string, parts []*part) error {
	dstDir := filepath.Dir(dst)
	dstBase := filepath.Base(dst)
	newFilename, err := helper_os.NewFilename(dst, 10, nil)
	if err != nil {
		return err
	}

	fs, err := os.Create(newFilename)
	defer fs.Close()
	if err != nil {
		return err
	}
	var offset int64 = 0
	for i := 0; i < len(parts); i++ {
		pf := filepath.Join(dstDir, fmt.Sprintf("%s%s%d", dstBase, suffixDp, i))
		buf, err := os.ReadFile(pf)
		if err != nil {
			return err
		}
		at, err := fs.WriteAt(buf, offset)
		if err != nil {
			return err
		}
		offset += int64(at)
	}

	os.Remove(dbPath(dst))
	for i := 0; i < len(parts); i++ {
		os.Remove(dpPath(i, dst))
	}

	return nil
}
func (d *Instance) genParts(url string) (parts []*part, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fileSize := resp.ContentLength
	if fileSize == -1 {
		return nil, errors.New("file size is error")
	}

	count := int(math.Ceil(float64(fileSize) / float64(d.partSize)))
	var offset int64 = 0
	remain := fileSize
	parts = make([]*part, count)
	for i := 0; i < count; i++ {
		ps := d.partSize
		if remain < d.partSize {
			ps = remain
		}
		remain -= d.partSize
		parts[i] = &part{
			id:          i,
			offset:      offset,
			size:        ps,
			isCompleted: false,
		}
		offset += d.partSize
	}
	return parts, nil
}

func dpPath(id int, dst string) string {
	dstDir := filepath.Dir(dst)
	dstBase := filepath.Base(dst)
	return filepath.Join(dstDir, fmt.Sprintf("%s%s%d", dstBase, suffixDp, id))
}

func loadDb(dbPath string) (parts []*part, err error) {
	data, err := os.ReadFile(dbPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &parts)
	if err != nil {
		return parts, err
	}
	return parts, nil
}

func downloadPartToFile(params *invokeParams) (int64, error) {
	fs, err := os.Create(params.dst)
	defer fs.Close()
	if err != nil {
		return 0, err
	}
	return downloadPartToWriter(fs, params)
}

func downloadPartToWriter(writer io.WriterAt, params *invokeParams) (int64, error) {
	var err error
	select {
	case <-params.ctx.Done():
		err = errors.New("context done")
		params.failed <- err
		return -1, err
	default:
	}

	if params.part.size < 1 {
		err = errors.New("part size error")
		params.failed <- err
		return -1, err
	}

	if params.buf == nil {
		params.buf = make([]byte, 1024*200)
	}

	var wn int64 = 0

	request, err := http.NewRequest("GET", params.url, nil)
	if err != nil {
		params.failed <- err
		return -1, err
	}

	request.Header.Set(
		"Range",
		"bytes="+strconv.FormatInt(params.part.offset, 10)+"-"+strconv.FormatInt(params.part.offset+params.part.size-1, 10),
	)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		params.failed <- err
		return -1, err
	}

	defer resp.Body.Close()
	for {
		select {
		case <-params.ctx.Done():
			params.failed <- err
			return wn, errors.New("context done")
		default:
		}
		var n int
		n, err = resp.Body.Read(params.buf)
		if err != nil && err != io.EOF {
			params.failed <- err
			return wn, err
		} else {

			writer.WriteAt(params.buf[:n], wn)
			wn += int64(n)
			if err == io.EOF {
				break
			}

		}
	}
	params.completed <- params.part
	return wn, nil
}
