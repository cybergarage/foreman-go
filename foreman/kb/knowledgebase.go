// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import (
	"fmt"
)

//KnowledgeBase includes all knowledge rules.
type KnowledgeBase struct {
	Rules     map[string]Rule
	Variables map[string]Variable
}

// NewKnowledgeBase returns a new knowledge base.
func NewKnowledgeBase() *KnowledgeBase {
	kb := &KnowledgeBase{
		Rules:     nil,
		Variables: nil,
	}
	kb.Clear()
	return kb
}

// Clear removes all rules and variables..
func (kb *KnowledgeBase) Clear() error {
	kb.Rules = make(map[string]Rule)
	kb.Variables = make(map[string]Variable)
	return nil
}

// SetRule adds a new rules.
func (kb *KnowledgeBase) SetRule(rule Rule) error {
	// Add all variables in the rule
	for _, clause := range rule.GetClauses() {
		for _, formula := range clause.GetFormulas() {
			variable := formula.GetVariable()
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

	kb.Rules[rule.GetName()] = rule

	return nil
}

// SetRules adds a new rules.
func (kb *KnowledgeBase) SetRules(rules []Rule) error {
	for _, rule := range rules {
		err := kb.SetRule(rule)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetRule returns a rule of the specified name.
func (kb *KnowledgeBase) GetRule(name string) (Rule, bool) {
	rule, ok := kb.Rules[name]
	return rule, ok
}

// GetRules returns all rules.
func (kb *KnowledgeBase) GetRules() []Rule {
	rules := make([]Rule, 0)
	for _, rule := range kb.Rules {
		rules = append(rules, rule)
	}
	return rules
}

// SetRuleString parses a specified rule string.
func (kb *KnowledgeBase) SetRuleString(factory Factory, ruleName string, ruleString string) error {
	rule, err := kb.ParseRuleString(factory, ruleString)
	if err != nil {
		return err
	}
	rule.SetName(ruleName)
	return kb.SetRule(rule)
}

// ParseRuleString parses a specified rule string.
func (kb *KnowledgeBase) ParseRuleString(factory Factory, ruleString string) (Rule, error) {
	parser := NewParser()
	rule, err := parser.ParseString(factory, ruleString)
	if err != nil {
		return nil, err
	}

	return rule, nil
}

// ParseFormulaString parses a specified formula string.
func (kb *KnowledgeBase) ParseFormulaString(factory Factory, formulaString string) (Formula, error) {
	formula, err := factory.CreateFormula(formulaString)
	if err != nil {
		return nil, err
	}
	err = formula.ParseString(factory, formulaString)
	if err != nil {
		return nil, err
	}
	return formula, nil
}
