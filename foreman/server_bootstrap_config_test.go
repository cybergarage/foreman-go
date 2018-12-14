// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"os"
	"path/filepath"
	"testing"
)

const (
	testBootstrapNodeCont        = 2
	testBootstrapConfigQueryFile = "server_bootstrap_config_test.fql"
)

func setupBootstrapConfigTestNodeWithConfigFlag(t *testing.T, nodeNo int, configFlag bool) (*Server, error) {
	server := NewServer()

	// Load monitoring configurations

	if configFlag {
		testDir, _ := os.Getwd()
		filename := filepath.Join(testDir, testBootstrapConfigQueryFile)
		err := server.LoadQuery(filename)
		if err != nil {
			return nil, err
		}
	}

	return server, nil
}

func setupBootstrapConfigTestNode(t *testing.T, nodeNo int) (*Server, error) {
	return setupBootstrapConfigTestNodeWithConfigFlag(t, nodeNo, true)
}

func setupBootstrapConfigTestNodeWithoutConfig(t *testing.T, nodeNo int) (*Server, error) {
	return setupBootstrapConfigTestNodeWithConfigFlag(t, nodeNo, false)
}

func setupBootstrapConfigTestNodes(t *testing.T) ([]*Server, error) {
	servers := make([]*Server, testFederatedNodeCont)
	for n := 0; n < testFederatedNodeCont; n++ {
		server, err := setupBootstrapConfigTestNode(t, n)
		if err != nil {
			return nil, err
		}
		servers[n] = server
	}
	return servers, nil
}

func TestSeverBootstrapConfigStaticTransport(t *testing.T) {
	// Export

	srcNode, err := setupBootstrapConfigTestNode(t, 0)
	if err != nil {
		t.Error(err)
		return
	}

	srcConfig, err := srcNode.exportBootstrapConfig()
	if err != nil {
		t.Error(err)
	}

	// Import

	dstNode, err := setupBootstrapConfigTestNodeWithoutConfig(t, 1)
	if err != nil {
		t.Error(err)
		return
	}

	err = dstNode.importBootstrapConfig(srcConfig)
	if err != nil {
		t.Error(err)
	}

	dstConfig, err := dstNode.exportBootstrapConfig()
	if err != nil {
		t.Error(err)
	}

	// Compare

	if !srcConfig.Equals(dstConfig) {
		t.Errorf("%v != %v", srcConfig, dstConfig)
	}
}

func TestSeverBootstrapConfig(t *testing.T) {
	// Start source nodes with default configuration

	nodes, err := setupBootstrapConfigTestNodes(t)
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

	newNode, err := setupBootstrapConfigTestNodeWithoutConfig(t, len(nodes)+1)
	if err != nil {
		t.Error(err)
		return
	}

	newNode.SetBootstrapEnabled(true)

	err = newNode.Start()
	if err != nil {
		// FIXME : Update uecho-go to be able to the neighborhood node on CentOS
		t.Skip(err)
		return
	}

	// Compare the new node with other running nodes

	newConfig, err := newNode.exportBootstrapConfig()
	if err != nil {
		t.Error(err)
	}

	for _, node := range nodes {
		nodeConfig, err := node.exportBootstrapConfig()
		if err != nil {
			t.Error(err)
		}

		if !newConfig.Equals(nodeConfig) {
			t.Errorf("%v != %v", newConfig, nodeConfig)
		}
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
