// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Tin Nguyen <tinntt@penlook.com>
package storage

import (
	gstorage "code.google.com/p/google-api-go-client/storage/v1"
	// "github.com/stretchr/testify/assert"
	"testing"
)

func TestStore(t *testing.T) {
	// assert := assert.New(t)

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
