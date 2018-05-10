// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"testing"
	"time"

	"github.com/cybergarage/foreman-go/foreman/discovery"
	"github.com/cybergarage/foreman-go/foreman/metric"
)

const (
	testNodeNamePrefix = "node%03d"
)

func setupTestNode(t *testing.T, nodeNo int) *Server {
	server := NewServer()
	server.SetName(fmt.Sprintf(testNodeNamePrefix, nodeNo))
	err := server.Start()
	if err != nil {
		t.Error(err)
	}

	ts := testFederatedMetricsStartTimestamp
	for n := 0; n < testFederatedMetricsCount; n++ {
		m := metric.NewMetric()
		m.Name = fmt.Sprintf(testFederatedMetricsPrefix, nodeNo, n)
		m.Timestamp = time.Unix(int64(ts), 0)
		m.Value = float64(n)
		ts += testFederatedMetricsInterval
		err := server.metricMgr.AddMetric(m)
		if err != nil {
			t.Error(err)
			break
		}
	}

	return server
}

func setupTestNodes(t *testing.T) []*Server {
	servers := make([]*Server, testFederatedNodeCont)
	for n := 0; n < testFederatedNodeCont; n++ {
		servers[n] = setupTestNode(t, n)
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
	remoteNode := NewRemoteNodeWithDiscoveryNode(server)

	if !NodeEqual(server, remoteNode) {
		t.Errorf("%s != %s", server.GetName(), remoteNode.GetName())
	}
}
