// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import (
	"strings"
)

// Rule represents a QoS rule.
type Rule struct {
	Clauses []Clause
}

// NewRule returns a new null rule.
func NewRule() *Rule {
	p := &Rule{
		Clauses: make([]Clause, 0),
	}
	return p
}

// ParseString parses a specified rule string.
func (rule *Rule) ParseString(factory Factory, ruleString string) error {
	clausesString := strings.Split(ruleString, ClauseSeparator)

	for _, clauseString := range clausesString {
		clause := NewClause()
		clauseString = strings.Trim(clauseString, (StartBracket + EndBracket))
		formulasString := strings.Split(clauseString, FormulaSeparator)
		for _, formulaString := range formulasString {
			formula := NewFormula()
			formulaString = strings.Trim(formulaString, (StartBracket + EndBracket))
			err := formula.ParseString(factory, formulaString)
			if err != nil {
				return err
			}
			err = clause.AddFormula(formula)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
