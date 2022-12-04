package download

import (
	"context"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/s3"
	"testing"
	"time"
)

func Test_S3(t *testing.T) {
	langgo.Run(&s3.Instance{
		PutTimeout:      time.Hour,
		DownloadTimeout: time.Hour,
		ReadTimeout:     time.Hour,
	})
	dl := Instance{
		Workers:  5,
		PartSize: "5m",
		BufSize:  "200k",
	}
	dl.Run()

	s3client, err := s3.Get().NewClient(&s3.Client{
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

	reader := S3Reader{
		ObjectName: "Homework - 1028.mp4",
		Client:     s3client,
	}
	dl.Download(context.Background(), "./s3.mp4", &reader, &Listener{})
}
