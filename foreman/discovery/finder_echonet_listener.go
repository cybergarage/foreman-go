// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"github.com/cybergarage/uecho-go/net/echonet"
	"github.com/cybergarage/uecho-go/net/echonet/protocol"

	discovery_echonet "github.com/cybergarage/foreman-go/foreman/discovery/echonet"
)

func (finder *EchonetFinder) ControllerMessageReceived(*protocol.Message) {

}

func (finder *EchonetFinder) ControllerNewNodeFound(echonetNode *echonet.RemoteNode) {
	reqMsg := discovery_echonet.NewRequestAllPropertiesMessage()
	resMsg, err := finder.PostMessage(echonetNode, reqMsg)
	if err != nil {
		return
	}

	candidateNode, err := discovery_echonet.NewFinderNodeWithResponseMesssage(resMsg)
	if err != nil {
		return
	}

	if finder.HasNode(candidateNode) {
		return
	}

	finder.addNode(candidateNode)
}
