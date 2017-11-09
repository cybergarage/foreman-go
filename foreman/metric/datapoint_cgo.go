// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog
import "C"

import (
	"time"
	"unsafe"
)

// NewDataPointWithCObject returns a new DataPoint from the C++ object.
func NewDataPointWithCObject(cObject unsafe.Pointer) *DataPoint {
	dp := NewDataPoint()
	dp.Value = (float64)(C.foreman_metric_datapoint_getvalue(cObject))
	dp.Timestamp = time.Unix(int64(C.foreman_metric_datapoint_gettimestamp(cObject)), 0)
	return dp
}
