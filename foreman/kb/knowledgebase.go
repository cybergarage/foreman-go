// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import (
	"fmt"
)

// KnowledgeBaseListener represents a listener interface.
type KnowledgeBaseListener interface {
	RuleListener
}

//KnowledgeBase includes all knowledge rules.
type KnowledgeBase struct {
	Rules        map[string]Rule
	Variables    map[string]*Variable
	relatedRules map[string][]Rule // relatedRules stores rules which are related the variable name
	listeners    []KnowledgeBaseListener
}

// NewKnowledgeBase returns a new knowledge base.
func NewKnowledgeBase() *KnowledgeBase {
	kb := &KnowledgeBase{
		Rules:        nil,
		Variables:    nil,
		relatedRules: nil,
		listeners:    make([]KnowledgeBaseListener, 0),
	}
	kb.Clear()
	return kb
}

// Clear removes all rules and variables..
func (kb *KnowledgeBase) Clear() error {
	kb.Rules = make(map[string]Rule)
	kb.Variables = make(map[string]*Variable)
	kb.relatedRules = make(map[string][]Rule)
	return nil
}

// notifyUpdatedRuleEvent check the specified rule and sent the result to the listeners.
func (kb *KnowledgeBase) notifyUpdatedRuleEvent(rule Rule) error {
	ok, err := rule.IsSatisfied()
	if err != nil {
		return err
	}

	for _, l := range kb.GetListeners() {
		if ok {
			l.RuleSatisfied(rule)
		} else {
			l.RuleUnsatisfied(rule)
		}
	}

	return nil
}

// UpdateVariableValue sets a new value to the specified variable.
func (kb *KnowledgeBase) UpdateVariableValue(name string, value interface{}) bool {
	v, ok := kb.Variables[name]
	if !ok {
		return false
	}

	err := v.SetValue(value)
	if err != nil {
		return false
	}

	// Exec rules and send the result to the listeners

	relatedRules, ok := kb.GetRelatedRules(name)
	if !ok {
		return false
	}

	for _, relatedRule := range relatedRules {
		err := kb.notifyUpdatedRuleEvent(relatedRule)
		if err != nil {
			return false
		}
	}

	return true
}

// GetVariable returns a variable of the specified name.
func (kb *KnowledgeBase) GetVariable(name string) (*Variable, bool) {
	v, ok := kb.Variables[name]
	return v, ok
}

// GetRelatedRules returns rules which are related the specified variable name.
func (kb *KnowledgeBase) GetRelatedRules(name string) ([]Rule, bool) {
	rules, ok := kb.relatedRules[name]
	return rules, ok
}

// addRelatedRules adds a new rule to the specified variable name.
func (kb *KnowledgeBase) addRelatedRules(name string, newRule Rule) error {
	rules, ok := kb.relatedRules[name]
	if !ok {
		rules = make([]Rule, 0)
	}

	for _, addedRule := range rules {
		// Is new rule added already ?
		if addedRule.GetName() == newRule.GetName() {
			return nil
		}
	}

	rules = append(rules, newRule)
	kb.relatedRules[name] = rules

	return nil
}

// removeRelatedRules removes the specified rule from the related rules.
func (kb *KnowledgeBase) removeRelatedRules(ruleName string) error {
	for varName, addedRules := range kb.relatedRules {
		for n, addedRule := range addedRules {
			if addedRule.GetName() == ruleName {
				newRules := append(addedRules[:n], addedRules[n+1:]...)
				kb.relatedRules[varName] = newRules
				break
			}
		}
	}
	return nil
}

// SetRule adds a new rules.
func (kb *KnowledgeBase) SetRule(rule Rule) error {
	// Add all variables in the rule
	for _, clause := range rule.GetClauses() {
		for _, formula := range clause.GetFormulas() {
			for _, operand := range formula.GetOperands() {
				variable, ok := operand.(*Variable)
				if !ok {
					continue
				}
				variableName := variable.GetName()

				mapVariable, ok := kb.GetVariable(variableName)
				if !ok {
					kb.Variables[variableName] = variable
					continue
				}

				err := kb.addRelatedRules(variableName, rule)
				if err != nil {
					return err
				}

				// Check whether these are a same instance
				if variable != mapVariable {
					return fmt.Errorf(errorInvalidRuleVariable, variableName, variable, mapVariable)
				}

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

// RemoveRule deletes a rule of the specified name.
func (kb *KnowledgeBase) RemoveRule(name string) bool {
	_, ok := kb.Rules[name]
	if !ok {
		return false
	}
	err := kb.removeRelatedRules(name)
	if err != nil {
		return false
	}
	delete(kb.Rules, name)
	return true
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

// HasListener checks whether the specified listener is already added
func (kb *KnowledgeBase) HasListener(listener KnowledgeBaseListener) bool {
	for _, l := range kb.listeners {
		if l == listener {
			return true
		}
	}
	return false
}

// AddListener adds a new listener
func (kb *KnowledgeBase) AddListener(listener KnowledgeBaseListener) bool {
	if kb.HasListener(listener) {
		return false
	}
	kb.listeners = append(kb.listeners, listener)
	return true
}

// RemoveListener removes the specified listener
func (kb *KnowledgeBase) RemoveListener(listener KnowledgeBaseListener) bool {
	for n, l := range kb.listeners {
		if l == listener {
			lastIndex := len(kb.listeners) - 1
			kb.listeners[lastIndex], kb.listeners[n] = kb.listeners[n], kb.listeners[lastIndex]
			kb.listeners = kb.listeners[:lastIndex]
			return true
		}
	}
	return false
}

// GetListeners returns the current listeners
func (kb *KnowledgeBase) GetListeners() []KnowledgeBaseListener {
	return kb.listeners
}
