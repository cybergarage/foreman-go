// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"testing"

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
	node := setupTestNode(t, 0)
	nodes := []*Server{node}
	finder := setupStaticFinderWithServers(t, nodes)
	node.AddFinder(finder)

	client := NewClient()
	client.AddFinder(finder)

	federatedMetricsTest(t, client, nodes)

	stopTestNodes(t, nodes)
}

func TestFederatedMultiNodeMetricsWithStaticFinder(t *testing.T) {
	nodes := setupTestNodes(t)
	finder := setupStaticFinderWithServers(t, nodes)

	for _, node := range nodes {
		node.AddFinder(finder)
	}

	client := NewClient()
	client.AddFinder(finder)

	federatedMetricsTest(t, client, nodes)

	stopTestNodes(t, nodes)
}
