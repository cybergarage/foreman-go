// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lstdc++
import "C"

import (
	"unsafe"
)

// CData returns a data object for Foreman C++.
func (self *Metric) CMetric() (unsafe.Pointer, error) {
	cm := C.foreman_metric_new()

	C.foreman_metric_setname(cm, C.CString(self.Name))
	C.foreman_metric_setvalue(cm, C.double(self.Value))
	C.foreman_metric_settimestamp(cm, (C.time_t)(self.Timestamp.Unix()))

	return cm, nil
}
