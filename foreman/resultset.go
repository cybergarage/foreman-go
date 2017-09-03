// Copyright 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lstdc++
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

// ResultSet represents a Foreman ResultSet.
type ResultSet struct {
	cResultSet unsafe.Pointer
}

// NewResultSet returns a new ResultSet.
func NewResultSet() *ResultSet {
	rs := &ResultSet{}
	runtime.SetFinalizer(rs, resultSetFinalizer)
	return rs
}

// NewResultSetWithCObject returns a new ResultSet from the C++ object.
func NewResultSetWithCObject(cObj unsafe.Pointer) *ResultSet {
	rs := NewResultSet()
	return rs
}

func resultSetFinalizer(self *ResultSet) {
	if self.cResultSet != nil {
		if C.foreman_resultset_delete(self.cResultSet) {
			self.cResultSet = nil
		}
	}
}

// GetDataPointCount returns a number of the data points.
func (self *ResultSet) GetDataPointCount() (int, error) {
	if self.cResultSet == nil {
		return 0, fmt.Errorf(errorClangObjectNotInitialized)
	}
	return int(C.foreman_resultset_getdatapointcount(self.cResultSet)), nil
}

// GetFirstDataPoints returns a first data points.
func (self *ResultSet) GetFirstDataPoints() *DataPoints {
	if self.cResultSet == nil {
		return nil
	}

	cDpsObject := C.foreman_resultset_firstdatapoints(self.cResultSet)
	if cDpsObject == nil {
		return nil
	}

	return NewDataPointsWithCObject(cDpsObject)
}

// GetNextDataPoints returns a first data points.
func (self *ResultSet) GetNextDataPoints() *DataPoints {
	if self.cResultSet == nil {
		return nil
	}

	cDpsObject := C.foreman_resultset_nextdatapoints(self.cResultSet)
	if cDpsObject == nil {
		return nil
	}

	return NewDataPointsWithCObject(cDpsObject)
}
