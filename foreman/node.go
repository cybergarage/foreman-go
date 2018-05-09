// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/discovery"
)

const (
	insertMetricQueryFormat = "SET (%s, %f, %d) INTO METRICS"
)

// Node represents an abstract node interface
type Node interface {
	// GetCluster returns the cluster name
	GetCluster() string
	// GetName returns the host name
	GetName() string
	// GetAddress returns the interface address
	GetAddress() string
	// GetRPCPort returns the RPC port
	GetRPCPort() int
	// GetRenderPort returns the Graphite render port
	GetRenderPort() int
	// GetCarbonPort returns the Graphite carbon port
	GetCarbonPort() int
	// PostQuery posts a query string
	PostQuery(query string) (interface{}, error)
}

// NodeEqual returns true if the other node is same with this node
func NodeEqual(this, other Node) bool {
	return discovery.NodeEqual(this, other)
}
