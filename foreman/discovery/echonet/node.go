// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import (
	"github.com/cybergarage/foreman-go/foreman/node"

	"github.com/cybergarage/uecho-go/net/echonet"
	"github.com/cybergarage/uecho-go/net/echonet/encoding"
	"github.com/cybergarage/uecho-go/net/echonet/protocol"
)

const (
	ManufacturerCode = echonet.NodeManufacturerUnknown
)

const (
	errorNodeNotFound = "Node Not Found : %s:%d"
)

type EchonetNode struct {
	*echonet.LocalNode
	sourceNode node.Node
}

// NewEchonetNodeWithNode returns a new finder node.
func NewEchonetNodeWithNode(srcNode node.Node) (*EchonetNode, error) {
	node := &EchonetNode{
		LocalNode:  echonet.NewLocalNode(),
		sourceNode: srcNode,
	}

	node.SetManufacturerCode(ManufacturerCode)
	node.SetListener(node)

	dev, err := NewDevice()
	if err != nil {
		return nil, err
	}

	err = node.AddDevice(dev)
	if err != nil {
		return nil, err
	}

	return node, nil
}

// GetNodeDevice returns the device which has the status and configuration.
func (node *EchonetNode) GetNodeDevice() (*echonet.Device, error) {
	return node.GetDevice(FinderDeviceCode)
}

// MessageReceived updates local properties for the source node.
func (node *EchonetNode) NodeMessageReceived(msg *protocol.Message) {
	if !msg.IsReadRequest() {
		return
	}

	srcNode := node.sourceNode
	if srcNode == nil {
		return
	}

	dev, err := node.GetNodeDevice()
	if err != nil {
		return
	}

	for _, propCode := range FinderDeviceAllPropertyCodes() {
		propData := []byte{}
		switch propCode {
		case FinderConditionCode:
			propData = []byte{byte(srcNode.GetCondition())}
		case FinderClusterCode:
			propData = []byte(srcNode.GetCluster())
		case FinderNameCode:
			propData = []byte(srcNode.GetName())
		case FinderAddressCode:
			propData = []byte(srcNode.GetAddress())
		case FinderRPCPortCode:
			propData := make([]byte, FinderRPCPortSize)
			encoding.IntegerToByte(uint(srcNode.GetRPCPort()), propData)
		case FinderRenderPortCode:
			propData := make([]byte, FinderRenderPortSize)
			encoding.IntegerToByte(uint(srcNode.GetRenderPort()), propData)
		case FinderCarbonPortCode:
			propData := make([]byte, FinderCarbonPortSize)
			encoding.IntegerToByte(uint(srcNode.GetCarbonPort()), propData)
		case FinderClockCode:
			propData := make([]byte, FinderClockSize)
			encoding.IntegerToByte(uint(srcNode.GetClock()), propData)
		case FinderVersionCode:
			propData := make([]byte, FinderVersionSize)
			encoding.IntegerToByte(uint(srcNode.GetVersion()), propData)
		default:
			continue
		}
		dev.SetPropertyData(echonet.PropertyCode(propCode), propData)
	}
}
