// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog
import "C"
import "unsafe"

// CObject returns a registry property for Foreman C++.
func (self *Property) CObject() (unsafe.Pointer, error) {
	prop := C.foreman_registry_property_new()

	C.foreman_registry_property_setname(prop, C.CString(self.Name))
	C.foreman_registry_property_setdata(prop, C.CString(self.Data))

	return prop, nil
}
