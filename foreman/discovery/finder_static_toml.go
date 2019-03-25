// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/node"
)

// StaticFinder represents a simple static finder.
type StaticTOMLFinder struct {
	*StaticFinder
}

// NewStaticTOMLFinderWithNodes returns a new static finder with specified nodes.
func NewStaticTOMLFinderWithNodes(nodes []Node) Finder {
	finder := &StaticFinder{
		baseFinder: newBaseFinder(),
	}

	for _, node := range nodes {
		finder.addNode(node)
	}

	return finder
}

// NewStaticTOMLFinderWithHosts returns a new static finder.
func NewStaticTOMLFinderWithHosts(hosts []string, cluster string) Finder {
	nodes := []Node{}
	for _, host := range hosts {
		node := node.NewBaseNode()
		node.Cluster = cluster
		node.Name = host
		nodes = append(nodes, node)
	}
	return NewStaticTOMLFinderWithNodes(nodes)
}

// NewStaticTOMLFinder returns a new static finder.
func NewStaticTOMLFinder() Finder {
	return NewStaticTOMLFinderWithNodes(nil)
}

// String returns the description
func (finder *StaticFinder) String() string {
	return fmt.Sprintf("%s", FinderStaticToml)
}
