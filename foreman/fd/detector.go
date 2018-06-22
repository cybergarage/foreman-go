// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

// Detector represents an abstract interface
type Detector interface {
	// SetFinder sets the finder to know target nodes
	SetFinder(Finder) error
	// GetFinder returns a current finder
	GetFinder() (Finder, error)
	// SetListener sets a listener
	SetListener(FailureDetectionListener) error
	// GetListener returns a current listener
	GetListener() (FailureDetectionListener, error)
	// Start starts the instance
	Start() error
	// Stop stop the instance
	Stop() error
}
