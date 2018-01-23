// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import (
	"fmt"
)

//KnowledgeBase includes all knowledge rules.
type KnowledgeBase struct {
	Rules     []*Rule
	Variables map[string]Variable
}

// NewKnowledgeBase returns a new knowledge base.
func NewKnowledgeBase() *KnowledgeBase {
	kb := &KnowledgeBase{}
	kb.Clear()
	return kb
}

// Clear removes all rules and variables..
func (kb *KnowledgeBase) Clear() error {
	kb.Rules = make([]*Rule, 0)
	kb.Variables = make(map[string]Variable)
	return nil
}

// AddRule adds a new rules.
func (kb *KnowledgeBase) AddRule(rule *Rule) error {
	// Add all variables in the rule
	for _, clause := range rule.Clauses {
		for _, formula := range clause.Formulas {
			variable := formula.Variable
			variableName := variable.GetName()

			mapVariable, ok := kb.Variables[variableName]
			if !ok {
				kb.Variables[variableName] = variable
				continue
			}

			// Check whether these are a same instance
			if variable != mapVariable {
				return fmt.Errorf(errorInvalidRuleVariable, variableName, variable, mapVariable)
			}
		}
	}

	kb.Rules = append(kb.Rules, rule)

	return nil
}

// AddRules adds a new rules.
func (kb *KnowledgeBase) AddRules(rules []*Rule) error {
	for _, rule := range rules {
		err := kb.AddRule(rule)
		if err != nil {
			return err
		}
	}

	return nil
}

// ParseRuleString parses a specified rule string.
func (kb *KnowledgeBase) ParseRuleString(factory Factory, ruleString string) error {
	parser := NewParser()
	rule, err := parser.ParseString(factory, ruleString)
	if err != nil {
		return err
	}

	return kb.AddRule(rule)
}

// ParseFormulaString parses a specified formula string.
func (kb *KnowledgeBase) ParseFormulaString(factory Factory, formulaString string) (*Formula, error) {
	formula := NewFormula()
	err := formula.ParseString(factory, formulaString)
	if err != nil {
		return nil, err
	}
	return formula, nil
}
