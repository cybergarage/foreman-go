// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// Parameters represents a parameter map.
type Parameters map[string]Parameter

// NewParameters returns a null parameter map.
func NewParameters() Parameters {
	params := make(Parameters)
	return params
}

// AddParameter adds a new parameter.
func (params Parameters) AddParameter(param Parameter) error {
	params[param.GetName()] = param
	return nil
}

// Equals returns true when the specified parameters are same, otherwise false.
func (params Parameters) Equals(others Parameters) bool {
	if len(params) != len(others) {
		return false
	}

	for name, param := range params {
		oparam, ok := others[name]
		if !ok {
			return false
		}

		if !param.Equals(oparam) {
			return false
		}
	}

	return true
}
