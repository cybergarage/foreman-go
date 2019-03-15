// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/node"
)

// StaticFinder represents a simple static finder.
type StaticFinder struct {
	*baseFinder
}

// NodeList is a format of nodelist.yaml.
type NodeList struct {
	clusterName string   `yaml:"cluster_name"`
	hosts       []string `yaml:"hosts"`
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

// NewStaticFinder returns a new static finder.
func NewStaticFinder() Finder {
	return NewStaticFinderWithNodes(nil)
}

// NewStaticFinderWithHosts returns a new static finder.
func NewStaticFinderWithHosts(hosts []string, cluster string) Finder {
	nodes := []Node{}
	for _, host := range hosts {
		node := node.NewBaseNode()
		node.Cluster = cluster
		node.Name = host
		nodes = append(nodes, node)
	}
	return NewStaticFinderWithNodes(nodes)
}

// Search searches all nodes.
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
