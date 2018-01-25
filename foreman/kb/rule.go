// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import (
	"bytes"
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

// AddClause adds a new clause.
func (rule *Rule) AddClause(clause Clause) error {
	rule.Clauses = append(rule.Clauses, clause)
	return nil
}

// ParseString parses a specified rule string.
func (rule *Rule) ParseString(factory Factory, ruleString string) error {
	clausesString := strings.Split(ruleString, ClauseSeparator)

	for _, clauseString := range clausesString {
		clause, err := factory.CreateClause(clausesString)
		if err != nil {
			return nil
		}
		clauseString = strings.Trim(clauseString, (StartBracket + EndBracket))
		formulasString := strings.Split(clauseString, FormulaSeparator)
		for _, formulaString := range formulasString {
			formula, err := factory.CreateFormula(formulasString)
			if err != nil {
				return err
			}
			formulaString = strings.TrimSpace(formulaString)
			formulaString = strings.TrimLeft(formulaString, StartBracket)
			formulaString = strings.TrimRight(formulaString, EndBracket)
			err = formula.ParseString(factory, formulaString)
			if err != nil {
				return err
			}
			err = clause.AddFormula(formula)
			if err != nil {
				return err
			}
		}

		if len(clause.GetFormulas()) <= 0 {
			continue
		}

		err = rule.AddClause(clause)
		if err != nil {
			return err
		}
	}

	return nil
}

// String returns a string description of the instance
func (rule *Rule) String() string {
	if len(rule.Clauses) == 0 {
		return ""
	}

	if len(rule.Clauses) == 1 {
		return rule.Clauses[0].String()
	}

	var b bytes.Buffer
	b.Write([]byte(StartBracket))
	for n, clause := range rule.Clauses {
		if 0 < n {
			b.Write([]byte(SpaceSeparator + ClauseSeparator + SpaceSeparator))
		}
		b.Write([]byte(clause.String()))
	}
	b.Write([]byte(EndBracket))
	return b.String()
}
