// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

import (
	"github.com/cybergarage/foreman-go/foreman/discovery"
)

// DetectorListener a listener for detector.
type DetectorListener interface {
	NodeFailureDetected(*discovery.Node)
}

// Detector represents an abstract interface
type Detector interface {
	// SetFinder sets the finder to know target nodes
	SetFinder(discovery.Finder) error
	// SetListener sets a listener to know failure nodes
	SetListener(DetectorListener) error
	// Start starts the instance
	Start() error
	// Stop stop the instance
	Stop() error
	// Execute runs the failure action
	Execute() error
}
