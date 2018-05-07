// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/metric"
)

const (
	insertMetricQueryFormat = "SET (%s, %f, %d) INTO METRICS"
)

// Node represents an abstract node interface
type Node interface {
	// GetCuster returns the cluster name
	GetCuster() string
	// GetName returns the host name
	GetName() string
	// GetAddress returns the interface address
	GetAddress() string
	// GetRPCPort returns the RPC port
	GetRPCPort() int
	// PostQuery posts a query string
	PostQuery(query string) (interface{}, error)

	// PostMetric posts a metric
	PostMetric(m *metric.Metric) error
}

// nodePostMetric posts a metric
func nodePostMetric(node Node, m *metric.Metric) error {
	query := fmt.Sprintf(insertMetricQueryFormat, m.GetName(), m.GetValue(), m.GetTimestamp().Unix())
	_, err := node.PostQuery(query)
	return err
}
