package oss

import "testing"

func TestUpload(t *testing.T) {
	Upload(testConfig.Minio)
}

func TestUpload_Tencent(t *testing.T) {
	Upload(testConfig.Tencent)
}
