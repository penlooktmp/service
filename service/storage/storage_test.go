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
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var s3 = S3{}

func TestS3Config(t *testing.T) {
	assert := assert.New(t)
	// s3.GetConfig()
	assert.NotNil(os.Getenv("AWS_S3_KEY"))
	assert.NotNil(os.Getenv("AWS_S3_SECRET"))

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
