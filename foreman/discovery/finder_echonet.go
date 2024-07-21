// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"fmt"
	"reflect"
	"time"

	foreman_echonet "github.com/cybergarage/foreman-go/foreman/discovery/echonet"
	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/cybergarage/foreman-go/foreman/node"

	uecho_echonet "github.com/cybergarage/uecho-go/net/echonet"
	uecho_protocol "github.com/cybergarage/uecho-go/net/echonet/protocol"
)

const (
	echonetFinderSearchSleepSecond = 1
)

const (
	errorEchonetFinderNoResponse     = "Echonet node (%s:%d) is not responding"
	msgEchonetFinderFoundEchonetNode = "Echonet node (%s:%d) is found"
	msgEchonetFinderFoundCadiateNode = "Candidate finder node (%s:%d) is found"
	msgEchonetFinderFoundNewNode     = "New finder node (%s:%d) is found"
)

// EchonetFinder represents a base finder.
type EchonetFinder struct {
	*baseFinder
	localNode node.Node
	*foreman_echonet.EchonetController
}

// NewEchonetFinderWithLocalNode returns a new finder with the specified node.
func NewEchonetFinderWithLocalNode(node node.Node) Finder {
	finder := &EchonetFinder{
		baseFinder:        newBaseFinder(),
		localNode:         node,
		EchonetController: foreman_echonet.NewController(),
	}
	finder.EchonetController.SetListener(finder)
	return finder
}

// NewEchonetFinder returns a new finder of Echonet.
func NewEchonetFinder() Finder {
	return NewEchonetFinderWithLocalNode(nil)
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

// IsLocalNode returns true when the specified node is the local node, otherwise false.
func (finder *EchonetFinder) IsLocalNode(candidateNode node.Node) bool {
	if finder.localNode == nil || reflect.ValueOf(finder.localNode).IsNil() {
		return false
	}

	return node.Equal(finder.localNode, candidateNode)
}

// Start starts the finder.
func (finder *EchonetFinder) Start() error {
	return finder.EchonetController.Start()
}

// Stop stops the finder.
func (finder *EchonetFinder) Stop() error {
	return finder.EchonetController.Stop()
}

// IsRunning returns true when the finder is running, otherwise false.
func (finder *EchonetFinder) IsRunning() bool {
	return finder.EchonetController.IsRunning()
}

// String returns the description
func (finder *EchonetFinder) String() string {
	return fmt.Sprintf("%s:%s", FinderEchonet, uecho_echonet.Version)
}

func (finder *EchonetFinder) ControllerMessageReceived(msg *uecho_protocol.Message) {
	if !msg.IsReadRequest() {
		return
	}

	finder.EchonetController.EchonetDevice.UpdatePropertyWithNode(finder.localNode)
}

func (finder *EchonetFinder) ControllerNewNodeFound(echonetNode *uecho_echonet.RemoteNode) {
	if !finder.IsRunning() {
		return
	}

	logging.Trace(msgEchonetFinderFoundEchonetNode, echonetNode.Address(), echonetNode.Port())

	reqMsg := foreman_echonet.NewRequestAllPropertiesMessage()
	resMsg, err := finder.PostMessage(echonetNode, reqMsg)
	if err != nil {
		logging.Error(errorEchonetFinderNoResponse, echonetNode.Address(), echonetNode.Port())
		logging.Error("%s", err.Error())
		return
	}

	candidateNode, err := foreman_echonet.NewFinderNodeWithResponseMesssage(resMsg)
	if err != nil {
		logging.Error("%s", err.Error())
		return
	}

	logging.Trace(msgEchonetFinderFoundCadiateNode, candidateNode.GetAddress(), candidateNode.GetRPCPort())

	if finder.IsLocalNode(candidateNode) {
		return
	}

	if finder.HasNode(candidateNode) {
		return
	}

	logging.Info(msgEchonetFinderFoundNewNode, candidateNode.GetAddress(), candidateNode.GetRPCPort())

	finder.addNode(candidateNode)
}
