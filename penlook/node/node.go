// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package node

import (
	"errors"
)

type Node struct {
	Name string
}

func (node Node) Abc(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}

	return a / b, nil
}
