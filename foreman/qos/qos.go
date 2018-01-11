// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/kb"
)

// QoS includes all QoS rules.
type QoS struct {
	kb.KnowledgeBase
}

// NewQoS returns a new null object.
func NewQoS() *QoS {
	qos := &QoS{}
	qos.KnowledgeBase = *kb.NewKnowledgeBaseWithFactory(qos)
	return qos
}

// ParseQoSString parses a specified QoS string.
func (qos *QoS) ParseQoSString(qosString string) error {
	return qos.ParseRuleString(qosString)
}

// CreateVariable is a interface of kb.Factory
func (qos *QoS) CreateVariable(variable interface{}) (kb.Variable, error) {
	varStr, ok := variable.(string)
	if !ok {
		return nil, fmt.Errorf(errorInvalidVariable, variable)
	}

	v, ok := qos.Variables[varStr]
	if ok {
		return v, nil
	}

	m := NewMetric()
	m.Name = varStr
	qos.Variables[varStr] = m

	return m, nil
}

// CreateObjective is a interface of kb.Factory
func (qos *QoS) CreateObjective(objective interface{}) (kb.Objective, error) {
	objValue, ok := objective.(float64)
	if !ok {
		return nil, fmt.Errorf(errorInvalidObjective, objective)
	}
	return NewThresholdWithValue(objValue), nil
}

// CreateOperator is a interface of kb.Factory
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
