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

// NewParametersWithCObject returns a new parameters map of the specified C++ object.
func NewParametersWithCObject(cObject unsafe.Pointer) Parameters {
	params := NewParameters()
	params.SetCObject(cObject)
	return params
}

// SetCObject sets parametes which are in the specified C++ object.
func (params Parameters) SetCObject(cObject unsafe.Pointer) error {
	paramCount := int(C.foreman_action_parameters_size(cObject))
	for n := 0; n < paramCount; n++ {
		cparam := C.foreman_action_parameters_getparameter(cObject, C.size_t(n))
		if cparam == nil {
			continue
		}
		param := NewParameterWithCObject(cparam)
		if param == nil {
			continue
		}
		params.AddParameter(param)
	}

	return nil
}

// CObject returns a parameters object for foreman-cc.
func (params Parameters) CObject() (unsafe.Pointer, error) {
	cparams := C.foreman_action_parameters_new()
	if cparams == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	for _, param := range params {
		cparam, err := param.CObject()
		if err != nil {
			return nil, err
		}
		if param == nil {
			return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
		}
		C.foreman_action_parameters_addparameter(cparams, cparam)
	}

	return cparams, nil
}
