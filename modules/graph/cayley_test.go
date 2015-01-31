// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package doc

import (
	"github.com/google/cayley/config"
	//"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGraph(t *testing.T) {

	graph := Graph{
		Config: &config.Config{
			ReplicationType: "single",
			Timeout:         300 * time.Second,
			DatabaseType:    "mongo",
			DatabasePath:    "127.0.0.1:27017/cayley",
			LoadSize:        1000,
		},
	}

	graph.Init()
}
