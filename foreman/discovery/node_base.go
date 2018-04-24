// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"net"
)

// Node represents a node.
type baseNode struct {
	Node
	Cluster string
	Address net.Addr
	RPCPort int
}

// NewNode returns a new node.
func newBaseNode() Node {
	node := &baseNode{}
	return node
}

// GetCuster returns the cluster name
func (node *baseNode) GetCuster() string {
	return node.Cluster
}

// GetAddress returns the interface address
func (node *baseNode) GetAddress() net.Addr {
	return node.Address
}

// GetRPCPort returns the RPC port
func (node *baseNode) GetRPCPort() int {
	return node.RPCPort
}
