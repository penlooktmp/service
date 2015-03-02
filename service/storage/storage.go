/**
 * Penlook Project
 *
 * Copyright (c) 2015 Penlook Development Team
 *
 * --------------------------------------------------------------------
 *
 * This program is free software: you can redistribute it and/or
 * modify it under the terms of the GNU Affero General Public License
 * as published by the Free Software Foundation, either version 3
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public
 * License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 *
 * --------------------------------------------------------------------
 *
 * Author:
 *     Loi Nguyen       <loint@penlook.com>
 *     Tin Nguyen       <tinntt@penlook.com>
 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	sdkS3 "github.com/awslabs/aws-sdk-go/gen/s3"
	"github.com/penlook/service/module/redis"
	"io/ioutil"
	"os"
	"strings"
)

type S3 struct {
	s3cli  *sdkS3.S3
	ID     string
	Secret string
}

func (s3 *S3) GetConfig() {
	redis_ := redis.Redis{
		Name:   "Penlook",
		Server: "localhost:6379",
	}.Connect()

	result, _ := redis.String(redis_.Do("GET", "aws.yml"))

	decoder := json.NewDecoder(strings.NewReader(result))
	var data map[string]map[string]interface{}
	decoder.Decode(&data)

	s3.ID = data["access_key"]["key_id"].(string)
	s3.Secret = data["access_key"]["key_secret"].(string)
}

func (s3 *S3) Create() {
	creds := aws.Creds(s3.ID, s3.Secret, "")
	s3.s3cli = sdkS3.New(creds, "ap-northeast-1", nil)
}

func (s3 *S3) CheckExistFile(fileName string) (os.FileInfo, *os.File) {
	fi, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("Error: no input file found in '%s'\n", os.Args[1])
		os.Exit(1)
	}
	fd, err := os.Open(fileName)
	if err != nil {
		panic(err)
		return nil, nil
	}
	return fi, fd
}

func (s3 *S3) PutObject(bucketName string, fileName string, contenttype string) {
	// open the file to upload
	fi, fd := s3.CheckExistFile(fileName)
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
	_, err := s3.s3cli.PutObject(&objectreq)
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

func (s3 *S3) GetObject(bucketName string, fileName string) {
	// list the content of the bucket
	getobjreq := sdkS3.GetObjectRequest{
		Bucket: aws.StringValue(&bucketName),
		Key:    aws.StringValue(&fileName),
	}
	getobjresp, err := s3.s3cli.GetObject(&getobjreq)
	if err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		if b, err := ioutil.ReadAll(getobjresp.Body); err == nil {
			err := ioutil.WriteFile(fileName, b, 0644)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (s3 *S3) DeleteObject(bucketName string, fileName string) {
	delobjreq := sdkS3.DeleteObjectRequest{
		Bucket: aws.StringValue(&bucketName),
		Key:    aws.StringValue(&fileName),
	}
	_, err := s3.s3cli.DeleteObject(&delobjreq)
	if err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
