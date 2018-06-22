// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package node

import "net"

// BaseNode represents a base node.
type BaseNode struct {
	Node
	Cluster    string
	Name       string
	Address    string
	RPCPort    int
	RenderPort int
	CarbonPort int
	Clock      Clock
	Condition  Condition
	Version    Version
}

// NewBaseNode returns a new base node.
func NewBaseNode() *BaseNode {
	node := &BaseNode{
		Condition: ConditionInitial,
		Clock:     0,
		Version:   0,
	}
	return node
}

// NewNode returns a new node.
func NewNode() Node {
	return NewBaseNode()
}

// GetCluster returns the cluster name
func (node *BaseNode) GetCluster() string {
	return node.Cluster
}

// GetName returns the host name
func (node *BaseNode) GetName() string {
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
func (node *BaseNode) GetAddress() string {
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
func (node *BaseNode) GetRPCPort() int {
	return node.RPCPort
}

// GetRenderPort returns the Graphite render port
func (node *BaseNode) GetRenderPort() int {
	return node.RenderPort
}

// GetCarbonPort returns the Graphite carbon port
func (node *BaseNode) GetCarbonPort() int {
	return node.CarbonPort
}

// GetCondition returns the current status
func (node *BaseNode) GetCondition() Condition {
	return node.Condition
}

// GetClock returns the current logical clock
func (node *BaseNode) GetClock() Clock {
	return node.Clock
}

// GetVersion returns the current repository version
func (node *BaseNode) GetVersion() Version {
	return node.Version
}
