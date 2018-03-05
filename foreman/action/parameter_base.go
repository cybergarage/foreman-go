// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"fmt"
	"strconv"
)

const (
	parameterStringFormat = "%s (%d) : %s"
)

// BaseParameter represents a parameter for action method.
type BaseParameter struct {
	Name  string
	Type  ParameterType
	Value interface{}
}

// NewParameterFromInteger returns a new integer parameter.
func NewParameterFromInteger(name string, value int64) *BaseParameter {
	param := &BaseParameter{
		Name:  name,
		Type:  ParameterIntegerType,
		Value: value,
	}
	return param
}

// NewParameterFromReal returns a new real parameter.
func NewParameterFromReal(name string, value float64) *BaseParameter {
	param := &BaseParameter{
		Name:  name,
		Type:  ParameterRealType,
		Value: value,
	}
	return param
}

// NewParameterFromBool returns a new boolean parameter.
func NewParameterFromBool(name string, value bool) *BaseParameter {
	param := &BaseParameter{
		Name:  name,
		Type:  ParameterBoolType,
		Value: value,
	}
	return param
}

// NewParameterFromString returns a new string parameter.
func NewParameterFromString(name string, value string) *BaseParameter {
	ivalue, err := strconv.ParseInt(value, 10, 64)
	if err == nil {
		return NewParameterFromInteger(name, ivalue)
	}

	fvalue, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return NewParameterFromReal(name, fvalue)
	}

	param := &BaseParameter{
		Name:  name,
		Type:  ParameterStringType,
		Value: value,
	}
	return param
}

// GetName returns a stored name.
func (param *BaseParameter) GetName() string {
	return param.Name
}

// GetType returns a stored data type.
func (param *BaseParameter) GetType() ParameterType {
	return param.Type
}

// GetValue returns a stored value.
func (param *BaseParameter) GetValue() interface{} {
	return param.Value
}

// Equals returns true when the specified parameter is same, otherwise false.
func (param *BaseParameter) Equals(other Parameter) bool {
	if param.GetName() != other.GetName() {
		return false
	}

	if param.GetType() != other.GetType() {
		return false
	}

	switch param.GetType() {
	case ParameterIntegerType:
		pv, _ := param.Value.(int64)
		ov, _ := other.GetValue().(int64)
		if pv != ov {
			return false
		}
	case ParameterRealType:
		pv, _ := param.Value.(float64)
		ov, _ := other.GetValue().(float64)
		if pv != ov {
			return false
		}
	case ParameterBoolType:
		pv, _ := param.Value.(bool)
		ov, _ := other.GetValue().(bool)
		if pv != ov {
			return false
		}
	case ParameterStringType:
		pv, _ := param.Value.(string)
		ov, _ := other.GetValue().(string)
		if pv != ov {
			return false
		}
	}

	return true
}

// IsInteger returns whether the parameter type is integer.
func (param *BaseParameter) IsInteger() bool {
	if param.Type != ParameterIntegerType {
		return false
	}
	return true
}

// IsReal returns whether the parameter type is real.
func (param *BaseParameter) IsReal() bool {
	if param.Type != ParameterRealType {
		return false
	}
	return true
}

// IsBool returns whether the parameter type is boolean.
func (param *BaseParameter) IsBool() bool {
	if param.Type != ParameterBoolType {
		return false
	}
	return true
}

// IsString returns whether the parameter type is string.
func (param *BaseParameter) IsString() bool {
	if param.Type != ParameterStringType {
		return false
	}
	return true
}

// GetInteger returns a stored integer value.
func (param *BaseParameter) GetInteger() (int64, error) {
	if !param.IsInteger() {
		return 0, fmt.Errorf(errorInvalidParameterType, param.Type, ParameterIntegerType)
	}
	v, ok := param.Value.(int64)
	if !ok {
		return 0, fmt.Errorf(errorInvalidParameterType, param.Type, ParameterIntegerType)
	}
	return v, nil
}

// GetReal returns a stored real value.
func (param *BaseParameter) GetReal() (float64, error) {
	if !param.IsReal() {
		return 0, fmt.Errorf(errorInvalidParameterType, param.Type, ParameterRealType)
	}
	v, ok := param.Value.(float64)
	if !ok {
		return 0, fmt.Errorf(errorInvalidParameterType, param.Type, ParameterIntegerType)
	}
	return v, nil
}

// GetBool returns a stored boolean value.
func (param *BaseParameter) GetBool() (bool, error) {
	if !param.IsBool() {
		return false, fmt.Errorf(errorInvalidParameterType, param.Type, ParameterBoolType)
	}
	v, ok := param.Value.(bool)
	if !ok {
		return false, fmt.Errorf(errorInvalidParameterType, param.Type, ParameterBoolType)
	}
	return v, nil
}

// GetString returns a stored integer value.
func (param *BaseParameter) GetString() (string, error) {
	if !param.IsString() {
		return "", fmt.Errorf(errorInvalidParameterType, param.Type, ParameterStringType)
	}
	v, ok := param.Value.(string)
	if !ok {
		return "", fmt.Errorf(errorInvalidParameterType, param.Type, ParameterStringType)
	}
	return v, nil
}

// String returns a string description of the instance
func (param *BaseParameter) String() string {
	switch param.Type {
	case ParameterIntegerType:
		value, _ := param.GetInteger()
		strValue := fmt.Sprintf("%d", value)
		return fmt.Sprintf(parameterStringFormat, param.Name, param.Type, strValue)
	case ParameterRealType:
		value, _ := param.GetReal()
		strValue := fmt.Sprintf("%f", value)
		return fmt.Sprintf(parameterStringFormat, param.Name, param.Type, strValue)
	case ParameterBoolType:
		value, _ := param.GetBool()
		strValue := "true"
		if value {
			strValue = "false"
		}
		return fmt.Sprintf(parameterStringFormat, param.Name, param.Type, strValue)
	case ParameterStringType:
		strValue, _ := param.GetString()
		return fmt.Sprintf(parameterStringFormat, param.Name, param.Type, strValue)
	}

	return ""
}
