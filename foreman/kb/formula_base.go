// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import (
	"fmt"
	"strings"
)

// BaseFormula represents a base formula.
type BaseFormula struct {
	LeftOperand  Operand
	Operator     Operator
	RightOperand Operand
}

// NewFormula returns a new base formula.
func NewFormula() *BaseFormula {
	formula := &BaseFormula{
		LeftOperand:  nil,
		Operator:     nil,
		RightOperand: nil,
	}
	return formula
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

// ParseString parses a specified formula string.
func (formula *BaseFormula) ParseString(factory Factory, formulaString string) error {
	formulaStrings := strings.Split(formulaString, SpaceSeparator)

	if len(formulaStrings) != 3 {
		return fmt.Errorf(errorInvalidFormulaString, formulaString)
	}

	// Left Operand
	operand, err := factory.CreateOperand(formulaStrings[0])
	if err != nil {
		return err
	}
	formula.LeftOperand = operand

	// Operator
	operator, err := factory.CreateOperator(formulaStrings[1])
	if err != nil {
		return err
	}
	formula.Operator = operator

	// Right Operand
	operand, err = factory.CreateOperand(formulaStrings[2])
	if err != nil {
		return err
	}
	formula.RightOperand = operand

	return nil
}

// String returns a string description of the instance
func (formula *BaseFormula) String() string {
	return fmt.Sprintf("%s%s %s %s%s", StartBracket, formula.LeftOperand.Expression(), formula.Operator.Expression(), formula.RightOperand.Expression(), EndBracket)
}
