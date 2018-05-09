// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

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
	// PostQuery posts a query string
	PostQuery(query string) (interface{}, error)
}
