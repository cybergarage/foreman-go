// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

// #include <foreman/foreman-c.h>
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

// CObject returns a registry query for foreman-cc.
func (q *Query) CObject() (unsafe.Pointer, error) {
	cq := C.foreman_registry_query_new()
	if cq == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cparentID := C.CString(q.ParentID)
	defer C.free(unsafe.Pointer(cparentID))
	C.foreman_registry_query_setparentid(cq, cparentID)

	return cq, nil
}
