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

// CObject returns a registry property for foreman-cc.
func (prop *Property) CObject() (unsafe.Pointer, error) {
	cprop := C.foreman_registry_property_new()
	if cprop == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	C.foreman_registry_property_setname(cprop, C.CString(prop.Name))
	C.foreman_registry_property_setdata(cprop, C.CString(prop.Data))

	return cprop, nil
}
