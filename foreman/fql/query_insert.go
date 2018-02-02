// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// InsertQuery represents an insert query.
type InsertQuery struct {
	*baseQuery
}

// NewInsertQuery returns a new insert query.
func NewInsertQuery() Query {
	q := &InsertQuery{
		baseQuery: newBaseQuery(),
	}
	return q
}

// GetType returns a stored type.
func (q *InsertQuery) GetType() QueryType {
	return QueryTypeInsert
}
