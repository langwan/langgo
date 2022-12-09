package helper_s3

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/langwan/langgo/helpers/http"
	"io"
	"strings"
	"time"
)

type Client struct {
	BucketName string

	s3              *s3.S3
	WriteTimeout    time.Duration
	DownloadTimeout time.Duration
	ReadTimeout     time.Duration
}

type Option func(*Client)

func WithTimeout(readTimeout, writeTimeout, downloadTimeout time.Duration) Option {
	return func(client *Client) {
		client.ReadTimeout = readTimeout
		client.WriteTimeout = writeTimeout
		client.DownloadTimeout = downloadTimeout
	}
}

func NewClient(endpoint, accessKeyId, accessKeySecret, bucketName, region string, options ...Option) (*Client, error) {
	var err error
	creds := credentials.NewStaticCredentials(accessKeyId, accessKeySecret, "")

	cfg := &aws.Config{
		Region:           aws.String(region),
		Endpoint:         &endpoint,
		S3ForcePathStyle: aws.Bool(false),
		Credentials:      creds,
	}
	s, err := session.NewSession(cfg)
	if err != nil {
		return nil, err
	}

	client := Client{}
	client.BucketName = bucketName
	client.s3 = s3.New(s)

	for _, option := range options {
		option(&client)
	}

	return &client, nil
}

func (c *Client) CreateFolder(key string) (err error) {
	ctx := context.Background()

	if c.WriteTimeout > 0 {
		ctx, _ = context.WithTimeout(ctx, c.WriteTimeout)
	}

	k := folderName(key)
	_, err = c.s3.PutObjectWithContext(ctx, &s3.PutObjectInput{Bucket: aws.String(c.BucketName), Key: aws.String(k)})
	return err
}

func (c *Client) GetObjectByRange(key string, offset, size int64) (io.ReadCloser, error) {

	ctx := context.Background()
	if c.DownloadTimeout > 0 {
		ctx, _ = context.WithTimeout(ctx, c.DownloadTimeout)
	}

	res, err := c.s3.GetObjectWithContext(ctx, &s3.GetObjectInput{Bucket: aws.String(c.BucketName), Key: aws.String(key), Range: aws.String(helper_http.GenRange(offset, size))})
	return res.Body, err
}

func (c *Client) List(prefix string) ([]*s3.Object, error) {
	delimiter := "/"

	ctx := context.Background()

	if c.ReadTimeout > 0 {
		ctx, _ = context.WithTimeout(ctx, c.ReadTimeout)
	}

	var objects []*s3.Object
	err := c.s3.ListObjectsV2PagesWithContext(ctx, &s3.ListObjectsV2Input{Bucket: aws.String(c.BucketName), Prefix: aws.String(prefix), Delimiter: aws.String(delimiter)}, func(page *s3.ListObjectsV2Output, isLastPage bool) bool {
		for _, folder := range page.CommonPrefixes {

			attributes, err := c.s3.GetObjectAttributesWithContext(ctx, &s3.GetObjectAttributesInput{
				Bucket:           aws.String(c.BucketName),
				Key:              folder.Prefix,
				ObjectAttributes: []*string{aws.String("Last-Modified")},
			})
			if err != nil {
				continue
			}
			objects = append(objects, &s3.Object{
				Key:          folder.Prefix,
				LastModified: attributes.LastModified,
				Size:         aws.Int64(0),
			})
		}

		for _, object := range page.Contents {
			if *object.Key == *aws.String(prefix) {
				continue
			}
			objects = append(objects, object)
		}

		return len(page.Contents) == 0
	})
	if err != nil {
		return nil, err
	}

	for _, object := range objects {
		object.Key = aws.String(strings.Replace(*object.Key, prefix, "", 1))
	}

	return objects, nil
}

func (c *Client) GetObjectAttributes(key string, attributes []*string) (*s3.GetObjectAttributesOutput, error) {
	ctx := context.Background()

	if c.ReadTimeout > 0 {
		ctx, _ = context.WithTimeout(ctx, c.ReadTimeout)
	}

	output, err := c.s3.GetObjectAttributesWithContext(ctx, &s3.GetObjectAttributesInput{
		Bucket:           aws.String(c.BucketName),
		Key:              aws.String(key),
		ObjectAttributes: attributes,
	})
	return output, err
}

func (c *Client) HeadObject(key string) (*s3.HeadObjectOutput, error) {
	ctx := context.Background()

	if c.ReadTimeout > 0 {
		ctx, _ = context.WithTimeout(ctx, c.ReadTimeout)
	}

	return c.s3.HeadObjectWithContext(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(c.BucketName),
		Key:    aws.String(key),
	})
}

func (c *Client) CreateMultipartUpload(key string) (resp *s3.CreateMultipartUploadOutput, err error) {
	ctx := context.Background()

	if c.ReadTimeout > 0 {
		ctx, _ = context.WithTimeout(ctx, c.ReadTimeout)
	}

	return c.s3.CreateMultipartUploadWithContext(ctx, &s3.CreateMultipartUploadInput{Bucket: aws.String(c.BucketName), Key: aws.String(key)})
}

func (c *Client) UploadPart(uploadId string, body []byte, key string, partNumber int64) (resp *s3.UploadPartOutput, err error) {
	ctx := context.Background()
	var cancelFn func()
	if c.ReadTimeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, c.ReadTimeout)
	}
	if cancelFn != nil {
		defer cancelFn()
	}
	return c.s3.UploadPartWithContext(ctx, &s3.UploadPartInput{
		Bucket:     aws.String(c.BucketName),
		Key:        aws.String(key),
		PartNumber: aws.Int64(partNumber),
		UploadId:   aws.String(uploadId),
		Body:       bytes.NewReader(body),
	})
}

func (c *Client) CompletedMultipartUpload(uploadId string, key string, completedParts []*s3.CompletedPart) (resp *s3.CompleteMultipartUploadOutput, err error) {
	ctx := context.Background()
	var cancelFn func()
	if c.ReadTimeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, c.ReadTimeout)
	}
	if cancelFn != nil {
		defer cancelFn()
	}
	return c.s3.CompleteMultipartUploadWithContext(ctx, &s3.CompleteMultipartUploadInput{
		Bucket: aws.String(c.BucketName),
		Key:    aws.String(key),

		UploadId: aws.String(uploadId),
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: completedParts,
		},
	})
}

func folderName(name string) string {
	if !strings.HasSuffix(name, "/") {
		return name + "/"
	}
	return name
}
