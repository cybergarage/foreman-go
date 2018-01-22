// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -llua -lpython
import "C"

import (
	"runtime"
	"unsafe"
)

// cgoResultSet represents a result set of foreman-cc.
type cgoResultSet struct {
	ResultSet
	cObject unsafe.Pointer
}

// NewResultSetWithCObject returns a new ResultSet from the C++ object.
func NewResultSetWithCObject(cObj unsafe.Pointer) *cgoResultSet {
	rs := &cgoResultSet{}
	rs.cObject = cObj
	runtime.SetFinalizer(rs, resultSetFinalizer)
	return rs
}

func resultSetFinalizer(rs *cgoResultSet) {
	if rs.cObject != nil {
		if C.foreman_metric_resultset_delete(rs.cObject) {
			rs.cObject = nil
		}
	}
}

// GetDataPointCount returns a number of the data points.
func (rs *cgoResultSet) GetDataPointCount() int {
	if rs.cObject == nil {
		return 0
	}
	return int(C.foreman_metric_resultset_getdatapointcount(rs.cObject))
}

// GetFirstDataPoints returns a first data points.
func (rs *cgoResultSet) GetFirstDataPoints() *DataPoints {
	if rs.cObject == nil {
		return nil
	}

	cDpsObject := C.foreman_metric_resultset_firstdatapoints(rs.cObject)
	if cDpsObject == nil {
		return nil
	}

	return NewDataPointsWithCObject(cDpsObject)
}

// GetNextDataPoints returns a first data points.
func (rs *cgoResultSet) GetNextDataPoints() *DataPoints {
	if rs.cObject == nil {
		return nil
	}

	cDpsObject := C.foreman_metric_resultset_nextdatapoints(rs.cObject)
	if cDpsObject == nil {
		return nil
	}

	return NewDataPointsWithCObject(cDpsObject)
}
