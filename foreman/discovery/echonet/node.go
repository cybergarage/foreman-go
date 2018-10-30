// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import (
	"reflect"

	"github.com/cybergarage/foreman-go/foreman/node"
	uecho_echonet "github.com/cybergarage/uecho-go/net/echonet"
	uecho_protocol "github.com/cybergarage/uecho-go/net/echonet/protocol"
)

const (
	errorNodeNotFound = "Node Not Found : %s:%d"
)

type EchonetNode struct {
	*uecho_echonet.LocalNode
	*EchonetDevice
	sourceNode node.Node
}

// NewEchonetNodeWithNode returns a new finder node.
func NewEchonetNodeWithNode(srcNode node.Node) (*EchonetNode, error) {
	node := &EchonetNode{
		LocalNode:     uecho_echonet.NewLocalNode(),
		EchonetDevice: NewDevice(),
		sourceNode:    srcNode,
	}

	node.SetConfig(NewDefaultConfig())
	node.SetManufacturerCode(ManufacturerCode)

	node.SetListener(node)

	err := node.AddDevice(node.EchonetDevice.Device)
	if err != nil {
		return nil, err
	}

	return node, nil
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
	if node.sourceNode == nil || reflect.ValueOf(node.sourceNode).IsNil() {
		return false
	}
	return true
}

// GetSourceNode returns the local souce node in the node.
func (node *EchonetNode) GetSourceNode() node.Node {
	return node.sourceNode
}

// NodeMessageReceived updates local properties for the source node.
func (node *EchonetNode) NodeMessageReceived(msg *uecho_protocol.Message) {
	if !msg.IsReadRequest() {
		return
	}

	if !node.HasSourceNode() {
		return
	}

	node.EchonetDevice.UpdatePropertyWithNode(node.GetSourceNode())
}
