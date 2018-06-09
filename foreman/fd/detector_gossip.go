// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

// GossipDetector represents a gossip based detector.
type GossipDetector struct {
	*baseDetector
}

// NewGossipDetector returns a new gossip detector.
func NewGossipDetector() Detector {
	detector := &GossipDetector{
		baseDetector: newBaseDetector(),
	}
	return detector
}

// Execute runs the failure action
func (detector *baseDetector) Execute() error {
	return nil
}
