// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

import (
	"testing"
	"time"

	"github.com/cybergarage/foreman-go/foreman/discovery"
	"github.com/cybergarage/foreman-go/foreman/node"
)

const (
	errorInvalidNewNodeCount = "Invalid new node count %d != %d"
)

var testDetectorNodeNames = []string{
	"org.cybergarage.foreman001",
	"org.cybergarage.foreman002",
	"org.cybergarage.foreman003",
}

type testDetector struct {
	Detector
	AddedNodeCount int
}

func (detector *testDetector) FailureDetectorNodeAdded(node Node) {
	detector.AddedNodeCount++
}

func (detector *testDetector) FailureDetectorNodeRemoved(node Node) {
}

func (detector *testDetector) FailureDetectorNodeStatusChanged(node Node) {
}

func (detector *testDetector) FailureDetectorNodeOutOfDate(node Node) {
}

func newTestDetectorWithDetector(targetDetector Detector) (Detector, []Node) {
	detector := &testDetector{
		Detector:       targetDetector,
		AddedNodeCount: 0,
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
	detector.SetListener(detector)

	return detector, nodes
}

func runTestDetector(t *testing.T, rawDetector interface{}, nodes []Node) {
	detector := rawDetector.(Detector)

	err := detector.Start()
	if err != nil {
		t.Error(err)
		detector.Stop()
	}

	time.Sleep(HeartbeatDetectorDefaultIntervalTime * 5)

	err = detector.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestNewBroadcastDetector(t *testing.T) {
	simpleDetector := NewBroadcastDetector()
	detector, nodes := newTestDetectorWithDetector(simpleDetector)

	runTestDetector(t, detector, nodes)

	testDetector := detector.(*testDetector)
	if testDetector.AddedNodeCount != len(nodes) {
		t.Errorf(errorInvalidNewNodeCount, testDetector.AddedNodeCount, len(nodes))
	}
}

func TestNewGossipDetector(t *testing.T) {
	gossipDetector := NewGossipDetector()
	detector, nodes := newTestDetectorWithDetector(gossipDetector)
	runTestDetector(t, detector, nodes)

	testDetector := detector.(*testDetector)
	if testDetector.AddedNodeCount <= 0 {
		t.Errorf(errorInvalidNewNodeCount, testDetector.AddedNodeCount, 1)
	}
}
