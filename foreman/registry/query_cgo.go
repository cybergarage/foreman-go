// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -lua -lpython
import "C"

import (
	"unsafe"
)

// CObject returns a registry query for Foreman C++.
func (q *Query) CObject() (unsafe.Pointer, error) {
	cq := C.foreman_registry_query_new()

	C.foreman_registry_query_setparentid(cq, C.CString(q.ParentID))

	return cq, nil
}
