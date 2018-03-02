// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// ResultSet is an alias of Parameters
type ResultSet struct {
	Parameters
}

// NewResultSetWithParameters returns a new result set of the specified parameters.
func NewResultSetWithParameters(params Parameters) *ResultSet {
	rs := &ResultSet{
		Parameters: params,
	}
	return rs
}

// NewResultSet returns a new result set.
func NewResultSet() *ResultSet {
	return NewResultSetWithParameters(NewParameters())
}
