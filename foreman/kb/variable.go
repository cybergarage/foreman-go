// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import "fmt"

// Variable represents a simple operand object which has a name and value.
type Variable struct {
	Object
	Name  string
	Value interface{}
}

// NewVariable returns a new variable instance.
func NewVariable() *Variable {
	return NewVariableWithName("")
}

// NewVariableWithName returns a new variable instance with the specified name.
func NewVariableWithName(name string) *Variable {
	return NewVariableWithNameAndValue(name, nil)

}

// NewVariableWithNameAndValue returns a new variable instance with the specified name and value.
func NewVariableWithNameAndValue(name string, value interface{}) *Variable {
	v := &Variable{
		Name:  name,
		Value: value,
	}
	return v
}

// SetName sets a specified name.
func (v *Variable) SetName(name string) {
	v.Name = name
}

// GetName returns the name.
func (v *Variable) GetName() string {
	return v.Name
}

// SetValue sets a specified value.
func (v *Variable) SetValue(value interface{}) error {
	v.Value = value
	return nil
}

// GetValue returns the stored value.
func (v *Variable) GetValue() (interface{}, error) {
	return v.Value, nil
}

// Expression returns a string for the formula expression.
func (v *Variable) Expression() string {
	return fmt.Sprintf("%s", v.Name)
}
