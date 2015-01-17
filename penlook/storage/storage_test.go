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
		bucketName:   "static.penlook.com",
		projectID:    "penlook-app",
		clientId:     "769231272797-jo8jbdshck6pfs1hb6dfki7rlkm407ko.apps.googleusercontent.com",
		clientSecret: "ODZ9HsFaMkiMEZeE9tYgKp7j",

		scope:       gstorage.DevstorageFull_controlScope,
		authURL:     "https://accounts.google.com/o/oauth2/auth",
		tokenURL:    "https://accounts.google.com/o/oauth2/token",
		entityName:  "allUsers",
		redirectURL: "urn:ietf:wg:oauth:2.0:oob",
		cacheFile:   "config.json",
		service:     nil,
	}
	storage.CreateService()
}
