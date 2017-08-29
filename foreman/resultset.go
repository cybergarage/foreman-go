// Copyright 2017 Satoshi Konno. All rights reserved.
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
