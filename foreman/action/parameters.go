// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import "fmt"

// Parameters represents a parameter map.
type Parameters map[string]*Parameter

// NewParameters returns a new parameter map.
func NewParameters() Parameters {
	params := make(Parameters)
	return params
}

// AddParameter adds a new parameter.
func (params Parameters) AddParameter(param *Parameter) error {
	params[param.GetName()] = param
	return nil
}

// Equals returns true when the specified parameters are same, otherwise false.
func (params Parameters) Equals(others Parameters) bool {
	if len(params) != len(others) {
		return false
	}

	for name, param := range params {
		otherParam, ok := others[name]
		if !ok {
			return false
		}

		if !param.Equals(otherParam) {
			return false
		}
	}

	return true
}

// String returns a string description of the instance
func (params Parameters) String() string {
	paramsStr := ""

	paramIdx := 0
	for _, param := range params {
		if 0 < paramIdx {
			paramsStr += ", "
		}
		paramsStr += fmt.Sprintf("{%s}", param.String())
		paramIdx++
	}

	return paramsStr
}
