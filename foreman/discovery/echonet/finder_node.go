// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import (
	"github.com/cybergarage/foreman-go/foreman/node"
	"github.com/cybergarage/uecho-go/net/echonet/protocol"
)

type finderNode struct {
	*node.BaseNode
}

// NewFinderNodeWithMesssage returns a new finder node with the specified message.
func NewFinderNodeWithMesssage(msg *protocol.Message) (node.Node, error) {
	node := &finderNode{
		BaseNode: node.NewBaseNode(),
	}

	return node, nil
}
