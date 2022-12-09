package helper_oss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"

	"io"
)

type Client struct {
	Bucket *oss.Bucket
}

func NewClient(endpoint, accessKeyId, accessKeySecret, bucketName string) (*Client, error) {
	client := Client{}
	ossClient, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	bucket, err := ossClient.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	client.Bucket = bucket
	return &client, nil
}

func (c *Client) PutObject(key string, content []byte) error {
	err := c.Bucket.PutObject(key, bytes.NewReader(content))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) PutObjectByReader(key string, reader io.Reader) error {
	err := c.Bucket.PutObject(key, reader)
	if err != nil {
		return nil
	}
	return nil
}

func (c *Client) GetObject(key string) ([]byte, error) {
	body, err := c.Bucket.GetObject(key)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetObjectByRange(key string, offset int64, size int64) ([]byte, error) {
	body, err := c.Bucket.GetObject(key, oss.Range(offset, offset+size-1))
	if err != nil {
		return nil, err
	}
	defer body.Close()
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetObjectDetailedMeta(key string) (http.Header, error) {
	return c.Bucket.GetObjectDetailedMeta(key)
}

func (c *Client) IsObjectExist(key string) (bool, error) {
	isExist, err := c.Bucket.IsObjectExist(key)
	if err != nil {
		return false, err
	}
	return isExist, nil
}

func (c *Client) DeleteObject(key string) error {
	err := c.Bucket.DeleteObject(key)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) RemoveSubDir(key string) error {
	marker := oss.Marker("")
	prefix := oss.Prefix(key)
	count := 0

	for {
		lor, err := c.Bucket.ListObjects(marker, prefix)
		if err != nil {
			return err
		}

		var objects []string
		for _, object := range lor.Objects {
			objects = append(objects, object.Key)
		}

		if objects == nil {
			return nil
		}

		delRes, err := c.Bucket.DeleteObjects(objects, oss.DeleteObjectsQuiet(true))
		if err != nil {
			return err
		}

		if len(delRes.DeletedObjects) > 0 {
			return err
		}

		count += len(objects)

		prefix = oss.Prefix(lor.Prefix)
		marker = oss.Marker(lor.NextMarker)
		if !lor.IsTruncated {
			break
		}
	}
	return nil
}

func (c *Client) InitiateMultipartUpload(key string) (oss.InitiateMultipartUploadResult, error) {
	return c.Bucket.InitiateMultipartUpload(key)
}

func (c *Client) UploadPart(imur oss.InitiateMultipartUploadResult, reader io.Reader,
	partSize int64, partNumber int, options ...oss.Option) (oss.UploadPart, error) {
	return c.Bucket.UploadPart(imur, reader, partSize, partNumber, options...)
}

func (c *Client) CompleteMultipartUpload(imur oss.InitiateMultipartUploadResult,
	parts []oss.UploadPart, options ...oss.Option) (oss.CompleteMultipartUploadResult, error) {
	return c.Bucket.CompleteMultipartUpload(imur, parts)
}
