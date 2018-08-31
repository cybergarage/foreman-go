// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"testing"
	"time"

	"github.com/cybergarage/foreman-go/foreman/metric"
)

const (
	testFederatedNodeCont              = 3
	testFederatedMetricsPrefix         = "node%03d.metric%03d"
	testFederatedQueryPrefix           = "*.metric%03d"
	testFederatedMetricsStartTimestamp = 1514732400 // 2018-01-01 00:00:00
	testFederatedMetricsEndTimestamp   = 1514746800 // 2018-01-01 04:00:00
	testFederatedMetricsInterval       = 300
	testFederatedMetricsCount          = (testFederatedMetricsEndTimestamp - testFederatedMetricsStartTimestamp) / testFederatedMetricsInterval
)

func setupFederatedTestNode(t *testing.T, nodeNo int) *Server {
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

func setupFederatedTestNodes(t *testing.T) []*Server {
	servers := make([]*Server, testFederatedNodeCont)
	for n := 0; n < testFederatedNodeCont; n++ {
		servers[n] = setupFederatedTestNode(t, n)
	}
	return servers
}

func federatedMetricsTest(t *testing.T, client *Client, nodes []*Server) {
	testNodeCount := len(nodes)

	q := metric.NewMetricQuery()
	q.SetFromUnix(testFederatedMetricsStartTimestamp)
	q.SetUntilUnix(testFederatedMetricsEndTimestamp)
	q.SetIntervalSecond(testFederatedMetricsInterval)

	for _, node := range nodes {
		client.SetHTTPPort(node.GetHTTPPort())
		client.SetCarbonPort(node.GetCarbonPort())
		client.SetRenderPort(node.GetRenderPort())

		for n := 0; n < testFederatedMetricsCount; n++ {
			q.Target = fmt.Sprintf(testFederatedQueryPrefix, n)

			rs, err := client.GetMetrics(q)
			if err != nil {
				t.Error(err)
				return
			}

			if rs.GetMetricsCount() != testNodeCount {
				t.Errorf("%s : %d != %d", node.GetName(), rs.GetMetricsCount(), testNodeCount)
				break
			}

			ms := rs.GetFirstMetrics()
			for ms != nil {
				if len(ms.Values) != testFederatedMetricsCount {
					t.Errorf("%s : %d != %d", node.GetName(), len(ms.Values), testFederatedMetricsCount)
					break
				}
				ms = rs.GetNextMetrics()
			}
		}
	}
}

func TestStandaloneNodeMetricsWithStaticFinder(t *testing.T) {
	node := setupFederatedTestNode(t, 0)
	nodes := []*Server{node}
	finder := setupStaticFinderWithServers(t, nodes)
	node.AddFinder(finder)

	client := NewClient()
	client.AddFinder(finder)

	federatedMetricsTest(t, client, nodes)

	stopTestNodes(t, nodes)
}

func TestFederatedMultiNodeMetricsWithStaticFinder(t *testing.T) {
	nodes := setupFederatedTestNodes(t)
	finder := setupStaticFinderWithServers(t, nodes)

	for _, node := range nodes {
		node.AddFinder(finder)
	}

	client := NewClient()
	client.AddFinder(finder)

	federatedMetricsTest(t, client, nodes)

	stopTestNodes(t, nodes)
}
