// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

// ResultSet represents an abstract interface of metric store for Foreman.
type ResultSet interface {
	GetDataPointCount() int
	GetFirstDataPoints() *DataPoints
	GetNextDataPoints() *DataPoints
}

// goResultSet represents a result set.
type goResultSet struct {
	ResultSet
	dataPoints  []*DataPoints
	iteratorPos int
}

// NewResultSet returns a new result set.
func NewResultSet() *goResultSet {
	return NewResultSetWithSize(0)
}

// NewResultSetWithSize returns a new result set with the specified size.
func NewResultSetWithSize(size int) *goResultSet {
	rs := &goResultSet{
		dataPoints:  make([]*DataPoints, size),
		iteratorPos: 0,
	}
	return rs
}

// AddDataPoints adds a new data point.
func (rs *goResultSet) AddDataPoints(dp *DataPoints) error {
	rs.dataPoints = append(rs.dataPoints, dp)
	return nil
}

// GetDataPointCount returns a number of the data points.
func (rs *goResultSet) GetDataPointCount() int {
	return len(rs.dataPoints)
}

// GetFirstDataPoints returns a first data points.
func (rs *goResultSet) GetFirstDataPoints() *DataPoints {
	rs.iteratorPos = 0
	return rs.GetNextDataPoints()
}

// GetNextDataPoints returns a first data points.
func (rs *goResultSet) GetNextDataPoints() *DataPoints {
	if len(rs.dataPoints) <= rs.iteratorPos {
		return nil
	}
	dp := rs.dataPoints[rs.iteratorPos]
	rs.iteratorPos++
	return dp
}
