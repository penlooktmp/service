// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Tin Nguyen <tinntt@penlook.com>
package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	bucketName   string "static.penlook.com"
	projectID    string "penlook-app"
	clientId     string "769231272797-jo8jbdshck6pfs1hb6dfki7rlkm407ko.apps.googleusercontent.com"
	clientSecret string "ODZ9HsFaMkiMEZeE9tYgKp7j"

	// For the basic sample, these variables need not be changed.
	scope       string gstorage.DevstorageFull_controlScope
	authURL     string "https://accounts.google.com/o/oauth2/auth"
	tokenURL    string "https://accounts.google.com/o/oauth2/token"
	entityName  string "allUsers"
	redirectURL string "urn:ietf:wg:oauth:2.0:oob"
}