// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -llua -lpython
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

// CObject returns a method objectfor foreman-cc.
func (method *Method) CObject() (unsafe.Pointer, error) {
	cmethod := C.foreman_action_method_new()
	if cmethod == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	C.foreman_action_method_setname(cmethod, C.CString(method.Name))
	C.foreman_action_method_setlanguage(cmethod, C.CString(method.Language))
	C.foreman_action_method_setstringcode(cmethod, C.CString(string(method.Code)))
	C.foreman_action_method_setencoding(cmethod, C.int(method.Encoding))

	return cmethod, nil
}
