// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"testing"
	"time"

	foreman_echonet "github.com/cybergarage/foreman-go/foreman/discovery/echonet"
	//echonet_log "github.com/cybergarage/uecho-go/net/echonet/log"
)

const (
	errorEchonetFinderIsNotRunning = "Finder is not running"
)

func setupTestEchonetFinderNodes() ([]*foreman_echonet.EchonetNode, error) {
	nodes := setupTestFinderNodes()
	echonetNodes := make([]*foreman_echonet.EchonetNode, len(nodes))
	for n, node := range nodes {
		echonetNode, err := foreman_echonet.NewEchonetNodeWithNode(node)
		if err != nil {
			return nil, err
		}
		echonetNodes[n] = echonetNode
	}
	return echonetNodes, nil
}

func finderEchonetTest(t *testing.T, finder Finder, nodes []*foreman_echonet.EchonetNode) {
	// Check all nodes

	_, err := finder.GetAllNodes()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestEchonetFinder(t *testing.T) {
	//echonet_log.SetSharedLogger(echonet_log.NewStdoutLogger(echonet_log.LevelTrace))

	nodes, err := setupTestEchonetFinderNodes()
	if err != nil {
		t.Error(err)
		return
	}

	for _, node := range nodes {
		err = node.Start()
		if err != nil {
			t.Error(err)
		}
	}

	finder := NewEchonetFinder()

	err = finder.Start()
	if err != nil {
		t.Error(err)
		return
	}

	err = finder.Search()
	if err != nil {
		t.Error(err)
		return
	}

	time.Sleep((500 * time.Millisecond) * time.Duration(len(nodes)))

	finderTest(t, finder)
	finderEchonetTest(t, finder, nodes)

	err = finder.Stop()
	if err != nil {
		t.Error(err)
	}

	for _, node := range nodes {
		err = node.Stop()
		if err != nil {
			t.Error(err)
			return
		}
	}
}
