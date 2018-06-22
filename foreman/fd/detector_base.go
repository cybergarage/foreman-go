// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

import (
	"fmt"
)

const (
	errorDetectorNotFoundFinder   = "Finder Not Found"
	errorDetectorNotFoundListener = "Listener Not Found"
)

// baseDetector represents a base detector.
type baseDetector struct {
	finder   Finder
	listener FailureDetectionListener
}

// newBaseDetector returns a new base detector.
func newBaseDetector() *baseDetector {
	detector := &baseDetector{
		finder:   nil,
		listener: nil,
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
func (detector *baseDetector) SetListener(l FailureDetectionListener) error {
	detector.listener = l
	return nil
}

// GetListener returns a current listener
func (detector *baseDetector) GetListener() (FailureDetectionListener, error) {
	if detector.listener == nil {
		return nil, fmt.Errorf(errorDetectorNotFoundListener)
	}
	return detector.listener, nil
}

// ExecuteFailureDetection runs the failure action
func (detector *baseDetector) ExecuteNodeFailureDetection(targetNode Node) error {
	return nil
}
