package download

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"strconv"
)

type OssReader struct {
	ObjectName string
	Bucket     *oss.Bucket
}

func (o OssReader) GetFileSize() (int64, error) {
	props, err := o.Bucket.GetObjectDetailedMeta(o.ObjectName)
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

func (o OssReader) OpenRange(offset, size int64) (io.ReadCloser, error) {
	body, err := o.Bucket.GetObject(o.ObjectName, oss.Range(offset, offset+size-1))
	if err != nil {
		return nil, err
	}
	return body, nil
}
