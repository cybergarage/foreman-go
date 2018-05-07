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

// baseNode represents a base node.
type baseNode struct {
	Node
}

// newBaseNode returns a new base node.
func newBaseNode() *baseNode {
	node := &baseNode{}
	return node
}

// PostMetric posts a metric
func (node *baseNode) PostMetric(m *metric.Metric) error {
	query := fmt.Sprintf(insertMetricQueryFormat, m.GetName(), m.GetValue(), m.GetTimestamp().Unix())
	_, err := node.PostQuery(query)
	return err
}
