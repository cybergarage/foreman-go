// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"testing"

	"github.com/cybergarage/foreman-go/foreman/kb"
	"github.com/cybergarage/foreman-go/foreman/metric"
)

const (
	errorInvalidFormulaOperator = "Invalid formula : %s"
)

func testQoSFactory(t *testing.T, factory kb.Factory) {
	factory.CreateVariable("var")
	factory.CreateOperator(">")
	factory.CreateObjective("th")
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

func testQoSOperators(t *testing.T, qos *QoS) {
	formulaStrings := []string{
		"x == 10.0",
		"x != 1.0",
		"x > 5.0",
		"x >= 10.0",
		"x < 20.0",
		"x <= 10.0",
	}

	xm := metric.NewMetricWithName("x")

	formulas := make([]kb.Formula, len(formulaStrings))
	for n, formulaString := range formulaStrings {
		formula, err := qos.parseFormulaString(formulaString)
		if err != nil {
			t.Error(err)
		}

		vm := formula.GetVariable()
		qm, _ := vm.(*Metric)
		qm.SetEntity(xm)

		formulas[n] = formula
	}

	// All formulas are good when x == 10.0

	xm.Value = 10.0
	for _, formula := range formulas {
		testQoSGoodFormula(t, qos, formula)
	}

	// x == 1

	xm.Value = 1.0
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

	xm.Value = 100.0
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

func TestNewQoS(t *testing.T) {
	qos := NewQoS()
	testQoSFactory(t, qos)
	testQoSOperators(t, qos)
}
