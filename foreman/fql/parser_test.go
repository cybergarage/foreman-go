// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

const (
	testFQLInsertCaseFilename = "parser_test_insert_case.csv"
	testFQLSetCaseFilename    = "parser_test_set_case.csv"
	testFQLSelectCaseFilename = "parser_test_select_case.csv"
	testFQLExportCaseFilename = "parser_test_export_case.csv"
)

const (
	errorInvalidTarget     = "Invalid target %s != %s"
	errorInvalidValue      = "Invalid value%s != %s"
	errorInvalidValueCount = "Invalid value count %d != %d"
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
		t.Error(err)
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

		fqlString := record[0]
		corrects := record[1:]
		testQuery(t, l, fqlString, corrects)
	}
}

func TestFQLCases(t *testing.T) {
	testCases := map[string]parserTestListener{
		testFQLSetCaseFilename:    &setQueryTestListener{},
		testFQLSelectCaseFilename: &selectQueryTestListener{},
		testFQLExportCaseFilename: &exportQueryTestListener{},
	}

	for filename, listener := range testCases {
		testCSVQueries(t, filename, listener)
	}

}

// Insert

type insertQueryTestListener struct{}

func (l *insertQueryTestListener) testCase(t *testing.T, q Query, corrects []string) error {
	return nil
}

// Set

type setQueryTestListener struct{}

func (l *setQueryTestListener) testCase(t *testing.T, q Query, corrects []string) error {
	sq, _ := q.(*SetQuery)

	values, _ := sq.GetValues()
	if len(values) != 1 {
		return fmt.Errorf(errorInvalidValueCount, len(values), 1)
	}
	if values[0] != corrects[0] {
		return fmt.Errorf(errorInvalidValue, values[0], corrects[0])
	}

	target, _ := sq.GetTarget()
	if target != corrects[1] {
		return fmt.Errorf(errorInvalidTarget, target, corrects[1])
	}

	return nil
}

// Select

type selectQueryTestListener struct{}

func (l *selectQueryTestListener) testCase(t *testing.T, q Query, corrects []string) error {
	sq, _ := q.(*SelectQuery)
	target, _ := sq.GetTarget()
	if target != corrects[0] {
		return fmt.Errorf(errorInvalidTarget, target, corrects[0])
	}
	return nil
}

// Export

type exportQueryTestListener struct{}

func (l *exportQueryTestListener) testCase(t *testing.T, q Query, corrects []string) error {
	sq, _ := q.(*ExportQuery)
	target, _ := sq.GetTarget()
	if target != corrects[0] {
		return fmt.Errorf(errorInvalidTarget, target, corrects[0])
	}
	return nil
}
