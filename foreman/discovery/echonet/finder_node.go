// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/node"

	uecho_echonet "github.com/cybergarage/uecho-go/net/echonet"
	uecho_transport "github.com/cybergarage/uecho-go/net/echonet/transport"
)

const (
	errorEchonetFinderInvalidNodeAddress     = "Invalid Echonet node address : '%s'"
	errorEchonetFinderInvalidMessage         = "Invalid Echonet message : %s"
	errorEchonetFinderMessageInvalidObject   = "Invalid Echonet object code : %X != %X"
	errorEchonetFinderMessageInvalidOPC      = "Invalid Echonet OPC : %d != %d"
	errorEchonetFinderMessageInvalidProperty = "Invalid Echonet property code : %X"
)

type finderNode struct {
	*node.BaseNode
}

// NewFinderNodeWithResponseMesssage returns a new finder node with the specified message.
func NewFinderNodeWithResponseMesssage(msg *uecho_echonet.Message) (node.Node, error) {
	// Valdate the specified message

	if msg == nil {
		return nil, fmt.Errorf(errorEchonetFinderInvalidMessage, msg)
	}

	if msg.SEOJ() != FinderDeviceCode {
		return nil, fmt.Errorf(errorEchonetFinderMessageInvalidObject, msg.SEOJ(), FinderDeviceCode)
	}

	for _, propCode := range FinderDeviceAllPropertyCodes() {
		if !msg.HasProperty(propCode) {
			return nil, fmt.Errorf(errorEchonetFinderInvalidMessage, msg)
		}
	}

	// Create a candidate from the specified message

	candidateNode := &finderNode{
		BaseNode: node.NewBaseNode(),
	}

	for _, prop := range msg.Properties() {
		switch prop.Code() {
		case FinderConditionCode:
			candidateNode.Condition = node.Condition(prop.IntegerData())
		case FinderClusterCode:
			candidateNode.Cluster = prop.StringData()
		case FinderNameCode:
			candidateNode.Name = prop.StringData()
		case FinderAddressCode:
			candidateNode.Address = prop.StringData()
		case FinderRPCPortCode:
			candidateNode.RPCPort = int(prop.IntegerData())
		case FinderRenderPortCode:
			candidateNode.RenderPort = int(prop.IntegerData())
		case FinderCarbonPortCode:
			candidateNode.CarbonPort = int(prop.IntegerData())
		case FinderClockCode:
			candidateNode.Clock = node.Clock(prop.IntegerData())
		case FinderVersionCode:
			candidateNode.Version = node.Version(prop.IntegerData())
		default:
			return nil, fmt.Errorf(errorEchonetFinderMessageInvalidProperty, prop.Code())
		}
	}

	// FIXME : Why invalid messages of empty or loopback address are sent
	addr := candidateNode.GetAddress()
	if !uecho_transport.IsCommunicableAddress(addr) {
		return nil, fmt.Errorf(errorEchonetFinderInvalidNodeAddress, addr)
	}

	return candidateNode, nil
}
