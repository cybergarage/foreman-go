// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import (
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

// GetNodeDevice returns the device which has the status and configuration.
func (node *EchonetNode) GetNodeDevice() (*uecho_echonet.Device, error) {
	return node.GetDevice(FinderDeviceCode)
}

// NodeMessageReceived updates local properties for the source node.
func (node *EchonetNode) NodeMessageReceived(msg *uecho_protocol.Message) {
	if !msg.IsReadRequest() {
		return
	}

	srcNode := node.sourceNode
	if srcNode == nil {
		return
	}

	node.EchonetDevice.UpdatePropertyWithNode(srcNode)
}
