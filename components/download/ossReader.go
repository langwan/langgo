package download

import (
	"bytes"
	helper_oss "github.com/langwan/langgo/helpers/oss"
	"io"
	"strconv"
)

type OssReader struct {
	ObjectName string
	Client     *helper_oss.Client
}

func (o OssReader) GetFileSize() (int64, error) {
	props, err := o.Client.GetObjectDetailedMeta(o.ObjectName)
	if err != nil {
		return -1, err
	}
	contentLength := props.Get("Content-Length")
	fileSize, err := strconv.ParseInt(contentLength, 10, 64)
	if err != nil {
		return -1, err
	}
	return fileSize, nil
}

func (o OssReader) GetObjectByRange(offset, size int64) (io.ReadCloser, error) {
	body, err := o.Client.GetObjectByRange(o.ObjectName, offset, size)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(body)
	return io.NopCloser(reader), nil
}
