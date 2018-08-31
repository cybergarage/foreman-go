// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// ExecuteQuery represents a execute query.
type ExecuteQuery struct {
	*baseQuery
}

// NewExecuteQuery returns a new execute query.
func NewExecuteQuery() Query {
	q := &ExecuteQuery{
		baseQuery: newBaseQueryWithType(QueryTypeExecute),
	}
	return q
}
