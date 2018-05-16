// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// AnalyzeQuery represents a Analyze query.
type AnalyzeQuery struct {
	*baseQuery
}

// NewAnalyzeQuery returns a new Analyze query.
func NewAnalyzeQuery() Query {
	q := &AnalyzeQuery{
		baseQuery: newBaseQuery(),
	}
	return q
}

// GetType returns a query type.
func (q *AnalyzeQuery) GetType() QueryType {
	return QueryTypeAnalyze
}

// IsStateChangeQuery returns whether state change query
func (q *AnalyzeQuery) IsStateChangeQuery() bool {
	return false
}
