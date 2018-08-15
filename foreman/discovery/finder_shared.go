// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

// SharedFinder represents a simple finder.
type SharedFinder struct {
	*baseFinder
}

var sharedFinder = &SharedFinder{
	baseFinder: newBaseFinder(),
}

// NewSharedFinder returns a new shared finder.
func NewSharedFinder() Finder {
	return sharedFinder
}

// SearchAll searches all nodes.
func (finder *SharedFinder) SearchAll() error {
	return nil
}

// Start starts the finder.
func (finder *SharedFinder) Start() error {
	return nil
}

// Stop stops the finder.
func (finder *SharedFinder) Stop() error {
	return nil
}
