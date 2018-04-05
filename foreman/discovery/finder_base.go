// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import "regexp"

// baseFinder represents a base finder.
type baseFinder struct {
	nodes          []*Node
	searchListener FinderSearchListener
	notifyListener FinderNotifyListener
}

// newBaseFinder returns a new base finder.
func newBaseFinder() *baseFinder {
	finder := &baseFinder{
		nodes:          make([]*Node, 0),
		searchListener: nil,
		notifyListener: nil,
	}
	return finder
}

// SetSearchListener sets the search listener.
func (finder *baseFinder) SetSearchListener(l FinderSearchListener) error {
	finder.searchListener = l
	return nil
}

// SetSearchListener sets the search listener.
func (finder *baseFinder) SetNotifyListener(l FinderNotifyListener) error {
	finder.notifyListener = l
	return nil
}

// addNodes adds a specified node.
func (finder *baseFinder) addNode(node *Node) error {
	finder.nodes = append(finder.nodes, node)
	return nil
}

// GetNodes returns only specified regex nodes.
func (finder *baseFinder) GetNodes(regexp.Regexp) ([]*Node, error) {
	return nil, nil
}

// GetAllNodes returns all found nodes.
func (finder *baseFinder) GetAllNodes() ([]*Node, error) {
	return finder.nodes, nil
}
