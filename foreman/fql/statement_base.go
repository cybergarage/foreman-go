// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// baseParameter represents a parameter.
type baseParameter struct {
	Name       string
	Type       StatementType
	Parameters Parameters
}

// NewStatementWithName returns a new statement.
func NewStatementWithName(name string, stype StatementType) Statement {
	stmt := &baseParameter{
		Name: name,
		Type: stype,
	}
	return stmt
}

// GetName returns a stored name.
func (stmt *baseParameter) GetName() string {
	return stmt.Name
}

// GetName returns a stored type.
func (stmt *baseParameter) GetType() StatementType {
	return stmt.Type
}

func (stmt *baseParameter) GetParameters() Parameters {
	return stmt.Parameters
}

func (stmt *baseParameter) GetParameter(name string) (Parameter, bool) {
	param, ok := stmt.Parameters[name]
	return param, ok
}

// String returns a string description of the instance
func (stmt *baseParameter) String() string {
	return stmt.GetName()
}
