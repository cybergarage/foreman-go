// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/cybergarage/foreman-go/foreman/kb"
)

const (
	qosOperandThresholdRegex = "^[+-]?[0-9]+([\\.][0-9]+)?"
	qosOperandMetricRegex    = "^[A-Za-z][0-9A-Za-z_\\.]*"
)

var qosOperandRegexes []*regexp.Regexp

// getQoSOperandRegexes returns a slice of the QoS operand Regexp formats.
func getQoSOperandRegexes() []*regexp.Regexp {
	if qosOperandRegexes == nil {
		qosOperandRegexes = []*regexp.Regexp{
			regexp.MustCompile(qosOperandThresholdRegex),
			regexp.MustCompile(qosOperandMetricRegex),
		}
	}
	return qosOperandRegexes
}

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

// CreateOperand is an interface method of kb.Factory
func (qos *QoS) CreateOperand(obj interface{}) (kb.Operand, error) {
	switch obj.(type) {
	case string:
		s, _ := obj.(string)
		for n, qosRegexp := range getQoSOperandRegexes() {
			if !qosRegexp.MatchString(s) {
				continue
			}
			switch n {
			case 0: // qosOperandThresholdRegex
				return qos.createThreshold(obj)
			case 1: // qosOperandMetricRegex
				return qos.createMetric(obj)
			}
		}
	case float64:
		return qos.createThreshold(obj)
	}

	return nil, fmt.Errorf(errorInvalidOperand, obj)
}

// createMetric is an interface method of kb.Factory
func (qos *QoS) createMetric(obj interface{}) (kb.Operand, error) {
	varStr, ok := obj.(string)
	if !ok {
		return nil, fmt.Errorf(errorInvalidOperand, obj)
	}

	v, ok := qos.Variables[varStr]
	if ok {
		return v, nil
	}

	m := NewMetricWithName(varStr)
	qos.Variables[varStr] = m

	return m, nil
}

// createThreshold is an interface method kb.Factory
func (qos *QoS) createThreshold(obj interface{}) (kb.Operand, error) {
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
