// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

// Metric represents a metric.
type Metrics struct {
	Name   string
	Values []*DataPoint
}

// NewMetricsWithSize returns a new metrics.
func NewMetricsWithSize(size int) *Metrics {
	m := &Metrics{
		Values: make([]*DataPoint, size),
	}
	return m
}

// NewMetrics returns a new metric.
func NewMetrics() *Metrics {
	return NewMetricsWithSize(0)
}
