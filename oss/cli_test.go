package oss

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"testing"
)

var testConfig = struct {
	Tencent Config `yaml:"tencent"`
	Ali     Config `yaml:"ali"`
	Minio   Config `yaml:"minio"`
}{}

func TestMain(t *testing.M) {
	bs, err := os.ReadFile("../test/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bs, &testConfig)
	os.Exit(t.Run())
}

func TestNewClient(t *testing.T) {
	NewClient(testConfig.Minio)
}

func TestNewClient_buckets(t *testing.T) {
	bs, err := NewClient(testConfig.Minio).buckets()
	if err != nil {
		t.Error(err)
		return
	}
	for _, item := range bs {
		t.Log(item)
	}
}

func TestNewClient_uploadFile(t *testing.T) {
	res, err := NewClient(testConfig.Minio).upload(testConfig.Minio.Bucket, "1231/123", []byte("123333333"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

func TestNewClient_get_file_and_folder(t *testing.T) {
	{
		res, err := NewClient(testConfig.Minio).getFileAndFolder(testConfig.Minio.Bucket, "")
		if err != nil {
			t.Error(err)
			return
		}
		for _, item := range res {
			fmt.Println("Name:         ", *item.Key)
			fmt.Println("Last modified:", *item.LastModified)
			fmt.Println("Size:         ", *item.Size)
			fmt.Println("Storage class:", *item.StorageClass)
			fmt.Println("")
		}
	}
	{
		//search the prefix eq "123"
		res, err := NewClient(testConfig.Minio).getFileAndFolder(testConfig.Minio.Bucket, "1231")
		if err != nil {
			t.Error(err)
			return
		}
		for _, item := range res {
			fmt.Println("Name:         ", *item.Key)
			fmt.Println("Last modified:", *item.LastModified)
			fmt.Println("Size:         ", *item.Size)
			fmt.Println("Storage class:", *item.StorageClass)
			fmt.Println("")
		}
	}
}

func TestNewClient_delete(t *testing.T) {
	res, err := NewClient(testConfig.Minio).delete(testConfig.Minio.Bucket, "1231/123")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

func TestNewClient_deleteOnPrefix(t *testing.T) {
	err := NewClient(testConfig.Minio).deleteOnPrefix(testConfig.Minio.Bucket, "")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestNewClient_buckets_ali(t *testing.T) {
	bs, err := NewClient(testConfig.Ali).buckets()
	if err != nil {
		t.Error(err)
		return
	}
	for _, item := range bs {
		t.Log(item)
	}
}

func TestNewClient_uploadFile_ali(t *testing.T) {
	res, err := NewClient(testConfig.Ali).upload(testConfig.Ali.Bucket, "/1231/123.txt", []byte("123333333"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

func TestNewClient_delete_ali(t *testing.T) {
	res, err := NewClient(testConfig.Ali).delete(testConfig.Ali.Bucket, "/1231/123.txt")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

func TestNewClient_get_file_and_folder_ali(t *testing.T) {
	{
		res, err := NewClient(testConfig.Ali).getFileAndFolder(testConfig.Ali.Bucket, "")
		if err != nil {
			t.Error(err)
			return
		}
		for _, item := range res {
			fmt.Println("Name:         ", *item.Key)
			fmt.Println("Last modified:", *item.LastModified)
			fmt.Println("Size:         ", *item.Size)
			fmt.Println("Storage class:", *item.StorageClass)
			fmt.Println("")
		}
	}
	{
		//search the prefix eq "123"
		res, err := NewClient(testConfig.Ali).getFileAndFolder(testConfig.Ali.Bucket, "1231")
		if err != nil {
			t.Error(err)
			return
		}
		for _, item := range res {
			fmt.Println("Name:         ", *item.Key)
			fmt.Println("Last modified:", *item.LastModified)
			fmt.Println("Size:         ", *item.Size)
			fmt.Println("Storage class:", *item.StorageClass)
			fmt.Println("")
		}
	}
}

func TestNewClient_buckets_tencent(t *testing.T) {
	bs, err := NewClient(testConfig.Tencent).buckets()
	if err != nil {
		t.Error(err)
		return
	}
	for _, item := range bs {
		t.Log(item)
	}
}

func TestNewClient_get_file_and_folder_tencent(t *testing.T) {
	{
		res, err := NewClient(testConfig.Tencent).getFileAndFolder(testConfig.Tencent.Bucket, "")
		if err != nil {
			t.Error(err)
			return
		}
		for _, item := range res {
			fmt.Println("Name:         ", *item.Key)
			fmt.Println("Last modified:", *item.LastModified)
			fmt.Println("Size:         ", *item.Size)
			fmt.Println("Storage class:", *item.StorageClass)
			fmt.Println("")
		}
	}
	{
		//search the prefix eq "123"
		res, err := NewClient(testConfig.Tencent).getFileAndFolder(testConfig.Tencent.Bucket, "123")
		if err != nil {
			t.Error(err)
			return
		}
		for _, item := range res {
			fmt.Println("Name:         ", *item.Key)
			fmt.Println("Last modified:", *item.LastModified)
			fmt.Println("Size:         ", *item.Size)
			fmt.Println("Storage class:", *item.StorageClass)
			fmt.Println("")
		}
	}
}

func TestNewClient_uploadFile_tencent(t *testing.T) {
	res, err := NewClient(testConfig.Tencent).upload(testConfig.Tencent.Bucket, "/1231/123.txt", []byte("123333333"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

func TestNewClient_deleteOnPrefix_tencent(t *testing.T) {
	err := NewClient(testConfig.Tencent).deleteOnPrefix(testConfig.Tencent.Bucket, "/1231")
	if err != nil {
		t.Error(err)
		return
	}
}
