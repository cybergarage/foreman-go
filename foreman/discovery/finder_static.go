// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import "fmt"

// StaticFinder represents a simple static finder.
type StaticFinder struct {
	*baseFinder
}

// NewStaticFinderWithNodes returns a new static finder with specified nodes.
func NewStaticFinderWithNodes(nodes []Node) Finder {
	finder := &StaticFinder{
		baseFinder: newBaseFinder(),
	}

	for _, node := range nodes {
		finder.addNode(node)
	}

	return finder
}

// SearchAll searches all nodes.
func (finder *StaticFinder) Search() error {
	return nil
}

// Start starts the finder.
func (finder *StaticFinder) Start() error {
	return nil
}

// Stop stops the finder.
func (finder *StaticFinder) Stop() error {
	return nil
}

// IsRunning returns true when the finder is running, otherwise false.
func (finder *StaticFinder) IsRunning() bool {
	return true
}

// String returns the description
func (finder *StaticFinder) String() string {
	return fmt.Sprintf("%s", FinderStatic)
}
