// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// #include <foreman/foreman-c.h>
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

// NewMethodWithCObject returns a new method of the specified object.
func NewMethodWithCObject(cObject unsafe.Pointer) *Method {
	method := NewMethod()

	var cname *C.char
	if C.foreman_action_method_getname(cObject, &cname) {
		method.Name = C.GoString(cname)
	}

	var clang *C.char
	if C.foreman_action_method_getlanguage(cObject, &clang) {
		method.Language = C.GoString(clang)
	}

	var ccode *C.char
	if C.foreman_action_method_getstringcode(cObject, &ccode) {
		method.Code = []byte(C.GoString(ccode))
	}

	return method
}

// CObject returns a method objectfor foreman-cc.
func (method *Method) CObject() (unsafe.Pointer, error) {
	clang := C.CString(method.Language)
	defer C.free(unsafe.Pointer(clang))

	cmethod := C.foreman_action_method_new(clang)
	if cmethod == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cname := C.CString(method.Name)
	defer C.free(unsafe.Pointer(cname))
	C.foreman_action_method_setname(cmethod, cname)

	ccode := C.CString(string(method.Code))
	defer C.free(unsafe.Pointer(ccode))
	C.foreman_action_method_setstringcode(cmethod, ccode)

	return cmethod, nil
}
