// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -llua -lpython
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

// CMetric returns a data object for foreman-cc.
func (m *Metric) CMetric() (unsafe.Pointer, error) {
	cm := C.foreman_metric_new()
	if cm == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	C.foreman_metric_setname(cm, C.CString(m.Name))
	C.foreman_metric_setvalue(cm, C.double(m.Value))
	C.foreman_metric_settimestamp(cm, (C.time_t)(m.Timestamp.Unix()))

	return cm, nil
}
