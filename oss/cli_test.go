package oss

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"testing"
)

var testConfig = struct {
	Tencent Config `yaml:"tencent"`
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
	res, err := NewClient(testConfig.Minio).upload("test123", "1231/123", []byte("123333333"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

func TestNewClient_get_file_and_folder(t *testing.T) {
	{
		res, err := NewClient(testConfig.Minio).getFileAndFolder("test123", "")
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
		res, err := NewClient(testConfig.Minio).getFileAndFolder("test123", "123")
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
	res, err := NewClient(testConfig.Minio).delete("test123", "1231/123")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

func TestNewClient_deleteOnPrefix(t *testing.T) {
	err := NewClient(testConfig.Minio).deleteOnPrefix("test123", "")
	if err != nil {
		t.Error(err)
		return
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
		res, err := NewClient(testConfig.Tencent).getFileAndFolder("/", "")
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
		res, err := NewClient(testConfig.Tencent).getFileAndFolder("", "123")
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
	res, err := NewClient(testConfig.Tencent).upload("", "/1231/123.txt", []byte("123333333"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

func TestNewClient_deleteOnPrefix_tencent(t *testing.T) {
	err := NewClient(testConfig.Tencent).deleteOnPrefix("", "/1231")
	if err != nil {
		t.Error(err)
		return
	}
}
