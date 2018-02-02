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
		baseQuery: newBaseQuery(),
	}
	return q
}

// GetType returns a stored type.
func (q *SelectQuery) GetType() QueryType {
	return QueryTypeSelect
}

// GetTarget returns the target name.
func (q *SelectQuery) GetTarget() (string, bool) {
	return q.GetParameterString(parameterTable)
}
