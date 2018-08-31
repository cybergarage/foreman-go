// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

// BroadcastDetector represents a simple failure detector which check all nodes everytime.
type BroadcastDetector struct {
	*HeartbeatDetector
}

// NewBroadcastDetector returns a new gossip detector.
func NewBroadcastDetector() Detector {
	detector := &BroadcastDetector{
		HeartbeatDetector: NewHeartbeatDetector(),
	}
	detector.SetExecutor(detector)
	return detector
}

// ExecuteFailureDetection runs the failure action
func (detector *BroadcastDetector) ExecuteFailureDetection(Detector) error {
	finder, err := detector.GetFinder()
	if err != nil {
		return err
	}

	targetNodes, err := finder.GetAllNodes()
	if err != nil {
		return err
	}

	targetAllNodeCount := len(targetNodes)
	for n := 0; n < targetAllNodeCount; n++ {
		targetNode := targetNodes[n]
		detector.ExecuteNodeFailureDetection(targetNode)
	}

	return nil
}
