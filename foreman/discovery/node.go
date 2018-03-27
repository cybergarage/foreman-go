// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"net"
)

// Node represents a node.
type Node struct {
	Cluster string
	Address net.Addr
}

// NewNode returns a new node.
func NewNode() *Node {
	node := &Node{}
	return node
}
