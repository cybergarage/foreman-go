// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// SelectQuery represents a select query.
type SelectQuery struct {
	*baseQuery
}

// NewSelectQuery returns a new select query.
func NewSelectQuery() Query {
	q := &SelectQuery{
		baseQuery: newBaseQueryWithType(QueryTypeSelect),
	}
	return q
}

// NewSelectAllQuery returns a new select query with an asterisk column.
func NewSelectAllQuery() Query {
	q := NewSelectQuery()
	q.AddColumn(NewColumnWithString(QueryColumnAll))
	return q
}
