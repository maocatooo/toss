package oss

import (
	"fmt"
	"os"
)

func Upload(c Config) {

	info := toInfo(c)
	if len(info.objects) == 0 {
		fmt.Println("no file to upload")
		return
	}

	cli := NewClient(c)
	// 删除目标文件夹下的所有文件
	err := cli.deleteOnPrefix(c.Bucket, c.Target)
	if err != nil {
		panic(err)
	}
	// 上传文件
	for _, o := range info.objects {
		bs, err := os.ReadFile(o.abs)
		if err != nil {
			panic(err)
		}

		_, err = cli.upload(c.Bucket, o.osspath, bs)
		if err != nil {
			panic(err)
		}
		fmt.Printf("uploaded file: %s ----> %s \n", o.abs, o.osspath)
	}
}
