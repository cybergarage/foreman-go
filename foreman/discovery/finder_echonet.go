// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	discovery_echonet "github.com/cybergarage/foreman-go/foreman/discovery/echonet"
	"github.com/cybergarage/foreman-go/foreman/node"

	"github.com/cybergarage/uecho-go/net/echonet"
)

// EchonetFinder represents a base finder.
type EchonetFinder struct {
	*baseFinder
	*echonet.Controller
}

// NewEchonetFinder returns a new finder of Echonet.
func NewEchonetFinder() Finder {
	finder := &EchonetFinder{
		baseFinder: newBaseFinder(),
		Controller: echonet.NewController(),
	}
	return finder
}

// SearchAll searches all nodes.
func (finder *EchonetFinder) SearchAll() error {
	return nil
}

// Start starts the finder.
func (finder *EchonetFinder) Start() error {
	return nil
}

// Stop stops the finder.
func (finder *EchonetFinder) Stop() error {
	return nil
}

// NewFinderNodeWithNode returns a new finder node with the specified node.
func (finder *EchonetFinder) createFinderNodeWithNode(echonetNode *echonet.RemoteNode) (node.Node, error) {
	reqMsg := discovery_echonet.NewRequestAllPropertiesMessage()
	resMsg, err := finder.PostMessage(echonetNode, reqMsg)
	if err != nil {
		return nil, err
	}

	node, err := discovery_echonet.NewFinderNodeWithMesssage(resMsg)
	if err != nil {
		return nil, err
	}

	return node, nil
}
