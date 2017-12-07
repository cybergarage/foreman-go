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
	ParameterUnknownType ParameterType = iota
	ParameterIntegerType
	ParameterRealType
	ParameterBoolType
	ParameterStringType
)

const (
	parameterStringFormat = "%s (%d) : %s"
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
		Type:   ParameterIntegerType,
		iValue: value,
	}
	return param
}

// NewParameterFromReal returns a new real parameter.
func NewParameterFromReal(name string, value float64) *Parameter {
	param := &Parameter{
		Name:   name,
		Type:   ParameterRealType,
		rValue: value,
	}
	return param
}

// NewParameterFromBool returns a new boolean parameter.
func NewParameterFromBool(name string, value bool) *Parameter {
	param := &Parameter{
		Name:   name,
		Type:   ParameterBoolType,
		bValue: value,
	}
	return param
}

// NewParameterFromString returns a new string parameter.
func NewParameterFromString(name string, value string) *Parameter {
	param := &Parameter{
		Name:   name,
		Type:   ParameterStringType,
		sValue: value,
	}
	return param
}

// GetName returns a stored name.
func (param *Parameter) GetName() string {
	return param.Name
}

// Equals returns true when the specified parameter is same, otherwise false.
func (param *Parameter) Equals(other *Parameter) bool {
	if param.Name != other.Name {
		return false
	}

	if param.Type != other.Type {
		return false
	}

	switch param.Type {
	case ParameterIntegerType:
		if param.iValue != other.iValue {
			return false
		}
	case ParameterRealType:
		if param.rValue != other.rValue {
			return false
		}
	case ParameterBoolType:
		if param.bValue != other.bValue {
			return false
		}
	case ParameterStringType:
		if param.sValue != other.sValue {
			return false
		}
	}

	return true
}

// IsInteger returns whether the parameter type is integer.
func (param *Parameter) IsInteger() bool {
	if param.Type != ParameterIntegerType {
		return false
	}
	return true
}

// IsReal returns whether the parameter type is real.
func (param *Parameter) IsReal() bool {
	if param.Type != ParameterRealType {
		return false
	}
	return true
}

// IsBool returns whether the parameter type is boolean.
func (param *Parameter) IsBool() bool {
	if param.Type != ParameterBoolType {
		return false
	}
	return true
}

// IsString returns whether the parameter type is string.
func (param *Parameter) IsString() bool {
	if param.Type != ParameterStringType {
		return false
	}
	return true
}

// GetInteger returns a stored integer value.
func (param *Parameter) GetInteger() (int64, error) {
	if !param.IsInteger() {
		return 0, fmt.Errorf(errorInvalidParameterType, param.Type, ParameterIntegerType)
	}
	return param.iValue, nil
}

// GetReal returns a stored real value.
func (param *Parameter) GetReal() (float64, error) {
	if !param.IsReal() {
		return 0, fmt.Errorf(errorInvalidParameterType, param.Type, ParameterRealType)
	}
	return param.rValue, nil
}

// GetBool returns a stored boolean value.
func (param *Parameter) GetBool() (bool, error) {
	if !param.IsBool() {
		return false, fmt.Errorf(errorInvalidParameterType, param.Type, ParameterBoolType)
	}
	return param.bValue, nil
}

// GetString returns a stored integer value.
func (param *Parameter) GetString() (string, error) {
	if !param.IsString() {
		return "", fmt.Errorf(errorInvalidParameterType, param.Type, ParameterStringType)
	}
	return param.sValue, nil
}

// String returns a string description of the instance
func (param *Parameter) String() string {
	switch param.Type {
	case ParameterIntegerType:
		strValue := fmt.Sprintf("%d", param.iValue)
		return fmt.Sprintf(parameterStringFormat, param.Name, param.Type, strValue)
	case ParameterRealType:
		strValue := fmt.Sprintf("%f", param.rValue)
		return fmt.Sprintf(parameterStringFormat, param.Name, param.Type, strValue)
	case ParameterBoolType:
		strValue := "true"
		if !param.bValue {
			strValue = "false"
		}
		return fmt.Sprintf(parameterStringFormat, param.Name, param.Type, strValue)
	case ParameterStringType:
		return fmt.Sprintf(parameterStringFormat, param.Name, param.Type, param.sValue)
	}

	return ""
}
