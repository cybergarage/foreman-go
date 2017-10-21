// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"fmt"
	"time"
	"unsafe"
)

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lstdc++
import "C"

// MetricQuery represents a Foreman MetricQuery.
type MetricQuery struct {
	Target   string
	From     *time.Time
	Until    *time.Time
	Interval time.Duration
}

// NewQuery returns a new MetricQuery.
func NewMetricQuery() *MetricQuery {
	q := &MetricQuery{
		From:     nil,
		Until:    nil,
		Interval: 0,
	}

	return q
}

// CQuery returns a MetricQuery object for Foreman C++.
func (self *MetricQuery) CQuery() (unsafe.Pointer, error) {
	cq := C.foreman_query_new()

	C.foreman_query_settarget(cq, C.CString(self.Target))
	C.foreman_query_setinterval(cq, (C.time_t)(self.Interval.Seconds()))
	if self.From != nil {
		C.foreman_query_setfrom(cq, (C.time_t)(self.From.Unix()))
	}
	if self.Until != nil {
		C.foreman_query_setuntil(cq, (C.time_t)(self.Until.Unix()))
	}

	return cq, nil
}

// String returns a string description of the instance
func (self *MetricQuery) String() string {
	return fmt.Sprintf("%s [%s - %s]", self.Target, self.From.String(), self.Until.String())
}
