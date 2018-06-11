// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

// Node represents an abstract node interface
type Node interface {
	NodeConfig
	NodeStatus
}

// NodeEqual returns true if the other node is same with this node
func NodeEqual(this, other Node) bool {
	if !NodeConfigEqual(this, other) {
		return false
	}
	return true
}
