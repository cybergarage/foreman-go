// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// ExportQuery is an alias for SelectQuery.
type ExportQuery struct {
	*SelectQuery
}

// NewExportQuery returns a new export query.
func NewExportQuery() Query {
	sq, _ := NewSelectQuery().(*SelectQuery)
	q := &ExportQuery{
		SelectQuery: sq,
	}
	return q
}
