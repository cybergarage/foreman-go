// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

import (
	"testing"

	"github.com/cybergarage/foreman-go/foreman/discovery"
	"github.com/cybergarage/foreman-go/foreman/node"
)

var testDetectorNodeNames = []string{
	"org.cybergarage.foreman001",
	"org.cybergarage.foreman002",
	"org.cybergarage.foreman003",
}

type testDetector struct {
	Detector
}

func newTestDetectorWithDetector(targetDetector Detector) (Detector, []Node) {
	detector := &testDetector{
		Detector: targetDetector,
	}

	nodes := make([]Node, len(testDetectorNodeNames))
	for n, name := range testDetectorNodeNames {
		baseNode := node.NewBaseNode()
		baseNode.Name = name
		baseNode.Condition = node.ConditionReady
		nodes[n] = baseNode
	}

	finder := discovery.NewStaticFinderWithNodes(nodes)
	detector.SetFinder(finder)

	return detector, nodes
}

func detectorTest(t *testing.T, detector Detector, nodes []Node) {
	err := detector.Start()
	if err != nil {
		t.Error(err)
		detector.Stop()
	}

	err = detector.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestNewGossipDetector(t *testing.T) {
	gossipDetector := NewGossipDetector()
	detector, nodes := newTestDetectorWithDetector(gossipDetector)
	detectorTest(t, detector, nodes)
}
