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
	errorInvalidNodeCount = "Invalid node count %d != %d"
	errorInvalidVersion   = "Invalid version %d != %d"
)

var testDetectorNodeNames = []string{
	"org.cybergarage.foreman001",
	"org.cybergarage.foreman002",
	"org.cybergarage.foreman003",
}

type testDetector struct {
	Detector
	AddedNodeCount     int
	OutOfDateNodeCount int
}

func (detector *testDetector) FailureDetectorNodeAdded(node Node) {
	detector.AddedNodeCount++
}

func (detector *testDetector) FailureDetectorNodeRemoved(node Node) {
}

func (detector *testDetector) FailureDetectorNodeStatusChanged(node Node) {
}

func (detector *testDetector) FailureDetectorNodeOutOfDate(node Node) {
	detector.OutOfDateNodeCount++
}

func newTestDetectorWithDetector(targetDetector Detector) (*testDetector, []*node.BaseNode) {
	detector := &testDetector{
		Detector:           targetDetector,
		AddedNodeCount:     0,
		OutOfDateNodeCount: 0,
	}

	nodes := make([]Node, len(testDetectorNodeNames))
	baseNodes := make([]*node.BaseNode, len(testDetectorNodeNames))
	for n, name := range testDetectorNodeNames {
		baseNode := node.NewBaseNode()
		baseNode.Name = name
		baseNode.Condition = node.ConditionReady
		nodes[n] = baseNode
		baseNodes[n] = baseNode
	}

	finder := discovery.NewStaticFinderWithNodes(nodes)
	detector.SetFinder(finder)
	detector.SetListener(detector)

	return detector, baseNodes
}

func runTestDetector(t *testing.T, rawDetector interface{}, nodes []*node.BaseNode, waitBaseTime time.Duration) {
	detector := rawDetector.(Detector)
	testDetector := rawDetector.(*testDetector)

	// Start detector

	err := detector.Start()
	if err != nil {
		t.Error(err)
		detector.Stop()
	}

	// Check new node detections

	time.Sleep(waitBaseTime * 3)

	// Check new node detections

	nodes[0].UpdateVersion()
	time.Sleep(waitBaseTime * 3)

	if testDetector.OutOfDateNodeCount <= 0 {
		t.Errorf(errorInvalidNodeCount, testDetector.OutOfDateNodeCount, 1)
	}

	// Stop detector

	err = detector.Stop()
	if err != nil {
		t.Error(err)
	}

}

func checkDetectorStatus(t *testing.T, testDetector *testDetector, targetDetector *baseDetector, nodes []*node.BaseNode) {
	cluster := targetDetector.GetCluster()
	ver := cluster.GetVersion()
	if ver <= 0 {
		t.Errorf(errorInvalidVersion, ver, 1)
	}
}

func TestNewBroadcastDetector(t *testing.T) {
	simpleDetector := NewBroadcastDetector().(*BroadcastDetector)
	simpleDetector.SetDetectionInterval(time.Second * 1)
	simpleDetector.SetSuspectionDuration(0)
	detector, nodes := newTestDetectorWithDetector(simpleDetector)

	runTestDetector(t, detector, nodes, simpleDetector.GetDetectionInterval())

	if detector.AddedNodeCount != len(nodes) {
		t.Errorf(errorInvalidNodeCount, detector.AddedNodeCount, len(nodes))
	}

	checkDetectorStatus(t, detector, simpleDetector.baseDetector, nodes)
}

func TestNewGossipDetector(t *testing.T) {
	gossipDetector := NewGossipDetector().(*GossipDetector)
	gossipDetector.SetDetectionInterval(time.Second * 1)
	gossipDetector.SetSuspectionDuration(0)
	detector, nodes := newTestDetectorWithDetector(gossipDetector)

	runTestDetector(t, detector, nodes, gossipDetector.GetDetectionInterval())

	if detector.AddedNodeCount <= 0 {
		t.Errorf(errorInvalidNodeCount, detector.AddedNodeCount, 1)
	}

	checkDetectorStatus(t, detector, gossipDetector.baseDetector, nodes)
}
