// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import (
	"fmt"
)

// BaseFormula represents a base formula.
type BaseFormula struct {
	LeftOperand  Operand
	Operator     Operator
	RightOperand Operand
}

// NewFormulaWithParams returns a new base formula with the specified parameters.
func NewFormulaWithParams(lop Operand, op Operator, rop Operand) *BaseFormula {
	formula := &BaseFormula{
		LeftOperand:  lop,
		Operator:     op,
		RightOperand: rop,
	}
	return formula
}

// NewFormula returns a new base formula.
func NewFormula() *BaseFormula {
	return NewFormulaWithParams(nil, nil, nil)
}

// GetLeftOperand returns a left operand object
func (formula *BaseFormula) GetLeftOperand() Operand {
	return formula.LeftOperand
}

// GetOperator returns an operator of the object
func (formula *BaseFormula) GetOperator() Operator {
	return formula.Operator
}

// GetRightOperand returns a right operand object
func (formula *BaseFormula) GetRightOperand() Operand {
	return formula.RightOperand
}

// GetOperands returns all operand objects
func (formula *BaseFormula) GetOperands() []Operand {
	return []Operand{formula.LeftOperand, formula.RightOperand}
}

// IsSatisfied checks whether the formula is valid
func (formula *BaseFormula) IsSatisfied() (bool, error) {
	return formula.Operator.IsSatisfied(formula.LeftOperand, formula.RightOperand)
}

// String returns a string description of the instance
func (formula *BaseFormula) String() string {
	return fmt.Sprintf("%s%s %s %s%s", StartBracket, formula.LeftOperand.Expression(), formula.Operator.Expression(), formula.RightOperand.Expression(), EndBracket)
}
