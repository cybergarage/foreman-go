// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cybergarage/foreman-go/foreman/node"
)

const (
	errorFinderHasSameNode = "Node (%s) is already added"
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

// HasNode returns true when the specified node is added already, otherwise false.
func (finder *baseFinder) HasNode(targetNode Node) bool {
	for _, addedNode := range finder.nodes {
		if node.Equal(targetNode, addedNode) {
			return true
		}
	}
	return false
}

// addNodes adds a specified node.
func (finder *baseFinder) addNode(node Node) error {
	if finder.HasNode(node) {
		return fmt.Errorf(errorFinderHasSameNode, node)
	}
	finder.nodes = append(finder.nodes, node)
	return nil
}

// GetAllNodes returns all found nodes.
func (finder *baseFinder) GetAllNodes() ([]Node, error) {
	return finder.nodes, nil
}

// GetPrefixNodes returns only nodes matching with a specified start string
func (finder *baseFinder) GetPrefixNodes(targetString string) ([]Node, error) {
	nodes, err := finder.GetAllNodes()
	if err != nil {
		return nil, err
	}

	matchedNodes := make([]Node, 0)

	for _, node := range nodes {
		port := node.GetRPCPort()
		addr := node.GetAddress()
		name := node.GetName()

		hosts := []string{
			addr,
			fmt.Sprintf("%s:%d", addr, port),
			name,
			fmt.Sprintf("%s:%d", name, port),
		}

		for _, host := range hosts {
			if len(host) <= 0 {
				continue
			}
			if strings.HasPrefix(targetString, host) {
				matchedNodes = append(matchedNodes, node)
				break
			}
		}
	}

	return matchedNodes, nil
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
		name := node.GetName()

		hosts := []string{
			addr,
			fmt.Sprintf("%s:%d", addr, port),
			name,
			fmt.Sprintf("%s:%d", name, port),
		}

		for _, host := range hosts {
			if len(host) <= 0 {
				continue
			}
			if re.MatchString(host) {
				matchedNodes = append(matchedNodes, node)
				break
			}
		}
	}

	return matchedNodes, nil
}
