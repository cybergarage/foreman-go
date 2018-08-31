// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

import (
	"math"
	"math/rand"
	"time"
)

var GossipExecutorDefaultPercentage = 1.0 / float64(((time.Second * 60) / HeartbeatDetectorDefaultIntervalTime))

// GossipDetector represents a gossip based detector.
type GossipDetector struct {
	*HeartbeatDetector
}

// NewGossipDetector returns a new gossip detector.
func NewGossipDetector() Detector {
	detector := &GossipDetector{
		HeartbeatDetector: NewHeartbeatDetector(),
	}
	detector.SetExecutor(detector)
	return detector
}

// ExecuteFailureDetection runs the failure action
func (detector *GossipDetector) ExecuteFailureDetection(Detector) error {
	finder, err := detector.GetFinder()
	if err != nil {
		return err
	}

	targetNodes, err := finder.GetAllNodes()
	if err != nil {
		return err
	}

	targetAllNodeCount := len(targetNodes)
	targetNodeCount := int(math.Ceil(GossipExecutorDefaultPercentage))
	for n := 0; n < targetNodeCount; n++ {
		targetNode := targetNodes[rand.Intn(targetAllNodeCount)]
		detector.ExecuteNodeFailureDetection(targetNode)
	}

	return nil
}
