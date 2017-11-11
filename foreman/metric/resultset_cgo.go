// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog
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

func resultSetFinalizer(r *CgoResultSet) {
	if r.cObject != nil {
		if C.foreman_metric_resultset_delete(r.cObject) {
			r.cObject = nil
		}
	}
}

// GetDataPointCount returns a number of the data points.
func (r *CgoResultSet) GetDataPointCount() int {
	if r.cObject == nil {
		return 0
	}
	return int(C.foreman_metric_resultset_getdatapointcount(r.cObject))
}

// GetFirstDataPoints returns a first data points.
func (r *CgoResultSet) GetFirstDataPoints() *DataPoints {
	if r.cObject == nil {
		return nil
	}

	cDpsObject := C.foreman_metric_resultset_firstdatapoints(r.cObject)
	if cDpsObject == nil {
		return nil
	}

	return NewDataPointsWithCObject(cDpsObject)
}

// GetNextDataPoints returns a first data points.
func (r *CgoResultSet) GetNextDataPoints() *DataPoints {
	if r.cObject == nil {
		return nil
	}

	cDpsObject := C.foreman_metric_resultset_nextdatapoints(r.cObject)
	if cDpsObject == nil {
		return nil
	}

	return NewDataPointsWithCObject(cDpsObject)
}
