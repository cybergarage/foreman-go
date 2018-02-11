// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

// ResultSet represents an abstract interface of metric store for Foreman.
type ResultSet interface {
	GetMetricsCount() int
	GetFirstMetrics() *Metrics
	GetNextMetrics() *Metrics
}

// goResultSet represents a result set.
type goResultSet struct {
	ResultSet
	Metrics     []*Metrics
	iteratorPos int
}

// NewResultSet returns a new result set.
func NewResultSet() *goResultSet {
	return NewResultSetWithSize(0)
}

// NewResultSetWithSize returns a new result set with the specified size.
func NewResultSetWithSize(size int) *goResultSet {
	rs := &goResultSet{
		Metrics:     make([]*Metrics, size),
		iteratorPos: 0,
	}
	return rs
}

// AddMetrics adds a new data point.
func (rs *goResultSet) AddMetrics(dp *Metrics) error {
	rs.Metrics = append(rs.Metrics, dp)
	return nil
}

// GetMetricsCount returns a number of the data points.
func (rs *goResultSet) GetMetricsCount() int {
	return len(rs.Metrics)
}

// GetFirstMetrics returns a first data points.
func (rs *goResultSet) GetFirstMetrics() *Metrics {
	rs.iteratorPos = 0
	return rs.GetNextMetrics()
}

// GetNextMetrics returns a first data points.
func (rs *goResultSet) GetNextMetrics() *Metrics {
	if len(rs.Metrics) <= rs.iteratorPos {
		return nil
	}
	dp := rs.Metrics[rs.iteratorPos]
	rs.iteratorPos++
	return dp
}
