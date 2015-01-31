// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package doc

import (
	"fmt"
	"github.com/google/cayley/config"
	"github.com/google/cayley/db"
	"github.com/google/cayley/graph"
	"sync"

	// Load all supported backends.
	_ "github.com/google/cayley/graph/bolt"
	_ "github.com/google/cayley/graph/leveldb"
	_ "github.com/google/cayley/graph/memstore"
	_ "github.com/google/cayley/graph/mongo"
	_ "github.com/google/cayley/writer"
)

type Graph struct {
	Config *config.Config
	Handle *graph.Handle
}

var err error
var create sync.Once

func (g Graph) Init() {

	create.Do(func() {

		if graph.IsPersistent(g.Config.DatabaseType) {
			err = db.Init(g.Config)
			if err != nil {
				fmt.Println("Could not initialize database: %v", err)
			}
		}

		g.Handle, err = db.Open(g.Config)

		if err != nil {
			fmt.Println("Failed to open %q: %v", g.Config.DatabasePath, err)
		}

	})
}
