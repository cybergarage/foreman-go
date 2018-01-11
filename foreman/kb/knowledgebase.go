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

// NewKnowledgeBase returns a new knowledge base instance.
func NewKnowledgeBase() *KnowledgeBase {
	kb := &KnowledgeBase{
		Rules:     make([]*Rule, 0),
		Variables: make(map[string]Variable),
	}
	return kb
}

// AddRule adds a new rules.
func (kb *KnowledgeBase) AddRule(rule *Rule) error {
	// Add all variables in the rule
	for _, clause := range rule.Clauses {
		for _, quality := range clause.Qualities {
			variable := quality.Variable
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
