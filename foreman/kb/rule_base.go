// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import (
	"bytes"
)

// BaseRule represents a rule.
type BaseRule struct {
	Rule
	Name    string
	Clauses []Clause
}

// NewRule returns a new null rule.
func NewRule() *BaseRule {
	return NewRuleWithName("")
}

// NewRuleWithName returns a new null rule.
func NewRuleWithName(name string) *BaseRule {
	p := &BaseRule{
		Name:    name,
		Clauses: make([]Clause, 0),
	}
	return p
}

// SetName set a name to the rule.
func (rule *BaseRule) SetName(name string) error {
	rule.Name = name
	return nil
}

// GetName returns a name of the rule.
func (rule *BaseRule) GetName() string {
	return rule.Name
}

// AddClause adds a new clause.
func (rule *BaseRule) AddClause(clause Clause) error {
	rule.Clauses = append(rule.Clauses, clause)
	return nil
}

// GetClauses return all clauses in the rule.
func (rule *BaseRule) GetClauses() []Clause {
	return rule.Clauses
}

// GetVariables returns all variables in the formula.
func (rule *BaseRule) GetVariables() []Variable {
	vars := []Variable{}
	for _, clause := range rule.GetClauses() {
		clauseVars := clause.GetVariables()
		if len(clauseVars) <= 0 {
			continue
		}
		vars = append(vars, clauseVars...)
	}
	return vars
}

// IsSatisfied returns whether a clause in the rule is satisfied.
func (rule *BaseRule) IsSatisfied() (bool, error) {
	var lastErr error
	for _, clause := range rule.Clauses {
		ok, err := clause.IsSatisfied()

		// Unknown error is considered to be satisfied
		// ex. the operand variable is not found
		if err != nil {
			lastErr = err
			continue
		}

		if ok {
			return true, lastErr
		}
	}
	return false, lastErr
}

// String returns a string description of the instance
func (rule *BaseRule) String() string {
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
