// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import (
	"fmt"
	"strings"
)

// A Formula represents a formula formula in a QoS clause.
type Formula struct {
	Variable  Variable
	Operator  Operator
	Objective Objective
}

// NewFormula returns a new formula.
func NewFormula() *Formula {
	formula := &Formula{
		Variable: nil,
	}
	return formula
}

// IsSatisfied checks whether the formula is valid
func (formula *Formula) IsSatisfied() (bool, error) {
	return formula.Operator.IsSatisfied(formula.Variable, formula.Objective)
}

// ParseString parses a specified formula string.
func (formula *Formula) ParseString(factory Factory, formulaString string) error {
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
	formula.Objective = objective

	return nil
}

// String returns a string description of the instance
func (formula *Formula) String() string {
	return fmt.Sprintf("%s%s %s %s%s", StartBracket, formula.Variable.GetName(), formula.Operator.String(), formula.Objective.String(), EndBracket)
}
