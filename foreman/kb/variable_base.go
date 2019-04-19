// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import "fmt"

// BaseVariable represents a simple operand object which has a name and value.
type BaseVariable struct {
	Name  string
	Value interface{}
}

// NewVariable returns a new BaseVariable instance.
func NewVariable() *BaseVariable {
	return NewVariableWithName("")
}

// NewVariableWithName returns a new BaseVariable instance with the specified name.
func NewVariableWithName(name string) *BaseVariable {
	return NewVariableWithNameAndValue(name, nil)

}

// NewVariableWithNameAndValue returns a new BaseVariable instance with the specified name and value.
func NewVariableWithNameAndValue(name string, value interface{}) *BaseVariable {
	v := &BaseVariable{
		Name:  name,
		Value: value,
	}
	return v
}

// SetName sets a specified name.
func (v *BaseVariable) SetName(name string) {
	v.Name = name
}

// GetName returns the name.
func (v *BaseVariable) GetName() string {
	return v.Name
}

// SetValue sets a specified value.
func (v *BaseVariable) SetValue(value interface{}) error {
	v.Value = value
	return nil
}

// GetValue returns the stored value.
func (v *BaseVariable) GetValue() (interface{}, error) {
	return v.Value, nil
}

// Expression returns a string for the formula expression.
func (v *BaseVariable) Expression() string {
	return fmt.Sprintf("%s", v.Name)
}
