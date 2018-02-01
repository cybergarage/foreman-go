// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

const (
	parameterStringFormat = "%s (%d) : %s"
)

// BaseParameter represents a parameter.
type BaseParameter struct {
	Name  string
	Value interface{}
}

// NewParameterWithObject returns a new parameter.
func NewParameterWithObject(name string, value interface{}) *BaseParameter {
	param := &BaseParameter{
		Name:  name,
		Value: value,
	}
	return param
}

// GetName returns a stored name.
func (param *BaseParameter) GetName() string {
	return param.Name
}

// GetValue returns a stored value.
func (param *BaseParameter) GetValue() interface{} {
	return param.Value
}

// GetString returns a stored string value.
func (param *BaseParameter) GetString() (string, bool) {
	value, ok := param.Value.(string)
	return value, ok
}
