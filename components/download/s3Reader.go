package download

import (
	"github.com/langwan/langgo/components/s3"

	"io"
)

type S3Reader struct {
	ObjectName string
	Client     *s3.Client
}

func (o S3Reader) GetFileSize() (int64, error) {
	head, err := o.Client.HeadObject(o.ObjectName)
	if err != nil {
		return 0, err
	}
	return *head.ContentLength, nil
}

func (o S3Reader) OpenRange(offset, size int64) (io.ReadCloser, error) {
	return o.Client.Download(o.ObjectName, offset, size)
}
