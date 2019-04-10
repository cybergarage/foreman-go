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
	Variable     Variable
	Operator     Operator
	RightOperand Operand
}

// NewFormula returns a new base formula.
func NewFormula() *BaseFormula {
	formula := &BaseFormula{
		Variable: nil,
	}
	return formula
}

// GetVariable returns a variable of the object
func (formula *BaseFormula) GetVariable() Variable {
	return formula.Variable
}

// GetOperator returns an operator of the object
func (formula *BaseFormula) GetOperator() Operator {
	return formula.Operator
}

// GetRightOperand returns a right operand object
func (formula *BaseFormula) GetRightOperand() Operand {
	return formula.RightOperand
}

// IsSatisfied checks whether the formula is valid
func (formula *BaseFormula) IsSatisfied() (bool, error) {
	return formula.Operator.IsSatisfied(formula.Variable, formula.RightOperand)
}

// ParseString parses a specified formula string.
func (formula *BaseFormula) ParseString(factory Factory, formulaString string) error {
	formulaStrings := strings.Split(formulaString, SpaceSeparator)

	if len(formulaStrings) != 3 {
		return fmt.Errorf(errorInvalidFormulaString, formulaString)
	}

	// Variable
	variableString := formulaStrings[0]
	variable, err := factory.CreateVariable(variableString)
	if err != nil {
		return err
	}
	formula.Variable = variable

	// Operator
	operatorString := formulaStrings[1]
	operator, err := factory.CreateOperator(operatorString)
	if err != nil {
		return err
	}
	formula.Operator = operator

	// Objective
	objectiveString := formulaStrings[2]
	objective, err := factory.CreateObjective(objectiveString)
	if err != nil {
		return err
	}
	formula.RightOperand = objective

	return nil
}

// String returns a string description of the instance
func (formula *BaseFormula) String() string {
	return fmt.Sprintf("%s%s %s %s%s", StartBracket, formula.Variable.GetName(), formula.Operator.String(), formula.RightOperand.String(), EndBracket)
}
