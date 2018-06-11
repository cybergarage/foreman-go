// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

import (
	"math"
	"math/rand"
	"time"
)

var GossipExecutorDefaultPercentage = 1.0 / float64(((time.Second * 60) / HeartbeatDetectorDefaultInterval))

// GossipDetector represents a gossip based detector.
type GossipDetector struct {
	*HeartbeatDetector
}

// NewGossipDetector returns a new gossip detector.
func NewGossipDetector() Detector {
	detector := &GossipDetector{
		HeartbeatDetector: NewHeartbeatDetector(),
	}
	return detector
}

// Execute runs the failure action
func (detector *GossipDetector) Execute() error {
	listener, err := detector.GetListener()
	if err != nil {
		return err
	}

	executor, err := detector.GetExecutor()
	if err != nil {
		return err
	}

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
		executor.ExecuteFailureDetection(detector, targetNode)
	}

	return nil
}
