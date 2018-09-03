// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"testing"
)

func TestNewQuery(t *testing.T) {
	NewInsertQuery()
	NewSetQuery()
	NewSelectQuery()
	NewExportQuery()
	NewDeleteQuery()
	NewAnalyzeQuery()
}

type testStatusChangeQuery struct {
	query          Query
	isStatusChange bool
}

func TestStateChangeQueries(t *testing.T) {
	testQueries := []*testStatusChangeQuery{
		&testStatusChangeQuery{NewInsertQuery(), true},
		&testStatusChangeQuery{NewSetQuery(), true},
		&testStatusChangeQuery{NewSelectQuery(), false},
		&testStatusChangeQuery{NewExportQuery(), false},
		&testStatusChangeQuery{NewDeleteQuery(), true},
		&testStatusChangeQuery{NewAnalyzeQuery(), false},
	}

	for n, testQuery := range testQueries {
		if testQuery.query.IsStateChangeQuery() != testQuery.isStatusChange {
			t.Errorf("[%d] %t != %t", n, testQuery.query.IsStateChangeQuery(), testQuery.isStatusChange)
		}
	}
}

func TestQueryToString(t *testing.T) {
	testQueries := []string{
		"INSERT INTO QOS VALUES (qos000)",
		"INSERT INTO QOS VALUES (qos000, \"m000 < 000\")",
	}

	for n, testQuery := range testQueries {
		parser := NewParser()
		queries, err := parser.ParseString(testQuery)
		if err != nil {
			t.Error(err)
			continue
		}
		query := queries[0]
		queryString := query.String()
		if queryString != testQuery {
			t.Errorf("[%d] %s != %s", n, queryString, testQuery)
		}
	}
}

func TestSelectAllQuery(t *testing.T) {
	testQuery := "SELECT * FROM QOS"

	for n := 0; n < 2; n++ {
		parser := NewParser()
		queries, err := parser.ParseString(testQuery)
		if err != nil {
			t.Error(err)
			return
		}

		query := queries[0]
		if !query.IsAllColumn() {
			t.Errorf("%s", testQuery)
		}

		testQuery = query.String()
	}
}
