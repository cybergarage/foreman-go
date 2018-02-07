// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// Column is a column of of FQL.
type Column struct {
	Value string
}

// NewColumnWithString returns a new Column with the specified string.
func NewColumnWithString(columnString string) *Column {
	col := &Column{
		Value: columnString,
	}
	return col
}

// String returns the value.
func (col *Column) String() string {
	return col.Value
}
