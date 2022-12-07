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
	"os"
	"path/filepath"
)

var pool *ants.PoolWithFunc

type Instance struct {
	Workers  int    `yaml:"workers"`
	PartSize string `yaml:"part_size"`
	partSize int64
	BufSize  string `yaml:"buf_size"`
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
	buf       []byte
	dst       string
	part      *part
	listener  helper_progress.ProgressListener
	fileSize  int64
	reader    FileReader
}

type FileReader interface {
	GetFileSize() (int64, error)
	OpenRange(offset, size int64) (io.ReadCloser, error)
}

func (d *Instance) Download(ctx context.Context, dst string, reader FileReader, listener helper_progress.ProgressListener) (err error) {
	db := dbPath(dst)
	var parts []*part
	var fileSize int64 = 0
	if helper_os.FileExists(db) {
		parts, err = loadDb(db)
	} else {
		fileSize, err = reader.GetFileSize()
		if err != nil {
			return err
		}
		parts, err = d.genParts(fileSize)
	}

	completedCount := 0
	for _, part := range parts {
		if part.isCompleted {
			completedCount++
		}
	}

	partCount := len(parts)
	completed := make(chan *part, partCount)
	failed := make(chan error)
	listener.ProgressChanged(&helper_progress.ProgressEvent{
		ConsumedBytes: 0,
		TotalBytes:    fileSize,
		RwBytes:       0,
		EventType:     helper_progress.TransferStartedEvent,
	})
	for _, part := range parts {
		if !part.isCompleted {
			pool.Invoke(&invokeParams{
				ctx:       ctx,
				completed: completed,
				failed:    failed,
				buf:       make([]byte, d.bufSize),
				dst:       dpPath(part.id, dst),
				part:      part,
				fileSize:  fileSize,
				reader:    reader,
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
	return merge(dst, parts, fileSize, listener)
}

func (d *Instance) Tune(size int) {
	pool.Tune(size)
}

func dbPath(dst string) string {
	dstDir := filepath.Dir(dst)
	dstBase := filepath.Base(dst)
	return filepath.Join(dstDir, fmt.Sprintf("%s%s", dstBase, suffixDb))
}

func merge(dst string, parts []*part, fileSize int64, listener helper_progress.ProgressListener) error {
	dstDir := filepath.Dir(dst)
	dstBase := filepath.Base(dst)
	newFilename, err := helper_os.GenUniqueFilename(dst, 10, nil)
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
	listener.ProgressChanged(&helper_progress.ProgressEvent{
		ConsumedBytes: fileSize,
		TotalBytes:    fileSize,
		RwBytes:       fileSize,
		EventType:     helper_progress.TransferCompletedEvent,
	})
	return nil
}
func (d *Instance) genParts(fileSize int64) (parts []*part, err error) {
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

	body, err := params.reader.OpenRange(params.part.offset, params.part.size)
	if err != nil {
		return 0, err
	}
	defer body.Close()
	for {
		select {
		case <-params.ctx.Done():
			params.failed <- err
			return wn, errors.New("context done")
		default:
		}
		var n int
		n, err = body.Read(params.buf)
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
