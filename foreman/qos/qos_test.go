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

	"github.com/cybergarage/foreman-go/foreman/kb"
)

const (
	testQoSCaseFilename = "parser_test.csv"
)

func testQoSFactory(t *testing.T, factory kb.Factory) {
	factory.CreateVariable("var")
	factory.CreateOperator(">")
	factory.CreateObjective("th")
}

func TestNewQoS(t *testing.T) {
	qos := NewQoS()
	testQoSFactory(t, qos)
}

func testQoSCase(t *testing.T, qos *QoS, qosString string, variables int, clauses int) {
	qos.ParseQoSString(qosString)
	qos.Clear()
}

func TestQoSCases(t *testing.T) {
	qos := NewQoS()

	kbStrings, err := ioutil.ReadFile(testQoSCaseFilename)
	if err != nil {
		t.Error(err)
		return
	}

	r := csv.NewReader(strings.NewReader(string(kbStrings)))
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

		qosString := record[0]
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

		testQoSCase(t, qos, qosString, variables, clauses)
	}
}
