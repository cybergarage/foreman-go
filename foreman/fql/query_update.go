// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// UpdateQuery represents a select query.
type UpdateQuery struct {
	*baseQuery
}

// NewUpdateQuery returns a new select query.
func NewUpdateQuery() Query {
	q := &UpdateQuery{
		baseQuery: newBaseQueryWithType(QueryTypeUpdate),
	}
	return q
}
