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
func NewResultSetWithCObject(cObj unsafe.Pointer) *ResultSet {
	rs := &ResultSet{}
	rs.cResultSet = cObj
	runtime.SetFinalizer(rs, resultSetFinalizer)
	return rs
}

func resultSetFinalizer(self *ResultSet) {
	if self.cResultSet != nil {
		if C.foreman_resultset_delete(self.cResultSet) {
			self.cResultSet = nil
		}
	}
}

// GetCount returns a number of the result set.
func (self *ResultSet) GetCount() (int64, error) {
	if self.cResultSet == nil {
		return 0, fmt.Errorf(errorClangObjectNotInitialized)
	}
	return int64(C.foreman_resultset_getcount(self.cResultSet)), nil
}

// GetValue returns a metric value of the specified index.
func (self *ResultSet) GetValue(n int64) (float64, error) {
	if self.cResultSet == nil {
		return 0, fmt.Errorf(errorClangObjectNotInitialized)
	}
	var value C.double
	isSuccess := C.foreman_resultset_getvalue(self.cResultSet, C.size_t(n), &value)
	if !isSuccess {
		return 0, fmt.Errorf(errorResultSetCouldGetValue, n)
	}
	return float64(value), nil
}
