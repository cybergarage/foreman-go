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
		baseQuery: newBaseQueryWithType(QueryTypeAnalyze),
	}
	return q
}
