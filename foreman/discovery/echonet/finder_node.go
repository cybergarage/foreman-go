// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/node"

	uecho_protocol "github.com/cybergarage/uecho-go/net/echonet/protocol"
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
func NewFinderNodeWithResponseMesssage(msg *uecho_protocol.Message) (node.Node, error) {
	// Valdate the specified message

	if msg == nil {
		return nil, fmt.Errorf(errorEchonetFinderInvalidMessage, msg)
	}

	if msg.GetSourceObjectCode() != FinderDeviceCode {
		return nil, fmt.Errorf(errorEchonetFinderMessageInvalidObject, msg.GetSourceObjectCode(), FinderDeviceCode)
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

	for _, prop := range msg.GetProperties() {
		switch prop.GetCode() {
		case FinderConditionCode:
			candidateNode.Condition = node.Condition(prop.GetIntegerData())
		case FinderClusterCode:
			candidateNode.Cluster = prop.GetStringData()
		case FinderNameCode:
			candidateNode.Name = prop.GetStringData()
		case FinderAddressCode:
			candidateNode.Address = prop.GetStringData()
		case FinderRPCPortCode:
			candidateNode.RPCPort = int(prop.GetIntegerData())
		case FinderRenderPortCode:
			candidateNode.RenderPort = int(prop.GetIntegerData())
		case FinderCarbonPortCode:
			candidateNode.CarbonPort = int(prop.GetIntegerData())
		case FinderClockCode:
			candidateNode.Clock = node.Clock(prop.GetIntegerData())
		case FinderVersionCode:
			candidateNode.Version = node.Version(prop.GetIntegerData())
		default:
			return nil, fmt.Errorf(errorEchonetFinderMessageInvalidProperty, prop.GetCode())
		}
	}

	// FIXME : Why invalid messages of empty or loopback address are sent
	addr := candidateNode.GetAddress()
	if !uecho_transport.IsCommunicableAddress(addr) {
		return nil, fmt.Errorf(errorEchonetFinderInvalidNodeAddress, addr)
	}

	return candidateNode, nil
}
