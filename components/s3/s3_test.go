package s3

import (
	"github.com/langwan/langgo"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_CreateFolder(t *testing.T) {
	langgo.Run(&Instance{

		PutTimeout: 30 * time.Second,
	})

	s3client, err := Get().NewClient(&Client{
		Endpoint:        "",
		AccessKeyId:     "",
		AccessKeySecret: "",
		BucketName:      "",
		Domain:          "",
		Region:          "oss-cn-hangzhou",
	})

	assert.NoError(t, err)

	s3client.CreateFolder("langwan2")
}

func Test_List(t *testing.T) {
	langgo.Run(&Instance{
		ReadTimeout:     30 * time.Second,
		PutTimeout:      30 * time.Second,
		DownloadTimeout: time.Hour,
	})

	s3client, err := Get().NewClient(&Client{
		Endpoint:        "",
		AccessKeyId:     "",
		AccessKeySecret: "",
		BucketName:      "",
		Domain:          "",
		Region:          "oss-cn-hangzhou",
	})
	list, err := s3client.List("新建文件夹/")
	if err != nil {
		return
	}
	for _, object := range list {
		t.Log(object)
	}
}
