// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog
import "C"

import (
	"unsafe"
)

// NewObjectWithCObject returns a new object from the C++ object.
func newObjectWithCObject(cObject unsafe.Pointer) *Object {
	obj := NewObject()

	if cObject != nil {
		var cString *C.char
		if C.foreman_registry_object_getid(cObject, &cString) {
			obj.ID = C.GoString(cString)
		}
		if C.foreman_registry_object_getparentid(cObject, &cString) {
			obj.ParentID = C.GoString(cString)
		}
		if C.foreman_registry_object_getname(cObject, &cString) {
			obj.Name = C.GoString(cString)
		}
		if C.foreman_registry_object_getdata(cObject, &cString) {
			obj.Data = C.GoString(cString)
		}
	}

	return obj
}

// CObject returns a registry object for Foreman C++.
func (obj *Object) CObject() (unsafe.Pointer, error) {
	cobj := C.foreman_registry_object_new()

	C.foreman_registry_object_setid(cobj, C.CString(obj.ID))
	C.foreman_registry_object_setparentid(cobj, C.CString(obj.ParentID))
	C.foreman_registry_object_setname(cobj, C.CString(obj.Name))
	C.foreman_registry_object_setdata(cobj, C.CString(obj.Data))

	return cobj, nil
}

// newObjectsWithCObjects returns a new object from the C++ object.
func newObjectsWithCObjects(cObjects unsafe.Pointer) []*Object {
	objCount := (int)(C.foreman_registry_objects_size(cObjects))
	goObjects := make([]*Object, objCount)
	for n := 0; n < objCount; n++ {
		cObject := C.foreman_registry_objects_getobject(cObjects, (C.size_t)(n))
		if cObject == nil {
			goObjects[n] = NewObject()
			continue
		}
		goObjects[n] = newObjectWithCObject(cObject)
	}

	return goObjects
}