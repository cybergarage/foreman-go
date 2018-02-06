// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

type QueryType int

const (
	QueryTypeUnknown QueryType = iota
	QueryTypeInsert
	QueryTypeSelect
	QueryTypeDelete
)

// Query represents a query interface.
type Query interface {
	GetType() QueryType

	AddCondition(*Condition) error
	GetConditions() Conditions

	SetTarget(*Target) error
	GetTarget() (*Target, bool)

	AddValue(*Value) error
	GetValues() (Values, bool)
}
