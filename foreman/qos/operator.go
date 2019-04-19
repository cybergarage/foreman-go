// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/kb"
)

// Operator represents an operator interface in qualities.
type Operator struct {
	kb.Operator
	Type int
}

// NewOperatorWithType returns a new operator with the specified type.
func NewOperatorWithType(typeID int) *Operator {
	operator := &Operator{
		Type: typeID,
	}
	return operator
}

// NewOperatorWithString returns a new operator with the specified string.
func NewOperatorWithString(typeString string) (*Operator, error) {
	operator := &Operator{}
	err := operator.SetTypeString(typeString)
	if err != nil {
		return nil, err
	}
	return operator, nil
}

// SetTypeString update the current type with the specified string.
func (operator *Operator) SetTypeString(typeString string) error {
	operatorType, ok := qosOperatorTypes[typeString]
	if ok {
		operator.Type = operatorType
		return nil
	}

	return fmt.Errorf(errorInvalidOperator, typeString)
}

// IsSatisfied is an interface method of kb.Operator
func (operator *Operator) IsSatisfied(v kb.Operand, obj kb.Operand) (bool, error) {
	m, ok := v.(*Metric)
	if !ok {
		return false, fmt.Errorf(errorInvalidOperand, v)
	}
	mObj, err := m.GetValue()
	if err != nil {
		return false, err
	}
	mValue, ok := mObj.(float64)
	if !ok {
		return false, fmt.Errorf(errorInvalidOperand, v)
	}

	th, ok := obj.(*Threshold)
	if !ok {
		return false, fmt.Errorf(errorInvalidOperand, obj)
	}
	thObj, err := th.GetValue()
	if err != nil {
		return false, err
	}
	thValue, ok := thObj.(float64)
	if !ok {
		return false, fmt.Errorf(errorInvalidOperand, v)
	}

	switch operator.Type {
	case qosOperatorTypeLT:
		if mValue < thValue {
			return true, nil
		}
		return false, nil
	case qosOperatorTypeLE:
		if mValue <= thValue {
			return true, nil
		}
		return false, nil
	case qosOperatorTypeGT:
		if mValue > thValue {
			return true, nil
		}
		return false, nil
	case qosOperatorTypeGE:
		if mValue >= thValue {
			return true, nil
		}
		return false, nil
	case qosOperatorTypeEQ:
		if mValue == thValue {
			return true, nil
		}
		return false, nil
	case qosOperatorTypeNotEQ:
		if mValue != thValue {
			return true, nil
		}
		return false, nil
	}

	return false, fmt.Errorf(errorInvalidOperator, operator)
}

// Expression returns a expression string in the formula
func (operator *Operator) Expression() string {
	for operatorString, operatorType := range qosOperatorTypes {
		if operator.Type == operatorType {
			return operatorString
		}
	}
	return ""
}
