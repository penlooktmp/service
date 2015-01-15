// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package neo

import (
	"fmt"
	neo4j "github.com/penlook/neo4j"
)

type Neo struct {
	*neo4j
}

func (neo Neo) Connect() {
	connection, err := neo4j.Connect("http://localhost:7474/db/data")

	if err != nil {
		panic(err)
	}

	return connection
}
