// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

import (
	"time"

	"github.com/cybergarage/foreman-go/foreman/node"
)

// cluster represents a current latest status.
type cluster struct {
	currentNodes  map[string]*node.BaseNode
	latestVersion node.Version
	latestClock   node.Clock
	timestamp     time.Time
}

// newCluster returns a new cluster
func newCluster() *cluster {
	cluster := &cluster{
		currentNodes:  make(map[string]*node.BaseNode),
		latestVersion: 0,
		latestClock:   0,
		timestamp:     time.Now(),
	}
	return cluster
}

// GetNode returns the added node by the specified ID
func (cluster *cluster) GetNode(nodeID string) (*node.BaseNode, bool) {
	node, ok := cluster.currentNodes[nodeID]
	return node, ok
}

// IsNodeOutOfDate checks whether the specified node is out of date
func (cluster *cluster) IsNodeOutOfDate(updatedNode Node) bool {
	if updatedNode.GetVersion() < cluster.latestVersion {
		return true
	}
	return false
}

// UpdateStatus updates the cluster status by the specified node
func (cluster *cluster) UpdateStatus(updatedNode Node) error {
	nodeID := updatedNode.GetUniqueID()
	currentNode, ok := cluster.currentNodes[nodeID]
	if !ok {
		currentNode = node.NewBaseNode()
		cluster.currentNodes[nodeID] = currentNode
	}

	updateNodeVersion := currentNode.GetVersion()
	if cluster.latestVersion < updateNodeVersion {
		cluster.latestVersion = updateNodeVersion
		cluster.timestamp = time.Now()
	}

	updateNodeClock := currentNode.GetClock()
	if cluster.latestClock < updateNodeClock {
		cluster.latestClock = updateNodeClock
		cluster.timestamp = time.Now()
	}

	currentNode.SetStatus(updatedNode)

	return nil
}

// GetVersion returns the latest version in the cluster nodes
func (cluster *cluster) GetVersion() node.Version {
	return cluster.latestVersion
}

// GetClock returns the latest logical clock in the cluster nodes
func (cluster *cluster) GetClock() node.Clock {
	return cluster.latestClock
}

// GetTimestamp returns the latest timestamp
func (cluster *cluster) GetTimestamp() time.Time {
	return cluster.timestamp
}
