// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package cql

type Cql struct {
	Name     string
	Server   string
}

func (cql Cql) Connect() Cql {
	return cql
}
