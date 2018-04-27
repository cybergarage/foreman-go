// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"fmt"
	"regexp"
)

// baseFinder represents a base finder.
type baseFinder struct {
	nodes          []Node
	searchListener FinderSearchListener
	notifyListener FinderNotifyListener
}

// newBaseFinder returns a new base finder.
func newBaseFinder() *baseFinder {
	finder := &baseFinder{
		nodes:          make([]Node, 0),
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
func (finder *baseFinder) addNode(node Node) error {
	finder.nodes = append(finder.nodes, node)
	return nil
}

// GetAllNodes returns all found nodes.
func (finder *baseFinder) GetAllNodes() ([]Node, error) {
	return finder.nodes, nil
}

// GetRegexpNodes returns only nodes matching with a specified regular expression
func (finder *baseFinder) GetRegexpNodes(re *regexp.Regexp) ([]Node, error) {
	nodes, err := finder.GetAllNodes()
	if err != nil {
		return nil, err
	}

	matchedNodes := make([]Node, 0)

	for _, node := range nodes {
		port := node.GetRPCPort()

		addr := node.GetAddress()
		if re.MatchString(addr) {
			matchedNodes = append(matchedNodes, node)
			break
		}
		if re.MatchString(fmt.Sprintf("%s:%d", addr, port)) {
			matchedNodes = append(matchedNodes, node)
			break
		}

		name := node.GetName()
		if re.MatchString(name) {
			matchedNodes = append(matchedNodes, node)
			break
		}
		if re.MatchString(fmt.Sprintf("%s:%d", name, port)) {
			matchedNodes = append(matchedNodes, node)
			break
		}
	}

	return matchedNodes, nil
}
