package s3

import (
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
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	Domain          string
	Region          string
	s3              *s3.S3
}

type Instance struct {
	PutTimeout      time.Duration `yaml:"put_timeout"`
	DownloadTimeout time.Duration `yaml:"download_timeout"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
}

const name = "s3"

var instance *Instance

func (inst *Instance) NewClient(client *Client) (*Client, error) {
	var err error
	creds := credentials.NewStaticCredentials(client.AccessKeyId, client.AccessKeySecret, "")
	region := client.Region
	endpoint := client.Endpoint
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
	client.s3 = s3.New(s)
	return client, nil
}

func (inst *Instance) Run() error {
	instance = inst
	return nil
}

func (inst *Instance) GetName() string {
	return name
}

func Get() *Instance {
	return instance
}

func (c *Client) CreateFolder(key string) (err error) {
	ctx := context.Background()
	var cancelFn func()
	if instance.PutTimeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, instance.PutTimeout)
	}
	if cancelFn != nil {
		defer cancelFn()
	}
	k := folderName(key)
	_, err = c.s3.PutObjectWithContext(ctx, &s3.PutObjectInput{Bucket: aws.String(c.BucketName), Key: aws.String(k)})
	return err
}

func (c *Client) Download(key string, offset, size int64) (io.ReadCloser, error) {
	ctx := context.Background()
	var cancelFn func()
	if instance.DownloadTimeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, instance.DownloadTimeout)
	}
	if cancelFn != nil {
		defer cancelFn()
	}
	k := folderName(key)
	res, err := c.s3.GetObjectWithContext(ctx, &s3.GetObjectInput{Bucket: aws.String(c.BucketName), Key: aws.String(k), Range: aws.String(helper_http.GenRange(offset, size))})
	return res.Body, err
}

func (c *Client) List(prefix string) ([]*s3.Object, error) {
	delimiter := "/"

	ctx := context.Background()
	var cancelFn func()
	if instance.ReadTimeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, instance.ReadTimeout)
	}
	if cancelFn != nil {
		defer cancelFn()
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
	var cancelFn func()
	if instance.ReadTimeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, instance.ReadTimeout)
	}
	if cancelFn != nil {
		defer cancelFn()
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
	var cancelFn func()
	if instance.ReadTimeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, instance.ReadTimeout)
	}
	if cancelFn != nil {
		defer cancelFn()
	}
	return c.s3.HeadObjectWithContext(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(c.BucketName),
		Key:    aws.String(key),
	})
}

func folderName(name string) string {
	if !strings.HasSuffix(name, "/") {
		return name + "/"
	}
	return name
}
