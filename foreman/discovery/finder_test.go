// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"regexp"
	"testing"

	"github.com/cybergarage/foreman-go/foreman/node"
)

const (
	testFinderNodeCountError     = "Node count error : %d != %d"
	testFinderMatchingError      = "Matching error (%s) : %s"
	testFinderMatchingCountError = "Matching count error (%s) : %d != %d"
)

var testFinderNodeNames = []string{
	"org.cybergarage.foreman001",
	"org.cybergarage.foreman002",
	"org.cybergarage.foreman003",
}

func setupTestFinderNodes() []Node {
	nodes := make([]Node, len(testFinderNodeNames))
	for n, name := range testFinderNodeNames {
		node := node.NewBaseNode()
		node.Name = name
		nodes[n] = node
	}
	return nodes
}

func finderTest(t *testing.T, finder Finder) {
	// Check all nodes

	nodes, err := finder.GetAllNodes()
	if err != nil {
		t.Error(err)
		return
	}

	if len(nodes) != len(testFinderNodeNames) {
		t.Errorf(testFinderNodeCountError, len(nodes), len(testFinderNodeNames))
		return
	}

	// Check regexp names for a node

	for _, nodeName := range testFinderNodeNames {
		re := regexp.MustCompile(nodeName)
		nodes, err := finder.GetRegexpNodes(re)
		if err != nil {
			t.Error(err)
			return
		}
		if len(nodes) != 1 {
			t.Errorf(testFinderNodeCountError, len(nodes), 1)
			return
		}
		if nodes[0].GetName() != nodeName {
			t.Errorf(testFinderMatchingError, nodeName, nodes[0].GetName())
			return
		}
	}

	// Check regexp names for all nodes

	reNames := []string{
		".*",
		"^org.cybergarage.foreman",
		"org.cybergarage.foreman.*",
		"org.cybergarage.foreman00[1-3]",
	}

	for _, reName := range reNames {
		re := regexp.MustCompile(reName)
		nodes, err := finder.GetRegexpNodes(re)
		if err != nil {
			t.Error(err)
			return
		}
		if len(nodes) != len(testFinderNodeNames) {
			t.Errorf(testFinderMatchingCountError, reName, len(nodes), len(testFinderNodeNames))
			return
		}
	}

	// Check start names for all nodes

	metricsNames := []string{
		"org.cybergarage.foreman001",
		"org.cybergarage.foreman001.m1",
		"org.cybergarage.foreman001.system.m1",
		"org.cybergarage.foreman002",
		"org.cybergarage.foreman002.m1",
		"org.cybergarage.foreman002.system.m1",
		"org.cybergarage.foreman003",
		"org.cybergarage.foreman003.m1",
		"org.cybergarage.foreman003.system.m1",
	}

	for _, metricsName := range metricsNames {
		nodes, err := finder.GetPrefixNodes(metricsName)
		if err != nil {
			t.Error(err)
			return
		}
		if len(nodes) != 1 {
			t.Errorf(testFinderMatchingCountError, metricsName, len(nodes), 1)
			return
		}
	}
}

func TestNewSharedFinder(t *testing.T) {
	nodes := setupTestFinderNodes()

	finder := NewSharedFinder().(*SharedFinder)
	for _, node := range nodes {
		finder.addNode(node)
	}

	err := finder.Start()
	if err != nil {
		t.Error(err)
		return
	}

	finderTest(t, finder)

	err = finder.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestNewStaticFinder(t *testing.T) {
	nodes := setupTestFinderNodes()

	finder := NewStaticFinderWithNodes(nodes)

	err := finder.Start()
	if err != nil {
		t.Error(err)
		return
	}

	finderTest(t, finder)

	err = finder.Stop()
	if err != nil {
		t.Error(err)
	}
}
