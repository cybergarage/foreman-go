// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"fmt"
)

// ParameterType defines parameter types
type ParameterType int

const (
	Unknown ParameterType = iota
	IntegerType
	RealType
	BoolType
	StringType
)

// Parameter represents a parameter for action method.
type Parameter struct {
	Name   string
	Type   ParameterType
	iValue int64
	rValue float64
	sValue string
	bValue bool
}

// NewParameterFromInteger returns a new integer parameter.
func NewParameterFromInteger(name string, value int64) *Parameter {
	param := &Parameter{
		Name:   name,
		Type:   IntegerType,
		iValue: value,
	}
	return param
}

// NewParameterFromReal returns a new real parameter.
func NewParameterFromReal(name string, value float64) *Parameter {
	param := &Parameter{
		Name:   name,
		Type:   RealType,
		rValue: value,
	}
	return param
}

// NewParameterFromBool returns a new boolean parameter.
func NewParameterFromBool(name string, value bool) *Parameter {
	param := &Parameter{
		Name:   name,
		Type:   BoolType,
		bValue: value,
	}
	return param
}

// NewParameterFromString returns a new string parameter.
func NewParameterFromString(name string, value string) *Parameter {
	param := &Parameter{
		Name:   name,
		Type:   StringType,
		sValue: value,
	}
	return param
}

// GetName returns a stored name.
func (param *Parameter) GetName() string {
	return param.Name
}

// IsInteger returns whether the parameter type is integer.
func (param *Parameter) IsInteger() bool {
	if param.Type != IntegerType {
		return false
	}
	return true
}

// IsReal returns whether the parameter type is real.
func (param *Parameter) IsReal() bool {
	if param.Type != RealType {
		return false
	}
	return true
}

// IsBool returns whether the parameter type is boolean.
func (param *Parameter) IsBool() bool {
	if param.Type != BoolType {
		return false
	}
	return true
}

// IsString returns whether the parameter type is string.
func (param *Parameter) IsString() bool {
	if param.Type != StringType {
		return false
	}
	return true
}

// GetInteger returns a stored integer value.
func (param *Parameter) GetInteger() (int64, error) {
	if !param.IsInteger() {
		return 0, fmt.Errorf(errorInvalidType, param.Type, IntegerType)
	}
	return param.iValue, nil
}

// GetReal returns a stored real value.
func (param *Parameter) GetReal() (float64, error) {
	if !param.IsReal() {
		return 0, fmt.Errorf(errorInvalidType, param.Type, RealType)
	}
	return param.rValue, nil
}

// GetBool returns a stored boolean value.
func (param *Parameter) GetBool() (bool, error) {
	if !param.IsBool() {
		return false, fmt.Errorf(errorInvalidType, param.Type, BoolType)
	}
	return param.bValue, nil
}

// GetString returns a stored integer value.
func (param *Parameter) GetString() (string, error) {
	if !param.IsString() {
		return "", fmt.Errorf(errorInvalidType, param.Type, StringType)
	}
	return param.sValue, nil
}
