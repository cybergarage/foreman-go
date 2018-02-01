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
	testFQLSelectCaseFilename = "parser_test_select_case.csv"
)

type parserTestListener interface {
	testCase(t *testing.T, q Query, corrects []string) error
}

type selectQueryTestListener struct{}

func (l *selectQueryTestListener) testCase(t *testing.T, q Query, corrects []string) error {
	sq, _ := q.(*SelectQuery)
	table, _ := sq.GetTable()
	fmt.Printf("talbe = %s\n", table)
	if table != corrects[0] {
		t.Error(fmt.Errorf("%s != %s", table, corrects[0]))
	}
	return nil
}

func testFQLCase(t *testing.T, l parserTestListener, fqlString string, corrects []string) {
	parser := NewParser()

	queries, err := parser.ParseString(fqlString)
	if err != nil {
		t.Error(err)
		return
	}

	l.testCase(t, queries[0], corrects)
}

func testCSVCases(t *testing.T, filename string, l parserTestListener) {

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
		testFQLCase(t, l, fqlString, corrects)
	}
}

func TestFQLCases(t *testing.T) {
	testCSVCases(t, testFQLSelectCaseFilename, &selectQueryTestListener{})
}
