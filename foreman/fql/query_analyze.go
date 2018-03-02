// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// AnalyzeQuery represents an analyze query.
type AnalyzeQuery struct {
	*baseQuery
}

// AnalyzeQuery returns a new analyze query.
func NewAnalyzeQuery() Query {
	q := &AnalyzeQuery{
		baseQuery: newBaseQuery(),
	}
	return q
}

// GetType returns a stored type.
func (q *AnalyzeQuery) GetType() QueryType {
	return QueryTypeAnalyze
}
