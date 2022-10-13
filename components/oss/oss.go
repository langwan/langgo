package oss

import (
	"bytes"
	"fmt"
	aliyunOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/log"
	"io"
	"io/ioutil"
	"os"
)

const name = "oss"

type Instance struct {
	Endpoint        string `yaml:"endpoint" json:"endpoint"`
	AccessKeyId     string `yaml:"access_key_id" json:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret" json:"access_key_secret"`
	BucketName      string `yaml:"bucket_name" json:"bucket_name"`
	Domain          string `yaml:"domain" json:"domain"`
}

var instance *Instance

func (i *Instance) Load() error {
	core.GetComponentConfiguration(name, i)
	return i.Run()
}

func (i *Instance) Run() error {
	instance = i
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func PutObject(uri string, content []byte) error {
	// 创建OSSClient实例。
	client, err := aliyunOss.New(instance.Endpoint, instance.AccessKeyId, instance.AccessKeySecret)
	if err != nil {
		return err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(instance.BucketName)
	if err != nil {
		return err
	}
	// 上传Byte数组。
	err = bucket.PutObject(uri, bytes.NewReader(content))
	if err != nil {
		return err
	}
	return nil
}

func PutObjectByReader(uri string, reader io.Reader) error {
	// 创建OSSClient实例。
	client, err := aliyunOss.New(instance.Endpoint, instance.AccessKeyId, instance.AccessKeySecret)
	if err != nil {
		return err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(instance.BucketName)
	if err != nil {
		return err
	}
	// 上传Byte数组。
	err = bucket.PutObject(uri, reader)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return nil
}

func GetObject(uri string) ([]byte, error) {
	// 创建OSSClient实例。
	client, err := aliyunOss.New(instance.Endpoint, instance.AccessKeyId, instance.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(instance.BucketName)
	if err != nil {
		return nil, err
	}

	// 下载文件到流。
	body, err := bucket.GetObject(uri)
	if err != nil {
		return nil, err

	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func IsObjectExist(uri string) (bool, error) {
	// 创建OSSClient实例。
	client, err := aliyunOss.New(instance.Endpoint, instance.AccessKeyId, instance.AccessKeySecret)
	if err != nil {
		return false, err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(instance.BucketName)
	if err != nil {
		return false, err
	}

	// 判断文件是否存在。
	isExist, err := bucket.IsObjectExist(uri)
	if err != nil {
		return false, err
	}
	return isExist, nil
}

func DeleteObject(uri string) error {
	// 创建OSSClient实例。
	client, err := aliyunOss.New(instance.Endpoint, instance.AccessKeyId, instance.AccessKeySecret)
	if err != nil {
		return err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(instance.BucketName)
	if err != nil {
		return err
	}
	// 上传Byte数组。
	err = bucket.DeleteObject(uri)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return nil
}

func Remove(uri string) error {
	client, err := aliyunOss.New(instance.Endpoint, instance.AccessKeyId, instance.AccessKeySecret)
	bucket, err := client.Bucket(instance.BucketName)
	if err != nil {
		return err
	}
	return bucket.DeleteObject(uri)
}

func RemoveSubDir(dir string) error {
	// 创建OSSClient实例。
	client, err := aliyunOss.New(instance.Endpoint, instance.AccessKeyId, instance.AccessKeySecret)
	bucket, err := client.Bucket(instance.BucketName)
	if err != nil {
		return err
	}
	// 列举所有包含指定前缀的文件并删除。
	marker := aliyunOss.Marker("")
	prefix := aliyunOss.Prefix(dir)
	count := 0

	for {
		lor, err := bucket.ListObjects(marker, prefix)
		if err != nil {
			log.Logger("component", name).Warn().Err(err).Send()
			return err
		}

		var objects []string
		for _, object := range lor.Objects {
			objects = append(objects, object.Key)
		}

		if objects == nil {
			return nil
		}

		delRes, err := bucket.DeleteObjects(objects, aliyunOss.DeleteObjectsQuiet(true))
		if err != nil {

			log.Logger("component", name).Warn().Err(err).Send()
			return err
		}

		if len(delRes.DeletedObjects) > 0 {
			log.Logger("component", name).Warn().Err(err).Send()
			return err
		}

		count += len(objects)

		prefix = aliyunOss.Prefix(lor.Prefix)
		marker = aliyunOss.Marker(lor.NextMarker)
		if !lor.IsTruncated {
			break
		}
	}
	fmt.Printf("success,total delete object count:%d\n", count)
	return nil
}
