// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// baseQuery represents a parameter.
type baseQuery struct {
	token string
	Parameters
}

// newQueryWithToken returns a new query with the specified string.
func newBaseQuery() *baseQuery {
	q := &baseQuery{
		Parameters: NewParameters(),
	}
	return q
}

// GetType returns a stored type.
func (q *baseQuery) GetType() QueryType {
	return QueryTypeUnknown
}

// GetParameters returns all parameters
func (q *baseQuery) GetParameters() Parameters {
	return q.Parameters
}
