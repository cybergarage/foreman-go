// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"time"

	foreman_echonet "github.com/cybergarage/foreman-go/foreman/discovery/echonet"
	"github.com/cybergarage/foreman-go/foreman/node"
	"github.com/cybergarage/uecho-go/net/echonet"
	"github.com/cybergarage/uecho-go/net/echonet/protocol"
)

const (
	echonetFinderSearchSleepSecond = 1
)

// EchonetFinder represents a base finder.
type EchonetFinder struct {
	node.Node
	echonet.ControllerListener
	*baseFinder
	*foreman_echonet.EchonetController
}

// NewEchonetFinderWithNode returns a new finder with the specified node.
func NewEchonetFinderWithNode(node node.Node) Finder {
	finder := &EchonetFinder{
		Node:              node,
		baseFinder:        newBaseFinder(),
		EchonetController: foreman_echonet.NewController(),
	}
	finder.EchonetController.SetListener(finder)
	return finder
}

// NewEchonetFinder returns a new finder of Echonet.
func NewEchonetFinder() Finder {
	return NewEchonetFinderWithNode(nil)
}

// Search searches all nodes.
func (finder *EchonetFinder) Search() error {
	err := finder.EchonetController.SearchAllObjects()
	if err != nil {
		return err
	}
	time.Sleep(time.Second * echonetFinderSearchSleepSecond)
	return nil
}

// Start starts the finder.
func (finder *EchonetFinder) Start() error {
	return finder.EchonetController.Start()
}

// Stop stops the finder.
func (finder *EchonetFinder) Stop() error {
	return finder.EchonetController.Stop()
}

func (finder *EchonetFinder) ControllerMessageReceived(msg *protocol.Message) {
	if !msg.IsReadRequest() {
		return
	}

	finder.EchonetController.EchonetDevice.UpdatePropertyWithNode(finder.Node)
}

func (finder *EchonetFinder) ControllerNewNodeFound(echonetNode *echonet.RemoteNode) {
	reqMsg := foreman_echonet.NewRequestAllPropertiesMessage()
	resMsg, err := finder.PostMessage(echonetNode, reqMsg)
	if err != nil {
		return
	}

	candidateNode, err := foreman_echonet.NewFinderNodeWithResponseMesssage(resMsg)
	if err != nil {
		return
	}

	if finder.HasNode(candidateNode) {
		return
	}

	finder.addNode(candidateNode)
}
