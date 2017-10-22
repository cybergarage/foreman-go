// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides interfaces for MetricStore of Foreman C++.
package metric

import (
	"fmt"
	"time"
	"unsafe"
)

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lstdc++
import "C"

// Query represents a Foreman Query.
type Query struct {
	Target   string
	From     *time.Time
	Until    *time.Time
	Interval time.Duration
}

// NewQuery returns a new Query.
func NewQuery() *Query {
	q := &Query{
		From:     nil,
		Until:    nil,
		Interval: 0,
	}

	return q
}

// CQuery returns a Query object for Foreman C++.
func (self *Query) CQuery() (unsafe.Pointer, error) {
	cq := C.foreman_metric_query_new()

	C.foreman_metric_query_settarget(cq, C.CString(self.Target))
	C.foreman_metric_query_setinterval(cq, (C.time_t)(self.Interval.Seconds()))
	if self.From != nil {
		C.foreman_metric_query_setfrom(cq, (C.time_t)(self.From.Unix()))
	}
	if self.Until != nil {
		C.foreman_metric_query_setuntil(cq, (C.time_t)(self.Until.Unix()))
	}

	return cq, nil
}

// String returns a string description of the instance
func (self *Query) String() string {
	return fmt.Sprintf("%s [%s - %s]", self.Target, self.From.String(), self.Until.String())
}
