package upload

import (
	"context"
	"errors"
	helper_os "github.com/langwan/langgo/helpers/os"
	helper_progress "github.com/langwan/langgo/helpers/progress"
	helper_string "github.com/langwan/langgo/helpers/string"
	"github.com/panjf2000/ants"
	"io"
	"math"
	"os"
)

type part struct {
	id          int    `json:"id"`
	offset      int64  `json:"offset"`
	size        int64  `json:"size"`
	isCompleted bool   `json:"is_completed"`
	etag        string `json:"etag"`
}

type PartList interface {
	LoadList() ([]*part, error)
	SetList(parts []*part)
	GetList() []*part
	Save() error
	GetUploadId() string
	SetUploadId(uploadId string) error
}

type invokeParams struct {
	ctx       context.Context
	completed chan *part
	failed    chan error
	dst       string
	part      *part
	listener  helper_progress.ProgressListener
	fileSize  int64
	uploadId  string
	writer    Writer
	srcFile   io.ReaderAt
}

type Writer interface {
	Create(key string) (string, error)
	UploadPart(key string, uploadId string, partId int64, data []byte) (string, error)
	Completed(key string, uploadId string, parts []*part) error
}

type Upload struct {
	Workers  int
	PartSize int64
}

var pool *ants.PoolWithFunc

func (up *Upload) Init() {
	pool, _ = ants.NewPoolWithFunc(up.Workers, func(i interface{}) {
		params := i.(*invokeParams)
		buf := make([]byte, params.part.size)
		params.srcFile.ReadAt(buf, params.part.offset)
		var err error
		params.part.etag, err = params.writer.UploadPart(params.dst, params.uploadId, int64(params.part.id), buf)
		if err != nil {
			params.failed <- err
			return
		} else {
			params.completed <- params.part
		}
	})
}

func (up *Upload) Upload(ctx context.Context, src string, dst string, partList PartList, writer Writer, listener helper_progress.ProgressListener) (err error) {

	if !helper_os.FileExists(src) {
		return errors.New("src file does not exist")
	}

	fi, err := helper_os.GetFileInfo(src)
	if err != nil {
		return err
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if helper_string.IsEmpty(partList.GetUploadId()) {
		parts, err := genParts(fi.Stat.Size(), up.PartSize)
		if err != nil {
			return err
		}
		uploadId, err := writer.Create(dst)
		if err != nil {
			return err
		}
		partList.SetUploadId(uploadId)
		partList.SetList(parts)
	}
	parts := partList.GetList()
	partCount := len(parts)
	chCompleted := make(chan *part, partCount)
	chFailed := make(chan error)
	completedCount := 0
	for _, part := range parts {
		if !part.isCompleted {
			pool.Invoke(&invokeParams{
				ctx:       ctx,
				completed: chCompleted,
				failed:    chFailed,
				part:      part,
				writer:    writer,
				srcFile:   srcFile,
				dst:       dst,
				uploadId:  partList.GetUploadId(),
			})
		} else {
			completedCount++
		}
	}
	var wm int64 = 0
	for completedCount < partCount {
		select {
		case rp := <-chCompleted:
			completedCount++
			rp.isCompleted = true
			partList.Save()
			wm += rp.size
			listener.ProgressChanged(&helper_progress.ProgressEvent{
				ConsumedBytes: wm,
				TotalBytes:    fi.Stat.Size(),
				RwBytes:       wm,
				EventType:     helper_progress.TransferDataEvent,
			})
		case err = <-chFailed:
			return err
		}

		if completedCount >= partCount {
			break
		}
	}
	writer.Completed(dst, partList.GetUploadId(), partList.GetList())
	return nil
}

func genParts(fileSize int64, partSize int64) (parts []*part, err error) {
	count := int(math.Ceil(float64(fileSize) / float64(partSize)))
	var offset int64 = 0
	remain := fileSize
	parts = make([]*part, count)
	for i := 0; i < count; i++ {
		ps := partSize
		if remain < partSize {
			ps = remain
		}
		remain -= partSize
		parts[i] = &part{
			id:          i + 1,
			offset:      offset,
			size:        ps,
			isCompleted: false,
		}
		offset += partSize
	}
	return parts, nil
}
