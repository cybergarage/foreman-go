// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"fmt"
	"strconv"
)

const (
	parameterStringFormat = "%s (%s) = %s"
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

// parameterTypeToString returns a string description of the specified parameter type
func parameterTypeToString(paramType ParameterType) string {
	paramTypeStrings := map[ParameterType]string{
		ParameterIntegerType: "int",
		ParameterRealType:    "real",
		ParameterBoolType:    "bool",
		ParameterStringType:  "string",
	}

	paramTypeString, ok := paramTypeStrings[paramType]
	if !ok {
		return "?"
	}

	return paramTypeString
}

// Parameter represents a parameter for action method.
type Parameter struct {
	Name  string
	Type  ParameterType
	Value interface{}
}

// NewParameterFromInteger returns a new integer parameter.
func NewParameterFromInteger(name string, value int64) *Parameter {
	param := &Parameter{
		Name:  name,
		Type:  ParameterIntegerType,
		Value: value,
	}
	return param
}

// NewParameterFromReal returns a new real parameter.
func NewParameterFromReal(name string, value float64) *Parameter {
	param := &Parameter{
		Name:  name,
		Type:  ParameterRealType,
		Value: value,
	}
	return param
}

// NewParameterFromBool returns a new boolean parameter.
func NewParameterFromBool(name string, value bool) *Parameter {
	param := &Parameter{
		Name:  name,
		Type:  ParameterBoolType,
		Value: value,
	}
	return param
}

// NewParameterFromString returns a new string parameter.
func NewParameterFromString(name string, value string) *Parameter {
	ivalue, err := strconv.ParseInt(value, 10, 64)
	if err == nil {
		return NewParameterFromInteger(name, ivalue)
	}

	fvalue, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return NewParameterFromReal(name, fvalue)
	}

	param := &Parameter{
		Name:  name,
		Type:  ParameterStringType,
		Value: value,
	}
	return param
}

// NewParameterFromInterface returns a new parameter.
func NewParameterFromInterface(name string, valueObj interface{}) (*Parameter, error) {
	switch valueObj.(type) {
	case string:
		value, _ := valueObj.(string)
		return NewParameterFromString(name, value), nil
	case bool:
		value, _ := valueObj.(bool)
		return NewParameterFromBool(name, value), nil
	case int:
		value, _ := valueObj.(int)
		return NewParameterFromInteger(name, int64(value)), nil
	case int8:
		value, _ := valueObj.(int8)
		return NewParameterFromInteger(name, int64(value)), nil
	case int16:
		value, _ := valueObj.(int16)
		return NewParameterFromInteger(name, int64(value)), nil
	case int32:
		value, _ := valueObj.(int32)
		return NewParameterFromInteger(name, int64(value)), nil
	case int64:
		value, _ := valueObj.(int64)
		return NewParameterFromInteger(name, value), nil
	case float32:
		value, _ := valueObj.(float32)
		return NewParameterFromReal(name, float64(value)), nil
	case float64:
		value, _ := valueObj.(float64)
		return NewParameterFromReal(name, value), nil
	}

	return nil, fmt.Errorf(errorUnknownParameterValueType, valueObj)
}

// GetName returns a stored name.
func (param *Parameter) GetName() string {
	return param.Name
}

// GetType returns a stored data type.
func (param *Parameter) GetType() ParameterType {
	return param.Type
}

// GetValue returns a stored value.
func (param *Parameter) GetValue() interface{} {
	return param.Value
}

// Equals returns true when the specified parameter is same, otherwise false.
func (param *Parameter) Equals(other *Parameter) bool {
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
	v, ok := param.Value.(int64)
	if !ok {
		return 0, fmt.Errorf(errorInvalidParameterType, param.Type, ParameterIntegerType)
	}
	return v, nil
}

// GetReal returns a stored real value.
func (param *Parameter) GetReal() (float64, error) {
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
func (param *Parameter) GetBool() (bool, error) {
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
func (param *Parameter) GetString() (string, error) {
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
func (param *Parameter) String() string {
	switch param.Type {
	case ParameterIntegerType:
		value, _ := param.GetInteger()
		strValue := fmt.Sprintf("%d", value)
		return fmt.Sprintf(parameterStringFormat, param.Name, parameterTypeToString(param.Type), strValue)
	case ParameterRealType:
		value, _ := param.GetReal()
		strValue := fmt.Sprintf("%f", value)
		return fmt.Sprintf(parameterStringFormat, param.Name, parameterTypeToString(param.Type), strValue)
	case ParameterBoolType:
		value, _ := param.GetBool()
		strValue := "true"
		if value {
			strValue = "false"
		}
		return fmt.Sprintf(parameterStringFormat, param.Name, parameterTypeToString(param.Type), strValue)
	case ParameterStringType:
		strValue, _ := param.GetString()
		return fmt.Sprintf(parameterStringFormat, param.Name, parameterTypeToString(param.Type), strValue)
	}

	return ""
}
