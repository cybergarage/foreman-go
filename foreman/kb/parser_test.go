// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package kb

import (
	"testing"
)

type TestParser struct {
	Factory
	variables map[string]Variable
}

func (parser *TestParser) CreateVariable(variable interface{}) (Variable, error) {
	varStr := variable.(string)
	v, ok := parser.variables[varStr]
	if ok {
		return v, nil
	}
	parser.variables[varStr] = newTestVariableWithName(varStr)
	return parser.variables[varStr], nil
}

func (parser *TestParser) CreateObjective(object interface{}) (Objective, error) {
	objStr := object.(string)
	return newTestObjectiveWithString(objStr), nil
}

func (parser *TestParser) CreateOperator(operator interface{}) (Operator, error) {
	opeStr := operator.(string)
	return newTestOperatorWithString(opeStr), nil
}

func newTestParser() *Parser {
	parser := &TestParser{
		variables: make(map[string]Variable),
	}
	return NewParserWithFactory(parser)
}

func TestNewParser(t *testing.T) {
	newTestParser()
}
