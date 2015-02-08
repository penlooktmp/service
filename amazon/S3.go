package main

import (
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/s3"
)

func main() {
	creds := aws.Creds("AKIAJTPBBS46Y7SJETQQ", "Zc/5Ii8qnb19xKiBxgLnRG2DlrRNpdMyWN10WSlr", "")
	cli := s3.New(creds, "us-west-2", nil)
	resp, err := cli.ListBuckets()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Buckets)
}
