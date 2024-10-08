// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

type QueryType int

const (
	QueryTypeUnknown QueryType = iota
	QueryTypeInsert
	QueryTypeSelect
	QueryTypeUpdate
	QueryTypeDelete
	QueryTypeAnalyze
	QueryTypeExecute
)

// Query represents a query interface.
type Query interface {
	GetType() QueryType

	SetTarget(*Target) error
	GetTarget() (*Target, bool)

	AddColumn(*Column) error
	GetColumns() (Columns, bool)
	HasColumn(string) bool
	HasOnlyColumn(string) bool
	IsAllColumn() bool

	AddValue(*Value) error
	GetValues() (Values, bool)

	AddCondition(*Condition) error
	GetConditions() (Conditions, bool)
	GetConditionByColumn(leftOpe string) (*Operator, string, bool)

	IsStateChangeQuery() bool

	SetRetransmissionFlag(bool)
	IsRetransmissionQuery() bool

	String() string
}
