package helper_http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	helper_os "github.com/langwan/langgo/helpers/os"
	"github.com/panjf2000/ants"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func DownloadPartToWriter(ctx context.Context, completed chan<- int, failed chan<- error, url string, buf []byte, writer io.WriterAt, part *DownloadFilePart) (int64, error) {
	var err error
	select {
	case <-ctx.Done():
		err = errors.New("context done")
		failed <- err
		return -1, err
	default:
	}

	if part.Size < 1 {
		err = errors.New("size < 0")
		failed <- err
		return -1, err
	}

	if buf == nil {
		buf = make([]byte, 1024*200)
	}

	var wn int64 = 0

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		failed <- err
		return -1, err
	}

	request.Header.Set(
		"Range",
		"bytes="+strconv.FormatInt(part.Start, 10)+"-"+strconv.FormatInt(part.Start+part.Size-1, 10),
	)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		failed <- err
		return -1, err
	}

	defer resp.Body.Close()
	for {
		select {
		case <-ctx.Done():
			failed <- err
			return wn, errors.New("context done")
		default:
		}
		var n int
		n, err = resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			failed <- err
			return wn, err
		} else {
			writer.WriteAt(buf[:n], wn)
			wn += int64(n)
			if err == io.EOF {
				break
			}
		}
	}
	completed <- part.Id
	return wn, nil
}

func DownloadPartToFile(ctx context.Context, completed chan<- int, failed chan<- error, url string, buf []byte, partName string, part *DownloadFilePart) (int64, error) {
	fmt.Println("DownloadPartToFile", partName, part.Start)
	fs, err := os.Create(partName)
	defer fs.Close()
	if err != nil {
		return 0, err
	}
	return DownloadPartToWriter(ctx, completed, failed, url, buf, fs, part)
}

type DownloadFilePartTask struct {
	Ctx       context.Context
	Completed chan int
	Failed    chan error
	Url       string
	Buf       []byte
	Dst       string
	Part      *DownloadFilePart
}

type DownloadFilePart struct {
	Id          int   `json:"id"`
	Start       int64 `json:"start"`
	Size        int64 `json:"size"`
	IsCompleted bool  `json:"is_completed"`
}

func DownloadFileMultiPart(ctx context.Context, url string, dst string, gp *ants.PoolWithFunc, partSize int64) error {
	dstDir := filepath.Dir(dst)
	dstBase := filepath.Base(dst)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fileSize := resp.ContentLength
	if fileSize == -1 {
		return errors.New("filesize -1")
	}
	partNumbers := int(math.Ceil(float64(fileSize) / float64(partSize)))
	var offset int64 = 0
	remain := fileSize

	completed := make(chan int, partNumbers)
	failed := make(chan error)
	db := filepath.Join(dstDir, fmt.Sprintf("%s.db", dstBase))

	var parts []*DownloadFilePart

	if helper_os.FileExists(db) {
		data, err := os.ReadFile(db)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, &parts)
		if err != nil {
			return err
		}
	} else {
		for i := 1; i < partNumbers+1; i++ {
			ps := partSize
			if remain < partSize {
				ps = remain
			}
			remain -= partSize
			//fmt.Printf("down %d %d\n", i, offset)
			parts = append(parts, &DownloadFilePart{
				Id:          i,
				Start:       offset,
				Size:        ps,
				IsCompleted: false,
			})
			offset += partSize
		}
	}
	completedNumbers := 0
	for _, part := range parts {
		if part.IsCompleted == true {
			completedNumbers++
		}
	}

	for _, part := range parts {
		if !part.IsCompleted {
			gp.Invoke(&DownloadFilePartTask{
				Ctx:       ctx,
				Completed: completed,
				Failed:    failed,
				Url:       url,
				Buf:       nil,
				Dst:       filepath.Join(dstDir, fmt.Sprintf("%s.dp%d", dstBase, part.Id)),
				Part:      part,
			})
		}
	}

	for completedNumbers < partNumbers {
		select {
		case partId := <-completed:
			completedNumbers++
			fmt.Println("completedNumbers", completedNumbers)
			for _, part := range parts {
				if part.Id == partId {
					part.IsCompleted = true
				}
				marshal, err := json.Marshal(parts)
				if err != nil {
					return err
				}
				os.WriteFile(db, marshal, 0666)
			}
		case err = <-failed:
			return err
		}

		if completedNumbers >= partNumbers {
			break
		}
	}

	newFilename, err := helper_os.NewFilename(dst, 10, nil)
	if err != nil {
		return err
	}

	fs, err := os.Create(newFilename)
	if err != nil {
		return err
	}
	offset = 0
	for i := 1; i < partNumbers+1; i++ {
		pf := filepath.Join(dstDir, fmt.Sprintf("%s.dp%d", dstBase, i))
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
	return nil
}
