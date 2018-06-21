// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package node

// Node represents an abstract node interface
type Node interface {
	Config
	Status
}

// Equal returns true if the other node is same with this node
func Equal(this, other Node) bool {
	if !ConfigEqual(this, other) {
		return false
	}
	return true
}
