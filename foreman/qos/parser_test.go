// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package qos

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

const (
	testQoSCaseFilename = "parser_test.csv"
)

type TestParser struct {
	Factory
	variables map[string]Variable
}

func (parser *TestParser) CreateVariable(varStr string) (Variable, error) {
	v, ok := parser.variables[varStr]
	if ok {
		return v, nil
	}
	parser.variables[varStr] = newTestVariableWithName(varStr)
	return parser.variables[varStr], nil
}

func (parser *TestParser) CreateObjective(objStr string) (Objective, error) {
	return newTestObjectiveWithString(objStr), nil
}

func (parser *TestParser) CreateOperator(opeStr string) (Operator, error) {
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

func testQoSCase(t *testing.T, qosString string, variables int, clauses int) {
	parser := newTestParser()
	parser.ParseString(qosString)
}

func TestQoSCases(t *testing.T) {
	qosStrings, err := ioutil.ReadFile(testQoSCaseFilename)
	if err != nil {
		t.Error(err)
		return
	}

	r := csv.NewReader(strings.NewReader(string(qosStrings)))
	r.Comment = rune('#')

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err)
			break
		}

		ruleString := record[0]
		variables, err := strconv.Atoi(record[1])
		if err != nil {
			t.Error(err)
			break
		}
		clauses, err := strconv.Atoi(record[2])
		if err != nil {
			t.Error(err)
			break
		}

		testQoSCase(t, ruleString, variables, clauses)
	}
}
