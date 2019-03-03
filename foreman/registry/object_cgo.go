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
		if C.foreman_registry_object_getpropertydata(cObject, &cString) {
			obj.propertyData = C.GoString(cString)
		}
	}

	return obj
}

// CObject returns a registry object for foreman-cc.
func (obj *Object) CObject() (unsafe.Pointer, error) {
	cobj := C.foreman_registry_object_new()
	if cobj == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cID := C.CString(obj.ID)
	defer C.free(unsafe.Pointer(cID))
	C.foreman_registry_object_setid(cobj, cID)

	cparentID := C.CString(obj.ParentID)
	defer C.free(unsafe.Pointer(cparentID))
	C.foreman_registry_object_setparentid(cobj, cparentID)

	cname := C.CString(obj.Name)
	defer C.free(unsafe.Pointer(cname))
	C.foreman_registry_object_setname(cobj, cname)

	cdata := C.CString(obj.Data)
	defer C.free(unsafe.Pointer(cdata))
	C.foreman_registry_object_setdata(cobj, cdata)

	cpropData := C.CString(obj.propertyData)
	defer C.free(unsafe.Pointer(cpropData))
	C.foreman_registry_object_setpropertydata(cobj, cpropData)

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
