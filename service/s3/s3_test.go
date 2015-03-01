package s3

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func abc() {

	// S3 service in here
	// var (
	// 	bucket      string = "penlook-abc" // change to your convenience
	// 	contenttype string = "binary/octet-stream"
	// )

	// s3 := S3{
	// 	s3cli: S3Create("AKIAJJCFXMQIASSWOLWQ", "/ubBzhnZN/poTTeLN1S+E6v27a8WQo2LpuwAuYqm"),
	// }

	// s3.PutObject(bucket, "abc1.txt", contenttype)
	// s3.ListObject(bucket)
	// s3.GetObject(bucket, "sample.txt")
	// s3.DeleteObject(bucket, "sample1.txt")
}

var s3 = S3{}

func TestS3Config(t *testing.T) {
	assert := assert.New(t)
	s3.GetConfig()
	assert.NotNil(s3.ID)
	assert.NotNil(s3.Secret)

	assert.Equal(20, len(s3.ID))
	assert.Equal(40, len(s3.Secret))
	fmt.Println()
}

func TestS3Create(t *testing.T) {
	assert := assert.New(t)
	s3.Create()
	assert.NotNil(s3.s3cli)
}

func TestS3CheckExistFile(t *testing.T) {
	filename := "test"
	assert := assert.New(t)
	os.Mkdir(filename, 0644)
	fi, fd := s3.CheckExistFile(filename)
	os.Remove(filename)
	assert.NotNil(fi)
	assert.NotNil(fd)
}
