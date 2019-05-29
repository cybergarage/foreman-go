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

// ParseString parses a specified rule string.
/*
func (rule *BaseRule) ParseString(factory Factory, ruleString string) error {
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
*/

// ParseString parses a specified rule string.
func (rule *BaseRule) ParseString(factory Factory, ruleString string) error {

	/*
		input := antlr.NewInputStream(ruleString)
		lexer := NewknowledgebaseLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := NewknowledgebaseParser(stream)
		//el := newANTLRParserErrorListener()
		//p.AddErrorListener(el)
		p.BuildParseTrees = true
		tree := p.Fql()
		pl := newANTLRParserListener()
		antlr.ParseTreeWalkerDefault.Walk(pl, tree)
		if !el.IsSuccess() {
			return nil, fmt.Errorf("%s (%s)", queryString, el.GetError().Error())
		}
	*/

	return nil
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
