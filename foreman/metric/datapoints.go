// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -lua -lpython
import "C"

import (
	"unsafe"
)

// DataPoints represents a Foreman DataPoints.
type DataPoints struct {
	Name   string
	Values []*DataPoint
}

// NewDataPoints returns a new DataPoint.
func NewDataPoints(size int) *DataPoints {
	dps := &DataPoints{
		Values: make([]*DataPoint, size),
	}
	return dps
}

// NewDataPointsWithCObject returns a new DataPoint from the C++ object.
func NewDataPointsWithCObject(cObject unsafe.Pointer) *DataPoints {
	dps := NewDataPoints(0)

	dataPointSize := int(C.foreman_metric_datapoints_size(cObject))
	dps.Name = C.GoString(C.foreman_metric_datapoints_getname(cObject))
	dps.Values = make([]*DataPoint, dataPointSize)

	for n := 0; n < dataPointSize; n++ {
		cDpObject := C.foreman_metric_datapoints_get(cObject, C.size_t(n))
		if cDpObject == nil {
			dps.Values[n] = nil
			continue
		}
		dps.Values[n] = NewDataPointWithCObject(cDpObject)
	}

	return dps
}
