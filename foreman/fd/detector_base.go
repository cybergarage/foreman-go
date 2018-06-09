// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

import (
	"github.com/cybergarage/foreman-go/foreman/discovery"
)

// baseDetector represents a base detector.
type baseDetector struct {
	finder   discovery.Finder
	listener DetectorListener
}

// newBaseDetector returns a new base detector.
func newBaseDetector() *baseDetector {
	detector := &baseDetector{
		listener: nil,
	}
	return detector
}

// SetFinder sets the finder to know target nodes
func (detector *baseDetector) SetFinder(f discovery.Finder) error {
	detector.finder = f
	return nil
}

// SetListener sets a listener to know failure nodes
func (detector *baseDetector) SetListener(l DetectorListener) error {
	detector.listener = l
	return nil
}

// Start starts the instance
func (detector *baseDetector) Start() error {
	return nil
}

// Stop starts the instance
func (detector *baseDetector) Stop() error {
	return nil
}
