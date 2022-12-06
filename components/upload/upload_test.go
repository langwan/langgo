package upload

import (
	"context"
	"fmt"
	"github.com/langwan/langgo"
	langgo_s3 "github.com/langwan/langgo/components/s3"
	helper_progress "github.com/langwan/langgo/helpers/progress"
	"testing"
	"time"
)

type MyPartList struct {
	uploadId string
	parts    []*Part
}

func (m *MyPartList) RomoveParts() error {
	return nil
}

func (m *MyPartList) Load() ([]*Part, error) {
	return nil, nil
}

func (m *MyPartList) SetList(parts []*Part) {
	m.parts = parts
}

func (m *MyPartList) GetList() []*Part {
	return m.parts
}

func (m *MyPartList) SavePart(part *Part) error {
	return nil
}

func (m *MyPartList) GetUploadId() string {
	return m.uploadId
}

func (m *MyPartList) SetUploadId(uploadId string) error {
	m.uploadId = uploadId
	return nil
}

type Listener struct {
}

func (l Listener) ProgressChanged(event *helper_progress.ProgressEvent) {
	fmt.Println(event)
}

func Test_s3Write(t *testing.T) {
	langgo.Run(&langgo_s3.Instance{
		PutTimeout:      time.Hour,
		DownloadTimeout: time.Hour,
		ReadTimeout:     time.Hour,
	})
	s3Client, err := langgo_s3.Get().NewClient(&langgo_s3.Client{
		Endpoint:        "",
		AccessKeyId:     "",
		AccessKeySecret: "",
		BucketName:      "",
		Domain:          "",
		Region:          "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	s3Writer := S3Writer{
		S3Client: s3Client,
	}
	upload := Upload{
		Workers:  5,
		PartSize: 1024 * 1024 * 1,
	}
	partList := MyPartList{}
	upload.Init()
	upload.Upload(context.Background(), "sample.mp4", "upload_test.mp4", &partList, &s3Writer, &Listener{})
}
