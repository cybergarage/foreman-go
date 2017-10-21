// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lstdc++
import "C"

import (
	"runtime"
	"unsafe"
)

// MetricResultSet represents a Foreman MetricResultSet.
type MetricResultSet struct {
	cObject unsafe.Pointer
}

// NewMetricResultSet returns a new MetricResultSet.
func NewMetricResultSet() *MetricResultSet {
	return NewMetricResultSetWithCObject(C.foreman_resultset_new())
}

// NewMetricResultSetWithCObject returns a new MetricResultSet from the C++ object.
func NewMetricResultSetWithCObject(cObj unsafe.Pointer) *MetricResultSet {
	rs := &MetricResultSet{}
	rs.cObject = cObj
	runtime.SetFinalizer(rs, resultSetFinalizer)
	return rs
}

func resultSetFinalizer(self *MetricResultSet) {
	if self.cObject != nil {
		if C.foreman_resultset_delete(self.cObject) {
			self.cObject = nil
		}
	}
}

// GetDataPointCount returns a number of the data points.
func (self *MetricResultSet) GetDataPointCount() int {
	if self.cObject == nil {
		return 0
	}
	return int(C.foreman_resultset_getdatapointcount(self.cObject))
}

// GetFirstDataPoints returns a first data points.
func (self *MetricResultSet) GetFirstDataPoints() *DataPoints {
	if self.cObject == nil {
		return nil
	}

	cDpsObject := C.foreman_resultset_firstdatapoints(self.cObject)
	if cDpsObject == nil {
		return nil
	}

	return NewDataPointsWithCObject(cDpsObject)
}

// GetNextDataPoints returns a first data points.
func (self *MetricResultSet) GetNextDataPoints() *DataPoints {
	if self.cObject == nil {
		return nil
	}

	cDpsObject := C.foreman_resultset_nextdatapoints(self.cObject)
	if cDpsObject == nil {
		return nil
	}

	return NewDataPointsWithCObject(cDpsObject)
}
