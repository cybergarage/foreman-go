// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

import (
	"unsafe"
)

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -lua -lpython
import "C"

// CQuery returns a Query object for Foreman C++.
func (q *Query) CQuery() (unsafe.Pointer, error) {
	cq := C.foreman_metric_query_new()

	C.foreman_metric_query_settarget(cq, C.CString(q.Target))
	C.foreman_metric_query_setinterval(cq, (C.time_t)(q.Interval.Seconds()))
	if q.From != nil {
		C.foreman_metric_query_setfrom(cq, (C.time_t)(q.From.Unix()))
	}
	if q.Until != nil {
		C.foreman_metric_query_setuntil(cq, (C.time_t)(q.Until.Unix()))
	}

	return cq, nil
}
