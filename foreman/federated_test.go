// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"testing"

	"github.com/cybergarage/foreman-go/foreman/discovery"
)

const (
	testFederatedNodePrefix     = "node%02d"
	testFederatedNodeCont       = 10
	testFederatedStartTimestamp = 1514732400 // 2018-01-01 00:00:00
	testFederatedEndTimestamp   = 1514775600 // 2018-01-01 12:00:00
	testFederatedInterval       = 300
)

func federatedMetricsTest(t *testing.T, client *Client, nodes []Node) {
}

func setupFederatedNodes(t *testing.T) []*Server {
	servers := make([]*Server, testFederatedNodeCont)

	for n := 0; n < testFederatedNodeCont; n++ {
		server := NewServer()
		err := server.Start()
		if err != nil {
			t.Error(err)
		}
		servers[n] = server
	}
	return servers
}

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

func stopFederatedNodes(t *testing.T, servers []*Server) {
	for _, server := range servers {
		err := server.Stop()
		if err != nil {
			t.Error(err)
		}
	}
}

func TestFederatedMetricsWithStaticFinder(t *testing.T) {
	nodes := setupFederatedNodes(t)
	finder := setupStaticFinderWithServers(t, nodes)

	client := NewClient()
	client.AddFinder(finder)

	//federatedMetricsTest(t, client, nodes)

	stopFederatedNodes(t, nodes)
}
