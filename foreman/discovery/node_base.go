// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import "net"

// Node represents a node.
type baseNode struct {
	Node
	Cluster    string
	Name       string
	Address    string
	RPCPort    int
	RenderPort int
	CarbonPort int
	Status     *NodeStatus
}

// NewNode returns a new node.
func newBaseNode() Node {
	node := &baseNode{}
	return node
}

// GetCluster returns the cluster name
func (node *baseNode) GetCluster() string {
	return node.Cluster
}

// GetName returns the host name
func (node *baseNode) GetName() string {
	if 0 < len(node.Name) {
		return node.Name
	}

	if len(node.Address) <= 0 {
		return ""
	}
	names, err := net.LookupAddr(node.Address)
	if err != nil {
		return ""
	}
	node.Name = names[0]

	return node.Name
}

// GetAddress returns the interface address
func (node *baseNode) GetAddress() string {
	if 0 < len(node.Address) {
		return node.Address
	}

	if len(node.Name) <= 0 {
		return ""
	}
	addrs, err := net.LookupIP(node.Name)
	if err != nil {
		return ""
	}
	node.Address = addrs[0].String()

	return node.Address
}

// GetRPCPort returns the RPC port
func (node *baseNode) GetRPCPort() int {
	return node.RPCPort
}

// GetRenderPort returns the Graphite render port
func (node *baseNode) GetRenderPort() int {
	return node.RenderPort
}

// GetCarbonPort returns the Graphite carbon port
func (node *baseNode) GetCarbonPort() int {
	return node.CarbonPort
}

// GetStatus returns the current condition status
func (node *baseNode) GetStatus() NodeStatus {
	return *node.Status
}
