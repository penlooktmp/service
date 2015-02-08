package main

import (
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	sdkS3 "github.com/awslabs/aws-sdk-go/gen/s3"
	"os"
)

type S3 struct {
	s3cli *sdkS3.S3
}

func main() {

	var (
		bucket      string = "penlook-abc" // change to your convenience
		contenttype string = "binary/octet-stream"
	)

	s3 := S3{
		s3cli: S3Create(),
	}

	s3.PutObject(bucket, "abc.txt", contenttype)
	s3.ListObject(bucket)
}

func S3Create() *sdkS3.S3 {
	creds := aws.Creds("AKIAJTPBBS46Y7SJETQQ", "Zc/5Ii8qnb19xKiBxgLnRG2DlrRNpdMyWN10WSlr", "")
	cli := sdkS3.New(creds, "ap-northeast-1", nil)
	return cli
}

func (s3 *S3) PutObject(bucketName string, fileName string, contenttype string) {
	// open the file to upload
	fi, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("Error: no input file found in '%s'\n", os.Args[1])
		os.Exit(1)
	}
	fd, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	// create a bucket upload request and send
	objectreq := sdkS3.PutObjectRequest{
		ACL:           aws.String("public-read"),
		Bucket:        aws.String(bucketName),
		Body:          fd,
		ContentLength: aws.Long(int64(fi.Size())),
		ContentType:   aws.String(contenttype),
		Key:           aws.String(fi.Name()),
	}
	_, err = s3.s3cli.PutObject(&objectreq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%s\n", "https://s3.amazonaws.com/"+bucketName+"/"+fi.Name())
	}
}

func (s3 *S3) ListObject(bucketName string) {
	// list the content of the bucket
	listreq := sdkS3.ListObjectsRequest{
		Bucket: aws.StringValue(&bucketName),
	}
	listresp, err := s3.s3cli.ListObjects(&listreq)
	if err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Content of bucket '%s': %d files\n", bucketName, len(listresp.Contents))
		for _, obj := range listresp.Contents {
			fmt.Println("-", *obj.Key)
		}
	}
}
