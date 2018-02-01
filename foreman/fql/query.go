// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

type QueryType int

const (
	QueryTypeUnknown QueryType = iota
	QueryTypeSelect
)

// Query represents a query interface.
type Query interface {
	GetType() QueryType
	AddParameter(param Parameter) error
	GetParameters() Parameters
	GetParameter(string) (Parameter, bool)
	GetParameterString(string) (string, bool)
}
