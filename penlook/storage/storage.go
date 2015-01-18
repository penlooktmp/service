// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Tin Nguyen <tinntt@penlook.com>
package storage

import (
	"code.google.com/p/goauth2/oauth"
	gstorage "code.google.com/p/google-api-go-client/storage/v1"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Storage struct {
	BucketName   string
	ProjectID    string
	ClientId     string
	ClientSecret string

	// For the basic sample, these variables need not be changed.
	Scope       string
	AuthURL     string
	TokenURL    string
	EntityName  string
	RedirectURL string
	CacheFile   string
	Service     *gstorage.Service
}

var (
	funcName   = flag.String("flag", "", "function name to excute.\n		add - to add new object into bucket.\n		delete - to delete object from bucket")
	fileName   = flag.String("file", "", "file name to upload to storage") // The name of the local file to upload.
	objectName = flag.String("object", "", "object name in storage")       // This can be changed to any valid object name.
	code       = flag.String("code", "", "Authorization Code")
)

func (storage Storage) Fatalf(errorMessage string, args ...interface{}) {
	storage.RestoreOriginalState()
	log.Fatalf("Dying with error:\n"+errorMessage, args...)
}

func (storage Storage) RestoreOriginalState() bool {
	succeeded := true

	// Delete an object from a bucket.
	if err := storage.Service.Objects.Delete(storage.BucketName, *objectName).Do(); err == nil {
		fmt.Printf("Successfully deleted %s/%s during cleanup.\n\n", storage.BucketName, objectName)
	} else {
		// If the object exists but wasn't deleted, the bucket deletion will also fail.
		fmt.Printf("Could not delete object during cleanup: %v\n\n", err)
	}

	// Delete a bucket in the project
	if err := storage.Service.Buckets.Delete(storage.BucketName).Do(); err == nil {
		fmt.Printf("Successfully deleted bucket %s during cleanup.\n\n", storage.BucketName)
	} else {
		succeeded = false
		fmt.Printf("Could not delete bucket during cleanup: %v\n\n", err)
	}

	if !succeeded {

		fmt.Println("WARNING: Final cleanup attempt failed. Original state could not be restored.\n")
	}
	return succeeded
}

func (storage Storage) CreateService() *gstorage.Service {
	var config = &oauth.Config{
		ClientId:     storage.ClientId,
		ClientSecret: storage.ClientSecret,
		Scope:        storage.Scope,
		AuthURL:      storage.AuthURL,
		TokenURL:     storage.TokenURL,
		TokenCache:   oauth.CacheFile(storage.CacheFile),
		RedirectURL:  storage.RedirectURL,
	}
	transport := &oauth.Transport{
		Config:    config,
		Transport: http.DefaultTransport,
	}

	token, err := config.TokenCache.Token()
	if err != nil {
		if *code == "" {
			url := config.AuthCodeURL("")
			fmt.Println(url)
			os.Exit(1)
		}

		token, err = transport.Exchange(*code)
		if err != nil {
			log.Fatal("exchange", err)
		}
		fmt.Printf("Token is chached in %v\n", config.TokenCache)
	}
	transport.Token = token

	httpClient := transport.Client()
	service, err := gstorage.New(httpClient)
	return service
}

func (storage Storage) ListAllBucket() {
	if res, err := storage.Service.Buckets.List(storage.ProjectID).Do(); err == nil {
		fmt.Println("Buckets:")
		for _, item := range res.Items {
			fmt.Println(item.Id)
		}
		fmt.Println()
	} else {
		storage.Fatalf("Buckets.List failed: %v", err)
	}
}

func (storage Storage) InsertObjectToBucket() {
	object := &gstorage.Object{Name: *objectName}
	file, err := os.Open(*fileName)
	if err != nil {
		storage.Fatalf("Error opening %q: %v", *fileName, err)
	}
	if res, err := storage.Service.Objects.Insert(storage.BucketName, object).Media(file).Do(); err == nil {
		fmt.Printf("Created object %v at location %v\n\n", res.Name, res.SelfLink)
	} else {
		storage.Fatalf("Objects.Insert failed: %v", err)
	}
}

func (storage Storage) DeleteObjectFromBucket() {
	// Delete an object from a bucket.
	if err := storage.Service.Objects.Delete(storage.BucketName, *objectName).Do(); err == nil {
		fmt.Printf("Successfully deleted %s/%s during cleanup.\n\n", storage.BucketName, *objectName)
	} else {
		// If the object exists but wasn't deleted, the bucket deletion will also fail.
		fmt.Printf("Could not delete object during cleanup: %v\n\n", err)
	}
}

func (storage Storage) InsertACLForObject() {
	objectAcl := &gstorage.ObjectAccessControl{
		Bucket: storage.BucketName, Entity: storage.EntityName, Object: *objectName, Role: "READER",
	}
	if res, err := storage.Service.ObjectAccessControls.Insert(storage.BucketName, *objectName, objectAcl).Do(); err == nil {
		fmt.Printf("Result of inserting ACL for %v/%v:\n%v\n\n", storage.BucketName, *objectName, res)
	} else {
		storage.Fatalf("Failed to insert ACL for %s/%s: %v.", storage.BucketName, *objectName, err)
	}
}

func main() {
	storage := Storage{
		BucketName:   "static.penlook.com",
		ProjectID:    "penlook-app",
		ClientId:     "769231272797-jo8jbdshck6pfs1hb6dfki7rlkm407ko.apps.googleusercontent.com",
		ClientSecret: "ODZ9HsFaMkiMEZeE9tYgKp7j",

		Scope:       gstorage.DevstorageFull_controlScope,
		AuthURL:     "https://accounts.google.com/o/oauth2/auth",
		TokenURL:    "https://accounts.google.com/o/oauth2/token",
		EntityName:  "allUsers",
		RedirectURL: "urn:ietf:wg:oauth:2.0:oob",
		CacheFile:   "cache.json",
		Service:     nil,
	}
	storage.Service = storage.CreateService()
	storage.ListAllBucket()
}
