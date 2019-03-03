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

// NewParameterWithCObject returns a new parameter of the specified C++ object.
func NewParameterWithCObject(cObject unsafe.Pointer) *Parameter {
	name := C.GoString(C.foreman_action_parameter_getname(cObject))

	if C.foreman_action_parameter_isinteger(cObject) {
		return NewParameterFromInteger(name, int64(C.foreman_action_parameter_getinteger(cObject)))
	}
	if C.foreman_action_parameter_isreal(cObject) {
		return NewParameterFromReal(name, float64(C.foreman_action_parameter_getreal(cObject)))
	}
	if C.foreman_action_parameter_isbool(cObject) {
		return NewParameterFromBool(name, bool(C.foreman_action_parameter_getbool(cObject)))
	}
	if C.foreman_action_parameter_isstring(cObject) {
		return NewParameterFromString(name, C.GoString(C.foreman_action_parameter_getstring(cObject)))
	}

	return nil
}

// CObject returns a parameter object for foreman-cc.
func (param *Parameter) CObject() (unsafe.Pointer, error) {
	var cparam unsafe.Pointer

	switch param.Type {
	case ParameterIntegerType:
		cparam = C.foreman_action_parameter_integer_new()
		if cparam == nil {
			return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
		}
		value, _ := param.GetInteger()
		C.foreman_action_parameter_setinteger(cparam, C.long(value))
	case ParameterRealType:
		cparam = C.foreman_action_parameter_real_new()
		if cparam == nil {
			return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
		}
		value, _ := param.GetReal()
		C.foreman_action_parameter_setreal(cparam, C.double(value))
	case ParameterBoolType:
		cparam = C.foreman_action_parameter_bool_new()
		if cparam == nil {
			return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
		}
		value, _ := param.GetBool()
		C.foreman_action_parameter_setbool(cparam, C.bool(value))
	case ParameterStringType:
		cparam = C.foreman_action_parameter_string_new()
		if cparam == nil {
			return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
		}
		value, _ := param.GetString()
		cvalue := C.CString(value)
		defer C.free(unsafe.Pointer(cvalue))
		C.foreman_action_parameter_setstring(cparam, cvalue)
	}

	if cparam == nil {
		return nil, fmt.Errorf(errorUnknownParameterType, param.Type)
	}

	cname := C.CString(param.Name)
	defer C.free(unsafe.Pointer(cname))

	C.foreman_action_parameter_setname(cparam, cname)

	return cparam, nil
}
