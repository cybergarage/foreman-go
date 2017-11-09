// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errors provides errors interfaces
package errors

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

// NewWithCObject creates a error object from the clang object.
func NewWithCObject(cobj unsafe.Pointer) error {
	var cMsg, cDetailMsg *C.char

	C.foreman_error_getmessage(cobj, &cMsg)
	msg := C.GoString(cMsg)

	C.foreman_error_getdetailmessage(cobj, &cDetailMsg)
	detailMsg := C.GoString(cDetailMsg)

	if (0 < len(msg)) || (0 < len(detailMsg)) {
		return fmt.Errorf("%s (%s)", msg, detailMsg)
	}

	if 0 < len(msg) {
		return fmt.Errorf("%s", msg, detailMsg)
	}

	if 0 < len(detailMsg) {
		return fmt.Errorf("%s", detailMsg)
	}

	return errors.New(ErrorClangNoMessage)
}
