// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"testing"

	"github.com/cybergarage/foreman-go/foreman/kb"
)

const (
	errorInvalidFormulaOperator = "Invalid formula : %s"
)

func TestQoSLiteralOperands(t *testing.T) {
	qos := NewQoS()

	operands := []string{
		"1",
		"+1",
		"-1",
		"1.0",
		"+1.0",
		"-1.0",
	}

	for _, operand := range operands {
		op, err := qos.CreateOperand(operand)
		if err != nil {
			t.Error(err)
			break
		}
		_, ok := op.(*Threshold)
		if !ok {
			t.Errorf(operand)
			break
		}
	}
}

func TestQoSMetricOperands(t *testing.T) {
	qos := NewQoS()

	operands := []string{
		"m",
		"host.m",
		"org.cybergarage.host.m",
		"m01",
		"host01.m01",
		"org.cybergarage.host01.m01",
	}

	for _, operand := range operands {
		op, err := qos.CreateOperand(operand)
		if err != nil {
			t.Error(err)
			break
		}
		_, ok := op.(*Metric)
		if !ok {
			t.Errorf(operand)
			break
		}
	}
}

func TestQoSOperators(t *testing.T) {
	qos := NewQoS()

	operators := []string{
		"==",
		"!=",
		">",
		">=",
		"<",
		"<=",
	}

	for _, operator := range operators {
		_, err := qos.CreateOperator(operator)
		if err != nil {
			t.Error(err)
			break
		}
	}
}

func testQoSGoodFormula(t *testing.T, qos *QoS, formula kb.Formula) {
	ok, err := formula.IsSatisfied()
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Errorf(errorInvalidFormulaOperator, formula)
	}
}

func testQoSBadFormula(t *testing.T, qos *QoS, formula kb.Formula) {
	ok, err := formula.IsSatisfied()
	if err != nil {
		t.Error(err)
		return
	}
	if ok {
		t.Errorf(errorInvalidFormulaOperator, formula)
	}
}

func TestQoSFormulas(t *testing.T) {
	qos := NewQoS()

	formulaStrings := []string{
		"x == 10.0", // 0
		"x != 1.0",  // 1
		"x > 5.0",   // 2
		"x >= 10.0", // 3
		"x < 20.0",  // 4
		"x <= 10.0", // 5
		"x == y",    // 6
		"x != y",    // 7
		"x >= y",    // 8
		"x > y",     // 9
		"x <= y",    // 10
		"x < y",     // 11
	}

	formulas := make([]kb.Formula, len(formulaStrings))
	for n, formulaString := range formulaStrings {
		formula, err := qos.parseFormulaString(formulaString)
		if err != nil {
			t.Error(err)
		}
		formulas[n] = formula
	}

	// Get a singleton instance

	mx, _ := qos.GetVariable("x")
	my, _ := qos.GetVariable("y")

	// x == 10.0, y == 10.0

	mx.SetValue(10.0)
	my.SetValue(10.0)

	goodFormulas := []kb.Formula{
		formulas[0],
		formulas[1],
		formulas[2],
		formulas[4],
		formulas[5],
		formulas[6],
		formulas[8],
		formulas[10],
	}
	for _, formula := range goodFormulas {
		testQoSGoodFormula(t, qos, formula)
	}
	badFormulas := []kb.Formula{
		formulas[7],
		formulas[9],
		formulas[11],
	}
	for _, formula := range badFormulas {
		testQoSBadFormula(t, qos, formula)
	}

	// x == 1.0, y == 10.0

	mx.SetValue(1.0)

	goodFormulas = []kb.Formula{
		formulas[4],
		formulas[5],
		formulas[7],
		formulas[10],
		formulas[11],
	}
	for _, formula := range goodFormulas {
		testQoSGoodFormula(t, qos, formula)
	}
	badFormulas = []kb.Formula{
		formulas[0],
		formulas[1],
		formulas[2],
		formulas[3],
		formulas[8],
		formulas[9],
	}
	for _, formula := range badFormulas {
		testQoSBadFormula(t, qos, formula)
	}

	// x == 100.0, y == 10.0

	mx.SetValue(100.0)

	goodFormulas = []kb.Formula{
		formulas[1],
		formulas[2],
		formulas[3],
		formulas[7],
		formulas[8],
		formulas[9],
	}
	for _, formula := range goodFormulas {
		testQoSGoodFormula(t, qos, formula)
	}
	badFormulas = []kb.Formula{
		formulas[0],
		formulas[4],
		formulas[5],
		formulas[6],
		formulas[10],
		formulas[11],
	}
	for _, formula := range badFormulas {
		testQoSBadFormula(t, qos, formula)
	}

	// x == 100.0, y == 1000.0

	my.SetValue(1000.0)

	goodFormulas = []kb.Formula{
		formulas[1],
		formulas[2],
		formulas[3],
		formulas[7],
		formulas[10],
		formulas[11],
	}
	for _, formula := range goodFormulas {
		testQoSGoodFormula(t, qos, formula)
	}
	badFormulas = []kb.Formula{
		formulas[0],
		formulas[4],
		formulas[5],
		formulas[6],
		formulas[8],
		formulas[9],
	}
	for _, formula := range badFormulas {
		testQoSBadFormula(t, qos, formula)
	}
}
