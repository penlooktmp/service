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
	bucketName   string
	projectID    string
	clientId     string
	clientSecret string

	// For the basic sample, these variables need not be changed.
	scope       string
	authURL     string
	tokenURL    string
	entityName  string
	redirectURL string
	service     *gstorage.Service
}

var (
	funcName   = flag.String("flag", "", "function name to excute.\n		add - to add new object into bucket.\n		delete - to delete object from bucket")
	fileName   = flag.String("file", "", "file name to upload to storage") // The name of the local file to upload.
	objectName = flag.String("object", "", "object name in storage")       // This can be changed to any valid object name.
	cacheFile  = flag.String("cache1", "cache.json", "Token cache file")
	code       = flag.String("code", "", "Authorization Code")
)

func (storage Storage) fatalf(errorMessage string, args ...interface{}) {
	storage.restoreOriginalState()
	log.Fatalf("Dying with error:\n"+errorMessage, args...)
}

func (storage Storage) restoreOriginalState() bool {
	succeeded := true

	// Delete an object from a bucket.
	if err := storage.service.Objects.Delete(storage.bucketName, *objectName).Do(); err == nil {
		fmt.Printf("Successfully deleted %s/%s during cleanup.\n\n", storage.bucketName, objectName)
	} else {
		// If the object exists but wasn't deleted, the bucket deletion will also fail.
		fmt.Printf("Could not delete object during cleanup: %v\n\n", err)
	}

	// Delete a bucket in the project
	if err := storage.service.Buckets.Delete(storage.bucketName).Do(); err == nil {
		fmt.Printf("Successfully deleted bucket %s during cleanup.\n\n", storage.bucketName)
	} else {
		succeeded = false
		fmt.Printf("Could not delete bucket during cleanup: %v\n\n", err)
	}

	if !succeeded {

		fmt.Println("WARNING: Final cleanup attempt failed. Original state could not be restored.\n")
	}
	return succeeded
}

func (storage Storage) createService() {
	var config = &oauth.Config{
		ClientId:     storage.clientId,
		ClientSecret: storage.clientSecret,
		Scope:        storage.scope,
		AuthURL:      storage.authURL,
		TokenURL:     storage.tokenURL,
		TokenCache:   oauth.CacheFile(*cacheFile),
		RedirectURL:  storage.redirectURL,
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
	storage.service, err = gstorage.New(httpClient)
}

func (storage Storage) listAllBucket() {
	if res, err := storage.service.Buckets.List(storage.projectID).Do(); err == nil {
		fmt.Println("Buckets:")
		for _, item := range res.Items {
			fmt.Println(item.Id)
		}
		fmt.Println()
	} else {
		storage.fatalf("Buckets.List failed: %v", err)
	}
}

func (storage Storage) insertObjectToBucket() {
	object := &gstorage.Object{Name: *objectName}
	file, err := os.Open(*fileName)
	if err != nil {
		storage.fatalf("Error opening %q: %v", *fileName, err)
	}
	if res, err := storage.service.Objects.Insert(storage.bucketName, object).Media(file).Do(); err == nil {
		fmt.Printf("Created object %v at location %v\n\n", res.Name, res.SelfLink)
	} else {
		storage.fatalf("Objects.Insert failed: %v", err)
	}
}

func (storage Storage) deleteObjectFromBucket() {
	// Delete an object from a bucket.
	if err := storage.service.Objects.Delete(storage.bucketName, *objectName).Do(); err == nil {
		fmt.Printf("Successfully deleted %s/%s during cleanup.\n\n", storage.bucketName, *objectName)
	} else {
		// If the object exists but wasn't deleted, the bucket deletion will also fail.
		fmt.Printf("Could not delete object during cleanup: %v\n\n", err)
	}
}

func (storage Storage) insertACLForObject() {
	objectAcl := &gstorage.ObjectAccessControl{
		Bucket: storage.bucketName, Entity: storage.entityName, Object: *objectName, Role: "READER",
	}
	if res, err := storage.service.ObjectAccessControls.Insert(storage.bucketName, *objectName, objectAcl).Do(); err == nil {
		fmt.Printf("Result of inserting ACL for %v/%v:\n%v\n\n", storage.bucketName, *objectName, res)
	} else {
		storage.fatalf("Failed to insert ACL for %s/%s: %v.", storage.bucketName, *objectName, err)
	}
}
