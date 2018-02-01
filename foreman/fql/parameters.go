// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// Parameters represents a parameter map.
type Parameters map[string]Parameter

// NewParameters returns a new parameter map.
func NewParameters() Parameters {
	params := make(Parameters)
	return params
}

// AddParameter adds a new parameter.
func (params Parameters) AddParameter(param Parameter) error {
	params[param.GetName()] = param
	return nil
}

// GetParameter returns a specified parameter.
func (params Parameters) GetParameter(name string) (Parameter, bool) {
	param, ok := params[name]
	return param, ok
}

// GetParameterString returns a string of the specified parameter.
func (params Parameters) GetParameterString(name string) (string, bool) {
	param, ok := params.GetParameter(name)
	if !ok {
		return "", false
	}
	return param.GetString()
}
