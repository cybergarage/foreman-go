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

// Metric represents a Foreman Metric.
type Metric struct {
	Name      string
	Value     float64
	Timestamp time.Time
}

// NewMetric returns a new Metric.
func NewMetric() *Metric {
	m := &Metric{}

	return m
}

// CMetric returns a Metric object for Foreman C++.
func (self *Metric) CMetric() (unsafe.Pointer, error) {
	cm := C.foreman_metric_new()

	C.foreman_metric_setname(cm, C.CString(self.Name))
	C.foreman_metric_setvalue(cm, C.double(self.Value))
	C.foreman_metric_settimestamp(cm, (C.time_t)(self.Timestamp.Unix()))

	return cm, nil
}

// String returns a string description of the instance
func (self *Metric) String() string {
	return fmt.Sprintf("%s : %f (%d)", self.Name, self.Value, self.Timestamp.Unix())
}
