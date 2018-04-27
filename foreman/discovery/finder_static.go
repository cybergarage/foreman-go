// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

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
func (finder *StaticFinder) SearchAll() error {
	return nil
}
