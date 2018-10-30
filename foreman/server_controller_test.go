// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"testing"
	"time"
)

const (
	testServerControllerCont = 3
)

func setupControllerTestServer(t *testing.T, nodeNo int) *Server {
	server := NewServer()
	server.SetName(fmt.Sprintf(testNodeNamePrefix, nodeNo))
	return server
}

func setupControllerTestServers(t *testing.T) []*Server {
	servers := make([]*Server, testServerControllerCont)
	for n := 0; n < testServerControllerCont; n++ {
		servers[n] = setupControllerTestServer(t, n)
	}
	return servers
}

func testServerController(t *testing.T, server *Server) {
	ctrl := server.GetController()

	ctrl.Search()
	time.Sleep(time.Second * 1)

	nodes, err := ctrl.GetAllNodes()
	if err != nil {
		t.Error(err)
	}

	for n, node := range nodes {
		t.Logf("[%d] %s:%d", n, node.GetAddress(), node.GetRPCPort())
	}

	if len(nodes) != (testServerControllerCont - 1) {
		t.Errorf("%d != %d", len(nodes), (testServerControllerCont - 1))
	}
}

func TestServerControllerDefaultFinders(t *testing.T) {
	servers := setupControllerTestServers(t)

	for _, server := range servers {
		err := server.Start()
		if err != nil {
			t.Error(err)
		}
	}

	for n, server := range servers {
		fmt.Printf("[%d] %s:%d\n", n, server.GetAddress(), server.GetRPCPort())
		t.Logf("[%d] %s:%d", n, server.GetAddress(), server.GetRPCPort())
	}

	for _, server := range servers {
		testServerController(t, server)
	}

	for _, server := range servers {
		err := server.Stop()
		if err != nil {
			t.Error(err)
		}
	}
}
