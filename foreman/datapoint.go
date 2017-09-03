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
	"time"
	"unsafe"
)

// DataPoint represents a Foreman DataPoint.
type DataPoint struct {
	Value     float64
	Timestamp time.Time
}

// NewDataPoint returns a new DataPoint.
func NewDataPoint() *DataPoint {
	dp := &DataPoint{}
	return dp
}

// NewDataPointWithCObject returns a new DataPoint from the C++ object.
func NewDataPointWithCObject(cObject unsafe.Pointer) *DataPoint {
	dp := NewDataPoint()
	dp.Value = (float64)(C.foreman_datapoint_getvalue(cObject))
	dp.Timestamp = time.Unix(int64(C.foreman_datapoint_gettimestamp(cObject)), 0)
	return dp
}

// String returns a string description of the instance
func (self *DataPoint) String() string {
	return fmt.Sprintf("[%d] %f", self.Timestamp.Unix(), self.Value)
}
