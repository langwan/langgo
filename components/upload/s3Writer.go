package upload

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	helper_s3 "github.com/langwan/langgo/helpers/s3"
)

type S3Writer struct {
	S3Client *helper_s3.Client
}

func (w *S3Writer) Create(key string) (string, error) {
	upload, err := w.S3Client.CreateMultipartUpload(key)
	if err != nil {
		return "", err
	}
	return *upload.UploadId, nil
}

func (w *S3Writer) UploadPart(key string, uploadId string, partId int64, partSize int64, data []byte) (string, error) {
	output, err := w.S3Client.UploadPart(uploadId, data, key, partId)
	if err != nil {
		return "", err
	}

	return *output.ETag, nil
}

func (w *S3Writer) Completed(key string, uploadId string, parts []*Part) error {
	var s3Parts []*s3.CompletedPart
	for _, part := range parts {
		s3Parts = append(s3Parts, &s3.CompletedPart{
			ETag:       aws.String(part.ETag),
			PartNumber: aws.Int64(int64(part.Id)),
		})
	}

	_, err := w.S3Client.CompletedMultipartUpload(uploadId, key, s3Parts)
	if err != nil {
		return err
	}
	return nil
}
