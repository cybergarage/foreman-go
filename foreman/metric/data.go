// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package metric

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lstdc++
import "C"

import (
	"fmt"
	"time"
	"unsafe"
)

// Data represents a Foreman Data.
type Data struct {
	Name      string
	Value     float64
	Timestamp time.Time
}

// NewData returns a new Data.
func NewData() *Data {
	m := &Data{}

	return m
}

// CMetric returns a Data object for Foreman C++.
func (self *Data) CMetric() (unsafe.Pointer, error) {
	cm := C.foreman_metric_new()

	C.foreman_metric_setname(cm, C.CString(self.Name))
	C.foreman_metric_setvalue(cm, C.double(self.Value))
	C.foreman_metric_settimestamp(cm, (C.time_t)(self.Timestamp.Unix()))

	return cm, nil
}

// String returns a string description of the instance
func (self *Data) String() string {
	return fmt.Sprintf("%s : %f (%d)", self.Name, self.Value, self.Timestamp.Unix())
}
