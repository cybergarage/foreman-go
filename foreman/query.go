// Copyright 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"time"
	"unsafe"
)

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lstdc++
import "C"

// Query represents a Foreman Query.
type Query struct {
	Target   string
	From     time.Time
	Until    time.Time
	Interval int
}

// NewQuery returns a new Query.
func NewQuery() *Query {
	q := &Query{}

	q.From = time.Unix(0, 0)
	q.Until = time.Unix(0, 0)
	q.Interval = 0

	return q
}

// CQuery returns a Query for Foreman C++.
func (self *Query) CQuery() (unsafe.Pointer, error) {
	cq := C.foreman_query_new()

	C.foreman_query_settarget(cq, C.CString(self.Target))
	C.foreman_query_setfrom(cq, (C.time_t)(self.From.Unix()))
	C.foreman_query_setuntil(cq, (C.time_t)(self.Until.Unix()))
	C.foreman_query_setinterval(cq, (C.time_t)(self.Interval))

	return cq, nil
}
