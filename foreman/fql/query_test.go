// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"testing"
)

func TestNewQuery(t *testing.T) {
	newBaseQuery()

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
