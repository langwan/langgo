package upload

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	helper_oss "github.com/langwan/langgo/helpers/oss"
)

type OssWriter struct {
	Client *helper_oss.Client
}

func (w *OssWriter) Create(key string) (string, error) {
	upload, err := w.Client.InitiateMultipartUpload(key)
	if err != nil {
		return "", err
	}
	return upload.UploadID, nil
}

func (w *OssWriter) UploadPart(key string, uploadId string, partId int64, partSize int64, data []byte) (string, error) {
	imur := oss.InitiateMultipartUploadResult{
		Bucket:   w.Client.Bucket.BucketName,
		Key:      key,
		UploadID: uploadId,
	}
	output, err := w.Client.UploadPart(imur, bytes.NewReader(data), partSize, int(partId))
	if err != nil {
		return "", err
	}

	return output.ETag, nil
}

func (w *OssWriter) Completed(key string, uploadId string, parts []*Part) error {
	var ossParts []oss.UploadPart
	for _, part := range parts {
		ossParts = append(ossParts, oss.UploadPart{
			PartNumber: part.Id,
			ETag:       part.ETag,
		})
	}
	imur := oss.InitiateMultipartUploadResult{
		Bucket:   w.Client.Bucket.BucketName,
		Key:      key,
		UploadID: uploadId,
	}
	_, err := w.Client.CompleteMultipartUpload(imur, ossParts)
	if err != nil {
		return err
	}
	return nil
}
