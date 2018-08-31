// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"testing"
)

const (
	testBoostrapNodeCont = 2
)

func setupBoostrapTestNode(t *testing.T, nodeNo int) *Server {
	server := NewServer()
	server.SetName(fmt.Sprintf(testNodeNamePrefix, nodeNo))

	return server
}

func setupBoostrapTestNodes(t *testing.T) []*Server {
	servers := make([]*Server, testFederatedNodeCont)
	for n := 0; n < testFederatedNodeCont; n++ {
		servers[n] = setupFederatedTestNode(t, n)
	}
	return servers
}

func TestSeverBoostrapExport(t *testing.T) {
	node := setupBoostrapTestNode(t, 0)

	err := node.Start()
	if err != nil {
		t.Error(err)
	}

	_, err = node.exportMonitoringConfigurations()
	if err != nil {
		t.Error(err)
	}

	err = node.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestSeverBoostrap(t *testing.T) {
	nodes := setupBoostrapTestNodes(t)

	for _, node := range nodes {
		err := node.Start()
		if err != nil {
			t.Error(err)
		}
	}

	for _, node := range nodes {
		err := node.Stop()
		if err != nil {
			t.Error(err)
		}
	}
}
