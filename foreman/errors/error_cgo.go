// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errors provides errors interfaces
package errors

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -llua -lpython
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

// NewWithCObject creates a error object from the clang object.
func NewWithCObject(cobj unsafe.Pointer) error {
	errMsg := ""

	// Source code infomation

	var cFileName, cFuncName *C.char
	var cLineNo C.int

	C.foreman_error_getfilename(cobj, &cFileName)
	filename := C.GoString(cFileName)

	C.foreman_error_getfuncname(cobj, &cFuncName)
	funcname := C.GoString(cFuncName)

	C.foreman_error_getlineno(cobj, &cLineNo)
	lineno := (int)(cLineNo)

	if 0 < len(filename) || 0 < len(funcname) {
		errMsg += fmt.Sprintf("%s (%d)", funcname, lineno)
	}

	// Basic infomation

	var cMsg *C.char
	var cCode C.int

	C.foreman_error_getmessage(cobj, &cMsg)
	msg := C.GoString(cMsg)

	C.foreman_error_getcode(cobj, &cCode)
	code := (int)(cCode)

	if 0 < len(msg) {
		errMsg += fmt.Sprintf(" : [%d] %s", code, msg)
	}

	// Extra infomation

	var cDetailMsg *C.char
	var cDetailCode C.int

	C.foreman_error_getdetailmessage(cobj, &cDetailMsg)
	detailMsg := C.GoString(cDetailMsg)

	C.foreman_error_getdetailcode(cobj, &cDetailCode)
	detailCode := (int)(cDetailCode)

	if 0 < len(detailMsg) {
		errMsg += fmt.Sprintf(" ([%d] %s)", detailCode, detailMsg)
	}

	return errors.New(errMsg)
}
