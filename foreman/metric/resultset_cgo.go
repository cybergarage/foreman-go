// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lstdc++
import "C"

import (
	"runtime"
	"unsafe"
)

// CgoResultSet represents a Foreman ResultSet.
type CgoResultSet struct {
	cObject unsafe.Pointer
}

// NewResultSet returns a new ResultSet.
func NewResultSet() *CgoResultSet {
	return NewResultSetWithCObject(C.foreman_metric_resultset_new())
}

// NewResultSetWithCObject returns a new ResultSet from the C++ object.
func NewResultSetWithCObject(cObj unsafe.Pointer) *CgoResultSet {
	rs := &CgoResultSet{}
	rs.cObject = cObj
	runtime.SetFinalizer(rs, resultSetFinalizer)
	return rs
}

func resultSetFinalizer(self *CgoResultSet) {
	if self.cObject != nil {
		if C.foreman_metric_resultset_delete(self.cObject) {
			self.cObject = nil
		}
	}
}

// GetDataPointCount returns a number of the data points.
func (self *CgoResultSet) GetDataPointCount() int {
	if self.cObject == nil {
		return 0
	}
	return int(C.foreman_metric_resultset_getdatapointcount(self.cObject))
}

// GetFirstDataPoints returns a first data points.
func (self *CgoResultSet) GetFirstDataPoints() *DataPoints {
	if self.cObject == nil {
		return nil
	}

	cDpsObject := C.foreman_metric_resultset_firstdatapoints(self.cObject)
	if cDpsObject == nil {
		return nil
	}

	return NewDataPointsWithCObject(cDpsObject)
}

// GetNextDataPoints returns a first data points.
func (self *CgoResultSet) GetNextDataPoints() *DataPoints {
	if self.cObject == nil {
		return nil
	}

	cDpsObject := C.foreman_metric_resultset_nextdatapoints(self.cObject)
	if cDpsObject == nil {
		return nil
	}

	return NewDataPointsWithCObject(cDpsObject)
}
