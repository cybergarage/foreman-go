// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

type OperatorType int

const (
	OperatorTypeUnknown OperatorType = iota
	OperatorTypeLE
	OperatorTypeLT
	OperatorTypeGE
	OperatorTypeGT
	OperatorTypeEQ
	OperatorTypeNEQ
)

var operatorMap = map[string]OperatorType{
	QueryConditionOperatorLe:  OperatorTypeLE,
	QueryConditionOperatorLt:  OperatorTypeLT,
	QueryConditionOperatorGe:  OperatorTypeGE,
	QueryConditionOperatorGt:  OperatorTypeGT,
	QueryConditionOperatorEq:  OperatorTypeEQ,
	QueryConditionOperatorNeq: OperatorTypeNEQ,
}

// Operator represents an operaton in select query.
type Operator struct {
	operatorType OperatorType
}

// NewOperatorWithType creates a new operator with the specified type.
func NewOperatorWithType(opeType OperatorType) *Operator {
	op := &Operator{
		operatorType: opeType,
	}
	return op
}

// NewOperatorWithString creates a new operator with the specified string.
func NewOperatorWithString(opeStr string) *Operator {
	opeType, ok := operatorMap[opeStr]
	if !ok {
		opeType = OperatorTypeUnknown
	}
	return NewOperatorWithType(opeType)
}

// GetType returns the operator type
func (op *Operator) GetType() OperatorType {
	return op.operatorType
}

// String returns a string description of the instance
func (op *Operator) String() string {
	for opeStr, opeType := range operatorMap {
		if op.GetType() == opeType {
			return opeStr
		}
	}
	return ""
}
