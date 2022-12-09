package download

import (
	helper_s3 "github.com/langwan/langgo/helpers/s3"
	"io"
)

type S3Reader struct {
	ObjectName string
	Client     *helper_s3.Client
}

func (o S3Reader) GetFileSize() (int64, error) {
	head, err := o.Client.HeadObject(o.ObjectName)
	if err != nil {
		return 0, err
	}
	return *head.ContentLength, nil
}

func (o S3Reader) GetObjectByRange(offset, size int64) (io.ReadCloser, error) {
	return o.Client.GetObjectByRange(o.ObjectName, offset, size)
}
