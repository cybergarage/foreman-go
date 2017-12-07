// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

import (
	"fmt"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -llua -lpython
import "C"

// CQuery returns a Query object for foreman-cc.
func (q *Query) CQuery() (unsafe.Pointer, error) {
	cq := C.foreman_metric_query_new()
	if cq == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

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
