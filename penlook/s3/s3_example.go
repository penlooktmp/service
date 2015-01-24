package main

import (
	// "fmt"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"log"
)

func main() {
	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := s3.New(auth, aws.APNortheast)
	bucket := client.Bucket("penlook-abc")

	data := []byte("Hello World!!")
	err = bucket.Put("sample.txt", data, "text/plain", s3.BucketOwnerFull)
	if err != nil {
		log.Fatal(err)
		// panic(err.Error())
	}
	// resp, err := client.ListBuckets()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Print(fmt.Sprintf("%T %+v", resp.Buckets[0], resp.Buckets[0]))
}
