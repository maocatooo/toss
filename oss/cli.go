package oss

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	_ "github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"", args...)
	os.Exit(1)
}

type Client struct {
	sess *session.Session
}

func NewClient(c Config) *Client {
	return &Client{sess: newSess(c)}
}

func newSess(c Config) *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(c.AK, c.SK, ""),
		Endpoint:    aws.String(c.Endpoint),
		Region:      aws.String(c.Region),
		DisableSSL:  aws.Bool(false),
		/*
			// false 会使用 virtual-host style方式， http://end_point -> http://{bucket}.end_point
			// true 会使用 强制使用路径方式， http://end_point -> http://end_point/bucket
		*/
		S3ForcePathStyle: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}
	return sess
}

func (c *Client) buckets() ([]*s3.Bucket, error) {
	svc := s3.New(c.sess)
	result, err := svc.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
		return nil, err
	}
	return result.Buckets, nil

}

func (c *Client) getFileAndFolder(bucket string, prefix string) ([]*s3.Object, error) {
	svc := s3.New(c.sess)
	params := &s3.ListObjectsInput{
		Bucket: aws.String(fmt.Sprint("/", bucket)),
		Prefix: aws.String(prefix),
	}
	resp, err := svc.ListObjects(params)

	if err != nil {
		exitErrorf("Unable to list items in bucket %q, %v", bucket, err)
		return nil, err
	}
	return resp.Contents, nil
}

func getFile(sess *session.Session, bucket, item string) {

	file, err := os.Create(item)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(""),
			Key:    aws.String(item),
		})
	if err != nil {
		exitErrorf("Unable to download item %q, %v", item, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func (c *Client) delete(bucket string, filename string) (*s3.DeleteObjectOutput, error) {
	svc := s3.New(c.sess)
	res, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(fmt.Sprint("/", bucket)), Key: aws.String(filename)})
	if err != nil {
		exitErrorf("Unable to delete object %q from bucket %q, %v", filename, bucket, err)
		return nil, err
	}
	return res, nil
}

func (c *Client) deleteOnPrefix(bucket string, prefix string) error {
	// deleteAll
	if len(prefix) > 0 && prefix[0] == '/' {
		prefix = prefix[1:]
	}
	res, err := c.getFileAndFolder(bucket, prefix)
	if err != nil {
		return err
	}
	for _, item := range res {
		_, err := c.delete(bucket, *item.Key)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) upload(bucket string, filename string, fileData []byte) (*s3manager.UploadOutput, error) {

	uploader := s3manager.NewUploader(c.sess)
	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(fmt.Sprint("/", bucket)),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(fileData),
	})
	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
		return nil, err
	}
	//fmt.Println("Successfully uploaded to", output.Location)
	return output, nil
}
