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
	testFQLCaseFilename = "parser_test.csv"
)

func testFQLCase(t *testing.T, fqlString string) {
	parser := NewParser()

	err := parser.ParseString(fqlString)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%s\n", fqlString)
}

func TestFQLCases(t *testing.T) {

	fqlStrings, err := ioutil.ReadFile(testFQLCaseFilename)
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

		testFQLCase(t, fqlString)
	}
}
