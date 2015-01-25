// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Tin Nguyen <tinntt@penlook.com>
package s3

import (
	"fmt"
	"github.com/mitchellh/goamz/aws"
	goamz_s3 "github.com/mitchellh/goamz/s3"
	"io/ioutil"
	"log"
)

type S3 struct {
}

func (s3 S3) Put(bucketName string, fileName string, objectName string, contType string) {
	auth, err := aws.EnvAuth()

	if err != nil {
		log.Fatal(err)
	}
	if contType == "" {
		contType = "text/plain"
	}

	client := goamz_s3.New(auth, aws.APNortheast)
	bucket := client.Bucket(bucketName)

	data, err := ioutil.ReadFile(fileName)
	err = bucket.Put(objectName, data, contType, goamz_s3.BucketOwnerFull)

	if err != nil {
		log.Fatal(err)
	}
}

func (s3 S3) Get(bucketName string, objectName string, fileName string, contType string) {
	auth, err := aws.EnvAuth()

	if err != nil {
		log.Fatal(err)
	}

	if contType == "" {
		contType = "text/plain"
	}

	client := goamz_s3.New(auth, aws.APNortheast)
	bucket := client.Bucket(bucketName)

	data, err := bucket.Get(objectName)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(fileName, data, 0666)

	if err != nil {
		log.Fatal(err)
	}
}

func (s3 S3) Delete(bucketName string, objectName string) {
	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := goamz_s3.New(auth, aws.APNortheast)
	bucket := client.Bucket(bucketName)

	err = bucket.Del(objectName)
	if err != nil {
		log.Fatal(err)
	}
}

func (s3 S3) ListAllBucket() *goamz_s3.ListBucketsResp {
	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}

	client := goamz_s3.New(auth, aws.APNortheast)
	resp, err := client.ListBuckets()

	if err != nil {
		log.Fatal(err)
	}

	return resp
}
