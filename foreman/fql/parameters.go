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
