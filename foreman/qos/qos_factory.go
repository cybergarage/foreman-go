// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"fmt"
	"strconv"

	"github.com/cybergarage/foreman-go/foreman/kb"
	"github.com/cybergarage/foreman-go/foreman/qos/function"
)

// CreateRule is an interface method of kb.Factory
func (qos *QoS) CreateRule() kb.Rule {
	return NewRule()
}

// CreateClause is an interface method of kb.Factory
func (qos *QoS) CreateClause() kb.Clause {
	return kb.NewClause()
}

// CreateFormula is an interface method of kb.Factory
func (qos *QoS) CreateFormula(lop kb.Operand, op kb.Operator, rop kb.Operand) kb.Formula {
	return NewFormulaWithParams(lop, op, rop)
}

// CreateLiteralOperand  is an interface method of kb.Factory
func (qos *QoS) CreateLiteralOperand(obj interface{}) (kb.Literal, error) {
	objValue, ok := obj.(float64)
	if !ok {
		objStr, ok := obj.(string)
		if !ok {
			return nil, fmt.Errorf(errorInvalidOperand, obj)
		}

		var err error
		objValue, err = strconv.ParseFloat(objStr, 64)
		if err != nil {
			return nil, fmt.Errorf(errorInvalidOperand, obj)
		}
	}

	return NewThresholdWithValue(objValue), nil
}

// CreateVariableOperand  is an interface method of kb.Factory
func (qos *QoS) CreateVariableOperand(name string) (kb.Variable, error) {
	v, ok := qos.Variables[name]
	if ok {
		return v, nil
	}

	m := NewMetricWithName(name)
	qos.Variables[name] = m

	return m, nil
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

// CreateFunctionOperand  is an interface method of kb.Factory
func (qos *QoS) CreateFunctionOperand(name string) (kb.Function, error) {
	var fn kb.Function

	switch name {
	case "ABS":
		fn = function.NewAbs()
	}

	if fn == nil {
		return nil, fmt.Errorf(errorUnknownFunction, name)
	}

	qosFn, ok := fn.(Function)
	if ok {
		qosFn.SetMetricManager(qos.GetMetricManager())
		qosFn.SetRegisterManager(qos.GetRegisterManager())
		qosFn.SetRegistryManager(qos.GetRegistryManager())
	}

	return fn, nil
}
