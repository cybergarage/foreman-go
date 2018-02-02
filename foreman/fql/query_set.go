// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// SetQuery represents a set query.
type SetQuery struct {
	*baseQuery
}

// NewSetQuery returns a new set query.
func NewSetQuery() Query {
	q := &SetQuery{
		baseQuery: newBaseQuery(),
	}
	return q
}

// GetType returns a stored type.
func (q *SetQuery) GetType() QueryType {
	return QueryTypeSet
}
