// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const (
	testBoostrapNodeCont        = 2
	testBoostrapConfigQueryFile = "server_boostrap_test.fql"
)

func setupBoostrapTestNodeWithConfigFlag(t *testing.T, nodeNo int, configFlag bool) (*Server, error) {
	server := NewServer()

	// Set basic configurations

	conf := NewDefaultConfig()
	conf.Server.Host = fmt.Sprintf(testNodeNamePrefix, nodeNo)
	conf.Server.Boostrap = 1

	// Load monitoring configurations

	if configFlag {
		testDir, _ := os.Getwd()
		filename := filepath.Join(testDir, testBoostrapConfigQueryFile)
		err := server.LoadQuery(filename)
		if err != nil {
			return nil, err
		}
	}

	return server, nil
}

func setupBoostrapTestNode(t *testing.T, nodeNo int) (*Server, error) {
	return setupBoostrapTestNodeWithConfigFlag(t, nodeNo, true)
}

func setupBoostrapTestNodeWithoutConfig(t *testing.T, nodeNo int) (*Server, error) {
	return setupBoostrapTestNodeWithConfigFlag(t, nodeNo, false)
}

func setupBoostrapTestNodes(t *testing.T) ([]*Server, error) {
	servers := make([]*Server, testFederatedNodeCont)
	for n := 0; n < testFederatedNodeCont; n++ {
		server, err := setupBoostrapTestNode(t, n)
		if err != nil {
			return nil, err
		}
		servers[n] = server
	}
	return servers, nil
}

func TestSeverBoostrapConfigStaticTransport(t *testing.T) {
	// Export

	srcNode, err := setupBoostrapTestNode(t, 0)
	if err != nil {
		t.Error(err)
		return
	}

	srcConfig, err := srcNode.exportBoostrapConfig()
	if err != nil {
		t.Error(err)
	}

	// Import

	dstNode, err := setupBoostrapTestNodeWithoutConfig(t, 1)
	if err != nil {
		t.Error(err)
		return
	}

	err = dstNode.importBoostrapConfig(srcConfig)
	if err != nil {
		t.Error(err)
	}

	dstConfig, err := dstNode.exportBoostrapConfig()
	if err != nil {
		t.Error(err)
	}

	// Compare

	if !srcConfig.Equals(dstConfig) {
		t.Errorf("%v != %v", srcConfig, dstConfig)
	}
}

func TestSeverBoostrap(t *testing.T) {
	// Start source nodes with default configuration

	nodes, err := setupBoostrapTestNodes(t)
	if err != nil {
		t.Error(err)
		return
	}

	for _, node := range nodes {
		err := node.Start()
		if err != nil {
			t.Error(err)
		}
	}

	// Start a new node with boostrap option

	newNode, err := setupBoostrapTestNodeWithoutConfig(t, len(nodes)+1)
	if err != nil {
		t.Error(err)
		return
	}

	newNode.SetBoostrapEnabled(true)

	err = newNode.Start()
	if err != nil {
		t.Error(err)
		return
	}

	// Compare

	nodeConfig, err := nodes[0].exportBoostrapConfig()
	if err != nil {
		t.Error(err)
	}

	newConfig, err := newNode.exportBoostrapConfig()
	if err != nil {
		t.Error(err)
	}

	if !newConfig.Equals(nodeConfig) {
		t.Errorf("%v != %v", newConfig, nodeConfig)
	}

	// Finalize

	for _, node := range nodes {
		err := node.Stop()
		if err != nil {
			t.Error(err)
		}
	}

	err = newNode.Stop()
	if err != nil {
		t.Error(err)
		return
	}
}
