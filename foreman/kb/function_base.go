// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import "fmt"

// BaseFunction represents a simple operand object which has a name and value.
type BaseFunction struct {
	Name   string
	Params []interface{}
}

// NewFunction returns a new BaseFunction instance.
func NewFunction() *BaseFunction {
	return NewFunctionWithName("")
}

// NewFunctionWithName returns a new BaseFunction instance with the specified name.
func NewFunctionWithName(name string) *BaseFunction {
	return NewFunctionWithNameAndParams(name, make([]interface{}, 0))
}

// NewFunctionWithNameAndParams returns a new base function instance with the specified name and parameters.
func NewFunctionWithNameAndParams(name string, params []interface{}) *BaseFunction {
	fn := &BaseFunction{
		Name:   name,
		Params: nil,
	}
	fn.SetParameters(params)
	return fn
}

// SetName sets a specified name.
func (fn *BaseFunction) SetName(name string) {
	fn.Name = name
}

// GetName returns the name.
func (fn *BaseFunction) GetName() string {
	return fn.Name
}

// SetParameters sets a specified parameters.
func (fn *BaseFunction) SetParameters(params []interface{}) {
	fn.Params = make([]interface{}, len(params))
	copy(fn.Params, params)
}

// GetParameters returns the parameters.
func (fn *BaseFunction) GetParameters() []interface{} {
	return fn.Params
}

// GetValue returns the stored value.
func (fn *BaseFunction) GetValue() (interface{}, error) {
	return fn.Execute(fn.Params)
}

// Execute returns the operand value with the specified parameters.
func (fn *BaseFunction) Execute([]interface{}) (interface{}, error) {
	return nil, fmt.Errorf(errorUnknownFunction, fn.Name)
}

// Expression returns a string for the formula expression.
func (fn *BaseFunction) Expression() string {
	return fmt.Sprintf("%s", fn.Name)
}
