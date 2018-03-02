// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

// #include <foreman/foreman-c.h>
import "C"

import (
	"unsafe"
)

// NewMetricsWithCObject returns a new DataPoint from the C++ object.
func NewMetricsWithCObject(cObject unsafe.Pointer) *Metrics {
	m := NewMetricsWithSize(0)

	dataPointSize := int(C.foreman_metric_metrics_getdatapointsize(cObject))
	m.Name = C.GoString(C.foreman_metric_metrics_getname(cObject))
	m.Values = make([]*DataPoint, dataPointSize)

	for n := 0; n < dataPointSize; n++ {
		cDpObject := C.foreman_metric_metrics_getdatapoint(cObject, C.size_t(n))
		if cDpObject == nil {
			m.Values[n] = nil
			continue
		}
		m.Values[n] = NewDataPointWithCObject(cDpObject)
	}

	return m
}
