// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import (
	"fmt"
	"strings"
)

// BaseFunction represents a simple operand object which has a name and value.
type BaseFunction struct {
	Name     string
	Params   []interface{}
	Executor Function
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
		Name:     name,
		Params:   nil,
		Executor: nil,
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

// AddParameter adds a specified parameters.
func (fn *BaseFunction) AddParameter(param interface{}) {
	fn.Params = append(fn.Params, param)
}

// GetParameters returns the parameters.
func (fn *BaseFunction) GetParameters() []interface{} {
	return fn.Params
}

// HasVariable returns true when the function has the specified parameter, otherwise false.
func (fn *BaseFunction) HasVariable(name string) bool {
	for _, param := range fn.Params {
		sparam, ok := param.(string)
		if !ok {
			continue
		}
		if sparam != name {
			continue
		}
		return true
	}
	return false
}

// GetVariableNames returns only parameter name in parameter.
func (fn *BaseFunction) GetVariableNames() []string {
	names := []string{}
	for _, param := range fn.Params {
		name, ok := param.(string)
		if !ok {
			continue
		}
		names = append(names, name)
	}
	return names
}

// SetExecutor sets a specified name.
func (fn *BaseFunction) SetExecutor(executor Function) {
	fn.Executor = executor
}

// GetValue returns the stored value.
func (fn *BaseFunction) GetValue() (interface{}, error) {
	if fn.Executor == nil {
		return nil, fmt.Errorf(errorUnknownFunction, fn.Name)
	}
	return fn.Executor.Execute(fn.Params)
}

// ExpressionWithParameters returns a string for the expression with the parameters.
func (fn *BaseFunction) ExpressionWithParameters(params []interface{}) string {
	paramStrings := make([]string, len(params))
	for n, param := range params {
		paramStrings[n] = fmt.Sprintf("%s", param)
	}
	return fmt.Sprintf("%s(%s)", fn.Name, strings.Join(paramStrings, ", "))
}

// Expression returns a string for the expression.
func (fn *BaseFunction) Expression() string {
	return fn.ExpressionWithParameters(fn.Params)
}
