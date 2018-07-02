// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

import (
	"fmt"
	"time"
)

const (
	DefaultSuspectionDuration = time.Minute * 5
)

const (
	errorDetectorNotFoundFinder   = "Finder Not Found"
	errorDetectorNotFoundListener = "Listener Not Found"
)

// baseDetector represents a base detector.
type baseDetector struct {
	finder             Finder
	listener           FailureDetectorListener
	cluster            *cluster
	suspectionDuration time.Duration
}

// newBaseDetector returns a new base detector.
func newBaseDetector() *baseDetector {
	detector := &baseDetector{
		finder:             nil,
		listener:           nil,
		cluster:            newCluster(),
		suspectionDuration: DefaultSuspectionDuration,
	}
	return detector
}

// SetFinder sets the finder to know target nodes
func (detector *baseDetector) SetFinder(f Finder) error {
	detector.finder = f
	return nil
}

// GetFinder returns a current finder
func (detector *baseDetector) GetFinder() (Finder, error) {
	if detector.finder == nil {
		return nil, fmt.Errorf(errorDetectorNotFoundFinder)
	}
	return detector.finder, nil
}

// SetListener sets a listener
func (detector *baseDetector) SetListener(l FailureDetectorListener) error {
	detector.listener = l
	return nil
}

// GetListener returns a current listener
func (detector *baseDetector) GetListener() (FailureDetectorListener, error) {
	if detector.listener == nil {
		return nil, fmt.Errorf(errorDetectorNotFoundListener)
	}
	return detector.listener, nil
}

// GetCluster returns a cluster status
func (detector *baseDetector) GetCluster() *cluster {
	return detector.cluster
}

// SetSuspectionDuration sets a suspection duration
func (detector *baseDetector) SetSuspectionDuration(duration time.Duration) error {
	detector.suspectionDuration = duration
	return nil
}

// GetSuspectionDuration returns a current suspection duration
func (detector *baseDetector) GetSuspectionDuration() time.Duration {
	return detector.suspectionDuration
}

// ExecuteFailureDetection runs the failure action
func (detector *baseDetector) ExecuteNodeFailureDetection(updatedNode Node) error {
	nodeID := updatedNode.GetUniqueID()

	// Is new node ?

	currentNode, ok := detector.cluster.GetNode(nodeID)
	if !ok {
		if detector.listener != nil {
			detector.listener.FailureDetectorNodeAdded(updatedNode)
		}
		detector.cluster.UpdateStatus(updatedNode)
		return nil
	}

	// Is status changed ?

	if currentNode.GetCondition() != updatedNode.GetCondition() {
		if detector.listener != nil {
			detector.listener.FailureDetectorNodeStatusChanged(updatedNode)
		}
	}

	// Is out of date ?

	if detector.cluster.IsNodeOutOfDate(updatedNode) {
		clusterTimestamp := detector.cluster.GetTimestamp()
		if detector.suspectionDuration <= clusterTimestamp.Sub(time.Now()) {
			if detector.listener != nil {
				detector.listener.FailureDetectorNodeOutOfDate(updatedNode)
			}
		}
	}

	// Update cluster status

	detector.cluster.UpdateStatus(updatedNode)

	return nil
}
