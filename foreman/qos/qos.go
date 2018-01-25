// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"fmt"
	"strconv"

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

// ParseQoSString parses a specified QoS string.
func (qos *QoS) ParseQoSString(qosString string) error {
	return qos.ParseRuleString(qos, qosString)
}

// parseFormulaString parses a specified QoS formula string.
func (qos *QoS) parseFormulaString(qosString string) (kb.Formula, error) {
	return qos.ParseFormulaString(qos, qosString)
}

// FindRelatedFormulas returns all QoS metrics of the the specified name.
func (qos *QoS) FindRelatedFormulas(q *Query) ([]*Formula, error) {
	name := q.Target
	qoSFormulas := make([]*Formula, 0)

	for _, rule := range qos.Rules {
		for _, clause := range rule.Clauses {
			for _, formula := range clause.Formulas {
				v := formula.GetVariable()
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

	return qoSFormulas, nil
}

// CreateFormula is an interface method of kb.Factory
func (qos *QoS) CreateFormula(obj interface{}) (kb.Formula, error) {
	return NewFormula(), nil
}

// CreateVariable is an interface method of kb.Factory
func (qos *QoS) CreateVariable(obj interface{}) (kb.Variable, error) {
	varStr, ok := obj.(string)
	if !ok {
		return nil, fmt.Errorf(errorInvalidVariable, obj)
	}

	v, ok := qos.Variables[varStr]
	if ok {
		return v, nil
	}

	m := NewMetricWithName(varStr)
	qos.Variables[varStr] = m

	return m, nil
}

// CreateObjective is an interface method kb.Factory
func (qos *QoS) CreateObjective(obj interface{}) (kb.Objective, error) {
	objValue, ok := obj.(float64)
	if !ok {
		objStr, ok := obj.(string)
		if !ok {
			return nil, fmt.Errorf(errorInvalidObjective, obj)
		}

		var err error
		objValue, err = strconv.ParseFloat(objStr, 64)
		if err != nil {
			return nil, fmt.Errorf(errorInvalidObjective, obj)
		}
	}

	return NewThresholdWithValue(objValue), nil
}

// CreateOperator is an interface method kb.Factory
func (qos *QoS) CreateOperator(obj interface{}) (kb.Operator, error) {
	operatorStr, ok := obj.(string)
	if !ok {
		return nil, fmt.Errorf(errorInvalidOperator, obj)
	}
	ope, err := NewOperatorWithString(operatorStr)
	if err != nil {
		return nil, err
	}
	return ope, nil
}
