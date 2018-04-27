// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

// Node represents a node.
type baseNode struct {
	Node
	Cluster string
	Address string
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
func (node *baseNode) GetAddress() string {
	return node.Address
}

// GetRPCPort returns the RPC port
func (node *baseNode) GetRPCPort() int {
	return node.RPCPort
}
