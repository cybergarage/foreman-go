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
func (qos *QoS) parseFormulaString(qosString string) (*kb.Formula, error) {
	return qos.ParseFormulaString(qos, qosString)
}

// CreateVariable is an interface method of kb.Factory
func (qos *QoS) CreateVariable(variable interface{}) (kb.Variable, error) {
	varStr, ok := variable.(string)
	if !ok {
		return nil, fmt.Errorf(errorInvalidVariable, variable)
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
func (qos *QoS) CreateObjective(objective interface{}) (kb.Objective, error) {
	objValue, ok := objective.(float64)
	if !ok {
		objStr, ok := objective.(string)
		if !ok {
			return nil, fmt.Errorf(errorInvalidObjective, objective)
		}

		var err error
		objValue, err = strconv.ParseFloat(objStr, 64)
		if err != nil {
			return nil, fmt.Errorf(errorInvalidObjective, objective)
		}
	}

	return NewThresholdWithValue(objValue), nil
}

// CreateOperator is an interface method kb.Factory
func (qos *QoS) CreateOperator(operator interface{}) (kb.Operator, error) {
	operatorStr, ok := operator.(string)
	if !ok {
		return nil, fmt.Errorf(errorInvalidOperator, operator)
	}
	ope, err := NewOperatorWithString(operatorStr)
	if err != nil {
		return nil, err
	}
	return ope, nil
}
