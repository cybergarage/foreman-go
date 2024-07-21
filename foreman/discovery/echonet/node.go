// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import (
	"fmt"
	"reflect"

	"github.com/cybergarage/foreman-go/foreman/node"
	uecho_echonet "github.com/cybergarage/uecho-go/net/echonet"
	uecho_protocol "github.com/cybergarage/uecho-go/net/echonet/protocol"
)

const (
	errorNodeNotRunning = "Node is not running"
	errorNodeNotFound   = "Node Not Found : %s:%d"
)

type EchonetNode struct {
	*uecho_echonet.LocalNode
	*EchonetDevice
	node.Node
}

// NewEchonetNodeWithNode returns a new finder node.
func NewEchonetNodeWithNode(srcNode node.Node) (*EchonetNode, error) {
	node := &EchonetNode{
		LocalNode:     uecho_echonet.NewLocalNode(),
		EchonetDevice: NewDevice(),
		Node:          srcNode,
	}

	node.SetConfig(NewDefaultConfig())
	node.SetManufacturerCode(ManufacturerCode)
	node.AddDevice(node.EchonetDevice.Device)
	node.SetListener(node)

	return node, nil
}

// GetAddress returns the interface address
func (node *EchonetNode) GetAddress() string {
	addrs := node.LocalNode.Addresses()
	if len(addrs) <= 0 {
		return ""
	}
	return addrs[0]
}

// GetLocalNode returns the local echonet node in the node
func (node *EchonetNode) GetLocalNode() *uecho_echonet.LocalNode {
	return node.LocalNode
}

// GetLocalDevice returns the local echonet device in the node.
func (node *EchonetNode) GetLocalDevice() *uecho_echonet.Device {
	return node.EchonetDevice.Device
}

// HasSourceNode returns true when this node has a source node, otherwise false.
func (node *EchonetNode) HasSourceNode() bool {
	if node.Node == nil || reflect.ValueOf(node.Node).IsNil() {
		return false
	}
	return true
}

// GetSourceNode returns the local souce node in the node.
func (node *EchonetNode) GetSourceNode() node.Node {
	return node.Node
}

// NodeMessageReceived updates local properties for the source node.
func (node *EchonetNode) NodeMessageReceived(msg *uecho_protocol.Message) error {
	if !node.IsRunning() {
		return fmt.Errorf(errorNodeNotRunning)
	}

	if !msg.IsReadRequest() {
		return nil
	}

	if !node.HasSourceNode() {
		return nil
	}

	node.EchonetDevice.UpdatePropertyWithNode(node)

	return nil
}
