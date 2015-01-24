package main

import (
	// "fmt"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"io/ioutil"
	"log"
)

func main() {
	// Put("penlook-abc", "/home/tinntt/228990_depotcache_1.csd", "sample2.txt", "text/plain")
	// Delete("penlook-abc", "sample1.txt")
	// resp := ListAllBucket()
	// log.Print(fmt.Sprintf("%T %+v", resp.Buckets[0], resp.Buckets[0]))
	Get("penlook-abc", "sample2.txt", "/home/tinntt/test.txt", "text/plain")
}

func Put(bucketName string, fileName string, objectName string, contType string) {
	auth, err := aws.EnvAuth()

	if err != nil {
		log.Fatal(err)
	}
	if contType == "" {
		contType = "text/plain"
	}

	client := s3.New(auth, aws.APNortheast)
	bucket := client.Bucket(bucketName)

	data, err := ioutil.ReadFile(fileName)
	err = bucket.Put(objectName, data, contType, s3.BucketOwnerFull)

	if err != nil {
		log.Fatal(err)
	}
}

func Get(bucketName string, objectName string, fileName string, contType string) {
	auth, err := aws.EnvAuth()

	if err != nil {
		log.Fatal(err)
	}

	if contType == "" {
		contType = "text/plain"
	}

	client := s3.New(auth, aws.APNortheast)
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

func Delete(bucketName string, objectName string) {
	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := s3.New(auth, aws.APNortheast)
	bucket := client.Bucket(bucketName)

	err = bucket.Del(objectName)
	if err != nil {
		log.Fatal(err)
	}
}

func ListAllBucket() *s3.ListBucketsResp {
	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}

	client := s3.New(auth, aws.APNortheast)
	resp, err := client.ListBuckets()

	if err != nil {
		log.Fatal(err)
	}

	return resp
}
