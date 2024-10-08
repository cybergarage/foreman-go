// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

const (
	testFQLInsertCaseFilename  = "parser_test_insert_case.csv"
	testFQLSetCaseFilename     = "parser_test_set_case.csv"
	testFQLSelectCaseFilename  = "parser_test_select_case.csv"
	testFQLUpdateCaseFilename  = "parser_test_update_case.csv"
	testFQLExportCaseFilename  = "parser_test_export_case.csv"
	testFQLDeleteCaseFilename  = "parser_test_delete_case.csv"
	testFQLExecuteCaseFilename = "parser_test_execute_case.csv"
	testFQLAnalyzeCaseFilename = "parser_test_analyze_case.csv"
)

const (
	errorInvalidTarget         = "Invalid target %s != %s"
	errorInvalidValue          = "Invalid value %s != %s"
	errorInvalidColumnCount    = "Invalid column count %d != %d"
	errorInvalidColumn         = "Invalid column %s != %s"
	errorInvalidValueCount     = "Invalid value count %d != %d"
	errorInvalidConditionCount = "Invalid condition count %d != %d"
	errorInvalidCorrectCount   = "Invalid correct count %d != %d"
	errorInvalidOperatorType   = "Invalid operator type : %d"
)

type parserTestListener interface {
	testCase(t *testing.T, q Query, corrects []string) error
}

func testQuery(t *testing.T, l parserTestListener, fqlString string, corrects []string) {
	parser := NewParser()

	queries, err := parser.ParseString(fqlString)
	if err != nil {
		t.Error(err)
		return
	}

	err = l.testCase(t, queries[0], corrects)
	if err != nil {
		t.Errorf("%s (%s)", fqlString, err.Error())
		return
	}
}

func testCSVQueries(t *testing.T, filename string, l parserTestListener) {

	fqlStrings, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
		return
	}

	r := csv.NewReader(strings.NewReader(string(fqlStrings)))
	r.Comma = ';'
	r.Comment = '#'
	r.LazyQuotes = true
	r.FieldsPerRecord = -1

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err)
			break
		}

		fqlString := record[0]
		corrects := record[1:]
		testQuery(t, l, fqlString, corrects)
	}
}

func TestFQLCases(t *testing.T) {
	testCases := map[string]parserTestListener{
		testFQLInsertCaseFilename:  &insertQueryTestListener{},
		testFQLSetCaseFilename:     &insertQueryTestListener{},
		testFQLSelectCaseFilename:  &selectQueryTestListener{},
		testFQLExportCaseFilename:  &selectQueryTestListener{},
		testFQLUpdateCaseFilename:  &updateQueryTestListener{},
		testFQLDeleteCaseFilename:  &deleteQueryTestListener{},
		testFQLExecuteCaseFilename: &executeQueryTestListener{},
		testFQLAnalyzeCaseFilename: &analyzeQueryTestListener{},
	}

	for filename, listener := range testCases {
		testCSVQueries(t, filename, listener)
	}

}

// Insert

type insertQueryTestListener struct{}

func (l *insertQueryTestListener) testCase(t *testing.T, q Query, corrects []string) error {
	// Target

	target, _ := q.GetTarget()
	if target.String() != corrects[0] {
		return fmt.Errorf(errorInvalidTarget, target.String(), corrects[0])
	}

	// Values

	values, _ := q.GetValues()
	for n := 0; n < (len(corrects) - 1); n++ {
		if values[n].String() != corrects[n+1] {
			return fmt.Errorf(errorInvalidValue, values[n].String(), corrects[n+1])
		}
	}

	return nil
}

// Select

type selectQueryTestListener struct{}

func (l *selectQueryTestListener) testCase(t *testing.T, q Query, corrects []string) error {
	if len(corrects) < 3 {
		return fmt.Errorf(errorInvalidCorrectCount, len(corrects), 3)
	}

	// Target

	target, _ := q.GetTarget()
	if target.String() != corrects[0] {
		return fmt.Errorf(errorInvalidTarget, target.String(), corrects[0])
	}

	// Columns

	columnsCnt, err := strconv.Atoi(corrects[1])
	if err != nil {
		return err
	}
	columns, _ := q.GetColumns()
	if len(columns) != columnsCnt {
		return fmt.Errorf(errorInvalidColumnCount, len(columns), columnsCnt)
	}

	// Conditions

	condCnt, err := strconv.Atoi(corrects[2])
	if err != nil {
		return err
	}
	conds, _ := q.GetConditions()
	if len(conds) != condCnt {
		return fmt.Errorf(errorInvalidConditionCount, len(conds), condCnt)
	}
	for _, cond := range conds {
		opeType := cond.GetOperator().GetType()
		if opeType == OperatorTypeUnknown {
			return fmt.Errorf(errorInvalidOperatorType, opeType)

		}
	}

	return nil
}

// Update

type updateQueryTestListener struct{}

func (l *updateQueryTestListener) testCase(t *testing.T, q Query, corrects []string) error {
	// Target

	target, _ := q.GetTarget()
	if target.String() != corrects[0] {
		return fmt.Errorf(errorInvalidTarget, target.String(), corrects[0])
	}

	// Columns

	columnsCnt, err := strconv.Atoi(corrects[1])
	if err != nil {
		return err
	}
	columns, _ := q.GetColumns()
	if len(columns) != columnsCnt {
		return fmt.Errorf(errorInvalidColumnCount, len(columns), columnsCnt)
	}

	// Columns

	valueCnt, err := strconv.Atoi(corrects[1])
	if err != nil {
		return err
	}
	values, _ := q.GetValues()
	if len(values) != valueCnt {
		return fmt.Errorf(errorInvalidValueCount, len(values), valueCnt)
	}

	// Conditions

	condCnt, err := strconv.Atoi(corrects[2])
	if err != nil {
		return err
	}
	conds, _ := q.GetConditions()
	if len(conds) != condCnt {
		return fmt.Errorf(errorInvalidConditionCount, len(conds), condCnt)
	}

	return nil
}

// Delete

type deleteQueryTestListener struct{}

func (l *deleteQueryTestListener) testCase(t *testing.T, q Query, corrects []string) error {
	if len(corrects) < 2 {
		return fmt.Errorf(errorInvalidCorrectCount, len(corrects), 2)
	}

	// Target

	target, _ := q.GetTarget()
	if target.String() != corrects[0] {
		return fmt.Errorf(errorInvalidTarget, target.String(), corrects[0])
	}

	// Conditions

	condCnt, err := strconv.Atoi(corrects[1])
	if err != nil {
		return err
	}
	conds, _ := q.GetConditions()
	if len(conds) != condCnt {
		return fmt.Errorf(errorInvalidConditionCount, len(conds), condCnt)
	}
	for _, cond := range conds {
		opeType := cond.GetOperator().GetType()
		if opeType == OperatorTypeUnknown {
			return fmt.Errorf(errorInvalidOperatorType, opeType)

		}
	}

	return nil
}

// Execute

type executeQueryTestListener struct{}

func (l *executeQueryTestListener) testCase(t *testing.T, q Query, corrects []string) error {
	if len(corrects) < 1 {
		return fmt.Errorf(errorInvalidCorrectCount, len(corrects), 1)
	}

	// Target

	target, _ := q.GetTarget()
	if target.String() != corrects[0] {
		return fmt.Errorf(errorInvalidTarget, target.String(), corrects[0])
	}

	// len(columns) == len(values)

	columns, _ := q.GetColumns()
	values, _ := q.GetValues()

	if len(columns) != len(values) {
		return fmt.Errorf(errorInvalidValueCount, len(columns), len(values))
	}

	return nil
}

// Analyze

type analyzeQueryTestListener struct{}

func (l *analyzeQueryTestListener) testCase(t *testing.T, q Query, corrects []string) error {
	if len(corrects) < 2 {
		return fmt.Errorf(errorInvalidCorrectCount, len(corrects), 2)
	}

	// Target (corrects[0])

	target, _ := q.GetTarget()
	if target.String() != corrects[0] {
		return fmt.Errorf(errorInvalidTarget, target.String(), corrects[0])
	}

	// Columns (corrects[1])

	columns, _ := q.GetColumns()
	if len(columns) != 1 {
		return fmt.Errorf(errorInvalidColumnCount, len(columns), 1)
	}

	colums := columns[0]
	if colums.String() != corrects[1] {
		return fmt.Errorf(errorInvalidValue, colums.String(), columns[0])
	}

	return nil
}
