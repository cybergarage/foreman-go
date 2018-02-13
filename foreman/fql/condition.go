// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"fmt"
)

const (
	operatorLeft  = 0
	operatorRight = 1
)

// Condition represents a condition in query.
type Condition struct {
	Operands []string
	Operator *Operator
}

// NewCondition returns a blank operator.
func NewCondition() *Condition {
	c := &Condition{
		Operands: []string{"", ""},
		Operator: NewOperatorWithType(OperatorTypeUnknown),
	}
	return c
}

// NewConditionWithObjects returns a new operator with the specified string objects.
func NewConditionWithObjects(condStrings []string) *Condition {
	if len(condStrings) < 3 {
		return NewCondition()
	}
	c := &Condition{
		Operands: []string{condStrings[0], condStrings[2]},
		Operator: NewOperatorWithString(condStrings[1]),
	}
	return c
}

// GetColumn returns the left operand string.
func (c *Condition) GetColumn() string {
	return c.Operands[operatorLeft]
}

// GetOperand returns the right operand string.
func (c *Condition) GetOperand() string {
	return c.Operands[operatorRight]
}

// GetOperator returns the operator.
func (c *Condition) GetOperator() *Operator {
	return c.Operator
}

// String returns a string description of the instance
func (c *Condition) String() string {
	return fmt.Sprintf("%s %s %s", c.Operands[0], c.Operator.String(), c.Operands[1])
}
