// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"fmt"
	"strconv"

	"github.com/cybergarage/foreman-go/foreman/kb"
)

// CreateClause is an interface method of kb.Factory
func (qos *QoS) CreateClause(obj interface{}) (kb.Clause, error) {
	return kb.NewClause(), nil
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
