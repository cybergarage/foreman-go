// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/cybergarage/foreman-go/foreman/logging"
)

const (
	testSharedNodeCont             = 3
	testSharedTestQosName          = "qos%03d"
	testSharedTestQosFormulaPrefix = "(m%03d <"
	testSharedTestQosSetQuery      = "SET (qos%03d, \"m%03d < %03d\") INTO QOS"
	testSharedTestQosGetQuery      = "EXPORT FROM QOS WHERE name == qos%03d"
)

func setupSharedTestNode(t *testing.T, nodeNo int) *Server {
	server := NewServer()
	server.SetName(fmt.Sprintf(testNodeNamePrefix, nodeNo))
	err := server.Start()
	if err != nil {
		t.Error(err)
	}

	return server
}

func setupSharedTestNodes(t *testing.T) []*Server {
	servers := make([]*Server, testSharedNodeCont)
	for n := 0; n < testSharedNodeCont; n++ {
		servers[n] = setupSharedTestNode(t, n)
	}
	return servers
}

func sharedRegistryTest(t *testing.T, client *Client, nodes []*Server) {
	testNodeCount := len(nodes)

	for _, node := range nodes {
		client.SetHTTPPort(node.GetHTTPPort())
		client.SetCarbonPort(node.GetCarbonPort())
		client.SetRenderPort(node.GetRenderPort())

		for n := 0; n < testNodeCount; n++ {
			q := fmt.Sprintf(testSharedTestQosGetQuery, n)
			resObj, err := client.PostQuery(q)
			if err != nil {
				t.Error(err)
				return
			}

			resMap, ok := resObj.(map[string]interface{})
			if !ok {
				t.Error("")
				return
			}
			qosObj, ok := resMap["qos"]
			if !ok {
				t.Errorf("%s : %s", node.GetName(), q)
				return
			}
			qosMap, ok := qosObj.(map[string]interface{})
			if !ok {
				t.Errorf("%s : %s", node.GetName(), q)
				return
			}

			qosName := fmt.Sprintf(testSharedTestQosName, n)
			resQosFormula, ok := qosMap[qosName]
			if !ok {
				t.Errorf("%s : %s", node.GetName(), q)
				return
			}
			resQosFormulaStr, ok := resQosFormula.(string)
			if !ok {
				t.Errorf("%s : %s", node.GetName(), q)
				return
			}

			//  (m000 < 0.000000) != (m000 < 000)

			qosFormulaPrefix := fmt.Sprintf(testSharedTestQosFormulaPrefix, n)
			if !strings.HasPrefix(resQosFormulaStr, qosFormulaPrefix) {
				t.Errorf("%s : %s != %s", qosName, resQosFormulaStr, qosFormulaPrefix)
				return
			}
		}
	}
}

/*
FIXME : Enable TestStandaloneSharedRegistryWithStaticFinder
func TestStandaloneSharedRegistryWithStaticFinder(t *testing.T) {
	node := setupSharedTestNode(t, 0)
	nodes := []*Server{node}
	finder := setupStaticFinderWithServers(t, nodes)
	node.AddFinder(finder)
	client := NewClient()
	client.AddFinder(finder)

	sharedRegistryTest(t, client, nodes)

	stopTestNodes(t, nodes)
}
*/

func TestMultiNodeSharedRegistryWithStaticFinder(t *testing.T) {
	logging.SetLogLevel(logging.LevelTrace)

	nodes := setupSharedTestNodes(t)
	finder := setupStaticFinderWithServers(t, nodes)

	for n, node := range nodes {
		node.AddFinder(finder)
		q := fmt.Sprintf(testSharedTestQosSetQuery, n, n, n)
		_, err := node.PostQuery(q)
		if err != nil {
			t.Error(err)
		}

		// Wait for asynchronous queries for other cluster nodes
		time.Sleep(1 * time.Second)
	}

	client := NewClient()
	client.AddFinder(finder)

	sharedRegistryTest(t, client, nodes)

	stopTestNodes(t, nodes)
}
