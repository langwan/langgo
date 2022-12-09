package download

import (
	"errors"
	"io"
	"net/http"
	"strconv"
)

type HttpReader struct {
	Url string
}

func (h HttpReader) GetFileSize() (int64, error) {
	resp, err := http.Get(h.Url)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	fileSize := resp.ContentLength
	if fileSize < 1 {
		return fileSize, errors.New("file size is error")
	}
	return fileSize, nil
}

func (h HttpReader) GetObjectByRange(offset, size int64) (io.ReadCloser, error) {
	request, err := http.NewRequest("GET", h.Url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set(
		"Range",
		"bytes="+strconv.FormatInt(offset, 10)+"-"+strconv.FormatInt(offset+size-1, 10),
	)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
