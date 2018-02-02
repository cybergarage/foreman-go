// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// SetQuery represents a set query.
type SetQuery struct {
	*InsertQuery
}

// NewSetQuery returns a new set query.
func NewSetQuery() Query {
	sq, _ := NewInsertQuery().(*InsertQuery)
	q := &SetQuery{
		InsertQuery: sq,
	}
	return q
}
