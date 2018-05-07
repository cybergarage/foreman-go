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
	testFederatedNodePrefix            = "node%03d"
	testFederatedNodeCont              = 10
	testFederatedMetricsPrefix         = "node%03d.metric%03d"
	testFederatedMetricsStartTimestamp = 1514732400 // 2018-01-01 00:00:00
	testFederatedMetricsEndTimestamp   = 1514775600 // 2018-01-01 12:00:00
	testFederatedMetricsInterval       = 300
	testFederatedMetricsCount          = (testFederatedMetricsEndTimestamp - testFederatedMetricsStartTimestamp) / testFederatedMetricsInterval
)

func setupFederatedNode(t *testing.T, nodeNo int) *Server {
	server := NewServer()

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

	err := server.Start()
	if err != nil {
		t.Error(err)
	}

	return server
}

func setupFederatedNodes(t *testing.T) []*Server {
	servers := make([]*Server, testFederatedNodeCont)
	for n := 0; n < testFederatedNodeCont; n++ {
		servers[n] = setupFederatedNode(t, n)
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

func federatedMetricsTest(t *testing.T, client *Client, nodes []*Server) {
}

func TestFederatedMetricsWithStaticFinder(t *testing.T) {
	nodes := setupFederatedNodes(t)
	finder := setupStaticFinderWithServers(t, nodes)

	client := NewClient()
	client.AddFinder(finder)

	federatedMetricsTest(t, client, nodes)

	stopFederatedNodes(t, nodes)
}
