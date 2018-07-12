// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"testing"

	"github.com/cybergarage/foreman-go/foreman/discovery"
)

const (
	testNodeNamePrefix = "node%03d"
)

func setupStaticFinderWithServers(t *testing.T, servers []*Server) discovery.Finder {
	nodes := make([]discovery.Node, len(servers))

	for n, server := range servers {
		node := discovery.Node(server)
		if node == nil {
			continue
		}
		nodes[n] = node
	}
	return discovery.NewStaticFinderWithNodes(nodes)
}

func stopTestNodes(t *testing.T, servers []*Server) {
	for _, server := range servers {
		err := server.Stop()
		if err != nil {
			t.Error(err)
		}
	}
}

func TestNewServer(t *testing.T) {
	server := NewServer()

	err := server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	err = server.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestNewServerEquals(t *testing.T) {
	server := NewServer()
	remoteNode := NewRemoteNodeWithNode(server)

	if !NodeEqual(server, remoteNode) {
		t.Errorf("%s != %s", server.GetName(), remoteNode.GetName())
	}
}
