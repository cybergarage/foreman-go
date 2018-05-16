// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// DeleteQuery represents a delete query.
type DeleteQuery struct {
	*baseQuery
}

// NewDeleteQuery returns a new delete query.
func NewDeleteQuery() Query {
	q := &DeleteQuery{
		baseQuery: newBaseQuery(),
	}
	return q
}

// GetType returns a query type.
func (q *DeleteQuery) GetType() QueryType {
	return QueryTypeDelete
}

// IsStateChangeQuery returns whether state change query
func (q *DeleteQuery) IsStateChangeQuery() bool {
	return true
}
