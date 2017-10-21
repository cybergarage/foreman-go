// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides interfaces for MetricStore of Foreman C++.
package metric

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lstdc++
import "C"

import (
	"runtime"
	"unsafe"
)

// ResultSet represents a Foreman ResultSet.
type ResultSet struct {
	cObject unsafe.Pointer
}

// NewResultSet returns a new ResultSet.
func NewResultSet() *ResultSet {
	return NewResultSetWithCObject(C.foreman_resultset_new())
}

// NewResultSetWithCObject returns a new ResultSet from the C++ object.
func NewResultSetWithCObject(cObj unsafe.Pointer) *ResultSet {
	rs := &ResultSet{}
	rs.cObject = cObj
	runtime.SetFinalizer(rs, resultSetFinalizer)
	return rs
}

func resultSetFinalizer(self *ResultSet) {
	if self.cObject != nil {
		if C.foreman_resultset_delete(self.cObject) {
			self.cObject = nil
		}
	}
}

// GetDataPointCount returns a number of the data points.
func (self *ResultSet) GetDataPointCount() int {
	if self.cObject == nil {
		return 0
	}
	return int(C.foreman_resultset_getdatapointcount(self.cObject))
}

// GetFirstDataPoints returns a first data points.
func (self *ResultSet) GetFirstDataPoints() *DataPoints {
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
func (self *ResultSet) GetNextDataPoints() *DataPoints {
	if self.cObject == nil {
		return nil
	}

	cDpsObject := C.foreman_resultset_nextdatapoints(self.cObject)
	if cDpsObject == nil {
		return nil
	}

	return NewDataPointsWithCObject(cDpsObject)
}
