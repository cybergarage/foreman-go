// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/discovery"
)

const (
	errorDetectorNotFoundFinder   = "Finder Not Found"
	errorDetectorNotFoundListener = "Listener Not Found"
	errorDetectorNotFoundExecutor = "Executor Not Found"
)

// baseDetector represents a base detector.
type baseDetector struct {
	finder   Finder
	listener FailureDetectionListener
	executor FailureDetectionExecutor
}

// newBaseDetector returns a new base detector.
func newBaseDetector() *baseDetector {
	detector := &baseDetector{
		finder:   nil,
		listener: nil,
		executor: nil,
	}
	return detector
}

// SetFinder sets the finder to know target nodes
func (detector *baseDetector) SetFinder(f discovery.Finder) error {
	detector.finder = f
	return nil
}

// GetFinder returns a current finder
func (detector *baseDetector) GetFinder() (discovery.Finder, error) {
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

// SetExecutor sets a executor
func (detector *baseDetector) SetExecutor(e FailureDetectionExecutor) error {
	detector.executor = e
	return nil
}

// GetExecutor returns a current executor
func (detector *baseDetector) GetExecutor() (FailureDetectionExecutor, error) {
	if detector.listener == nil {
		return nil, fmt.Errorf(errorDetectorNotFoundExecutor)
	}
	return detector.executor, nil
}
