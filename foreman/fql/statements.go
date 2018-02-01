// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// Parameters represents a parameter map.
type Statements []Statement

// NewStatements returns a new parameter map.
func NewStatements() Statements {
	stmts := make([]Statement, 0)
	return stmts
}

// AddStatement adds a new statement.
func (stmts Statements) AddStatement(stmt Statement) error {
	stmts = append(stmts, stmt)
	return nil
}
