package oss

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type object struct {
	name    string
	abs     string
	osspath string
}

type info struct {
	abs     string
	objects []object
}

// 递归src目录
func recursiveSrc(src string) []string {
	// 读取src目录下的所有文件和目录
	source, err := os.ReadDir(src)
	if err != nil {
		fmt.Println(src)
		panic(err)
	}
	var files []string
	for _, file := range source {
		if file.IsDir() {
			// 递归src目录下的目录
			files = append(files, recursiveSrc(filepath.Join(src, file.Name()))...)
		} else {
			// 将src目录下的文件添加到files切片中
			files = append(files, filepath.Join(src, file.Name()))
		}
	}
	return files
}

func toInfo(c Config) info {

	var files []string
	if isFile(c.Source) {
		files = append(files, c.Source)
	} else {
		files = recursiveSrc(c.Source)
	}
	return format(c, files)
}

func isFile(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		panic(err)
	}
	return !info.IsDir()
}

func format(c Config, path []string) (fo info) {

	abs, _ := filepath.Abs(c.Source)
	fo.abs = abs
	for _, item := range path {
		res, _ := filepath.Abs(item)
		relativePath, _ := filepath.Rel(abs, res)
		fo.objects = append(fo.objects, object{
			abs:     res,
			osspath: ossPathJoin(c.Target, relativePath),
			name:    relativePath})
	}
	return fo
}

func ossPathJoin(ss ...string) string {
	return strings.ReplaceAll(path.Join(ss...), "\\", "/")
}
