// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

import (
	"fmt"
	"time"

	"github.com/cybergarage/foreman-go/foreman/node"
)

const (
	DefaultSuspectionDuration = time.Minute * 2
)

const (
	errorDetectorNotFoundFinder   = "Finder Not Found"
	errorDetectorNotFoundListener = "Listener Not Found"
)

// baseDetector represents a base detector.
type baseDetector struct {
	finder             Finder
	listener           FailureDetectorListener
	currentNodes       map[string]*node.BaseNode
	suspectionDuration time.Duration
}

// newBaseDetector returns a new base detector.
func newBaseDetector() *baseDetector {
	detector := &baseDetector{
		finder:             nil,
		listener:           nil,
		currentNodes:       make(map[string]*node.BaseNode),
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

// SetSuspectionDuration sets a suspection duration for failure detection
func (detector *baseDetector) SetSuspectionDuration(d time.Duration) error {
	detector.suspectionDuration = d
	return nil
}

// GetSuspectionDuration gets a suspection duration for failure detection
func (detector *baseDetector) GetSuspectionDuration() time.Duration {
	return detector.suspectionDuration
}

// ExecuteFailureDetection runs the failure action
func (detector *baseDetector) ExecuteNodeFailureDetection(updatedNode Node) error {
	nodeID := updatedNode.GetUniqueID()

	// New node ?

	currentNode, ok := detector.currentNodes[nodeID]
	if !ok {
		currentNode = node.NewBaseNode()
		detector.currentNodes[nodeID] = currentNode
		if detector.listener != nil {
			detector.listener.FailureDetectorNodeAdded(updatedNode)
		}
	}

	// Is status change ?

	if currentNode.GetCondition() != updatedNode.GetCondition() {
		if detector.listener != nil {
			detector.listener.FailureDetectorNodeStatusChanged(updatedNode)
		}
	}

	// Update node status

	currentNode.SetStatus(updatedNode)

	return nil
}
