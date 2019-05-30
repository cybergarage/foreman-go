// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"github.com/cybergarage/foreman-go/foreman/kb"
)

// QoS includes all QoS rules.
type QoS struct {
	kb.KnowledgeBase
}

// NewQoS returns a new null object.
func NewQoS() *QoS {
	qos := &QoS{
		KnowledgeBase: *kb.NewKnowledgeBase(),
	}
	return qos
}

// ParseString parses a specified QoS string.
func (qos *QoS) ParseString(qosString string) (kb.Rule, error) {
	return qos.ParseRuleString(qos, qosString)
}

// FindRelatedRules returns all rules of the the specified name.
func (qos *QoS) FindRelatedRules(q *Query) ([]*Rule, error) {
	name := q.Target
	qosRules := make([]*Rule, 0)

	for _, rule := range qos.Rules {
		for _, clause := range rule.GetClauses() {
			for _, formula := range clause.GetFormulas() {
				for _, operand := range formula.GetOperands() {
					v, ok := operand.(kb.Variable)
					if !ok {
						continue
					}
					if v.GetName() != name {
						continue
					}
					qosRule, ok := rule.(*Rule)
					if !ok {
						continue
					}
					qosRules = append(qosRules, qosRule)

				}
			}
		}
	}

	return qosRules, nil
}

// FindRelatedFormulas returns all QoS metrics of the the specified name.
func (qos *QoS) FindRelatedFormulas(q *Query) ([]*Formula, error) {
	name := q.Target
	qoSFormulas := make([]*Formula, 0)

	for _, rule := range qos.Rules {
		for _, clause := range rule.GetClauses() {
			for _, formula := range clause.GetFormulas() {
				for _, operand := range formula.GetOperands() {
					v, ok := operand.(kb.Variable)
					if !ok {
						continue
					}
					if v.GetName() != name {
						continue
					}
					qosFormula, ok := formula.(*Formula)
					if !ok {
						continue
					}
					qoSFormulas = append(qoSFormulas, qosFormula)
				}
			}
		}
	}

	return qoSFormulas, nil
}
