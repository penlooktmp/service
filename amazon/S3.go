package main

import (
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/s3"
	"os"
)

func main() {

	var (
		err         error
		bucket      string = "penlook-abc" // change to your convenience
		fd          *os.File
		contenttype string = "binary/octet-stream"
	)

	creds := aws.Creds("AKIAJTPBBS46Y7SJETQQ", "Zc/5Ii8qnb19xKiBxgLnRG2DlrRNpdMyWN10WSlr", "")
	cli := s3.New(creds, "ap-northeast-1", nil)

	// open the file to upload
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <inputfile>\n", os.Args[0])
		os.Exit(1)
	}
	fi, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Printf("Error: no input file found in '%s'\n", os.Args[1])
		os.Exit(1)
	}
	fd, err = os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	// create a bucket upload request and send
	objectreq := s3.PutObjectRequest{
		ACL:           aws.String("public-read"),
		Bucket:        aws.String(bucket),
		Body:          fd,
		ContentLength: aws.Long(int64(fi.Size())),
		ContentType:   aws.String(contenttype),
		Key:           aws.String(fi.Name()),
	}
	_, err = cli.PutObject(&objectreq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%s\n", "https://s3.amazonaws.com/"+bucket+"/"+fi.Name())
	}

	// list the content of the bucket
	listreq := s3.ListObjectsRequest{
		Bucket: aws.StringValue(&bucket),
	}
	listresp, err := cli.ListObjects(&listreq)
	if err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Content of bucket '%s': %d files\n", bucket, len(listresp.Contents))
		for _, obj := range listresp.Contents {
			fmt.Println("-", *obj.Key)
		}
	}
}
