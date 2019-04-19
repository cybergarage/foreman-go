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
		"x == 10.0",
		"x != 1.0",
		"x > 5.0",
		"x >= 10.0",
		"x < 20.0",
		"x <= 10.0",
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

	xm, _ := qos.GetVariable("x")

	// All formulas are good when x == 10.0

	err := xm.SetValue(10.0)
	if err != nil {
		t.Error(err)
		return
	}

	for _, formula := range formulas {
		testQoSGoodFormula(t, qos, formula)
	}

	// x == 1

	xm.SetValue(1.0)
	if err != nil {
		t.Error(err)
		return
	}
	goodFormulas := []kb.Formula{
		formulas[4],
		formulas[5],
	}
	for _, formula := range goodFormulas {
		testQoSGoodFormula(t, qos, formula)
	}
	badFormulas := []kb.Formula{
		formulas[0],
		formulas[1],
		formulas[2],
		formulas[3],
	}
	for _, formula := range badFormulas {
		testQoSBadFormula(t, qos, formula)
	}

	// x == 100.0

	err = xm.SetValue(100.0)
	if err != nil {
		t.Error(err)
		return
	}
	goodFormulas = []kb.Formula{
		formulas[1],
		formulas[2],
		formulas[3],
	}
	for _, formula := range goodFormulas {
		testQoSGoodFormula(t, qos, formula)
	}
	badFormulas = []kb.Formula{
		formulas[0],
		formulas[4],
		formulas[5],
	}
	for _, formula := range badFormulas {
		testQoSBadFormula(t, qos, formula)
	}
}
