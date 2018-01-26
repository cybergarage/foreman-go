// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// ParameterType defines parameter types
type ParameterType int

const (
	ParameterUnknownType ParameterType = iota
	ParameterIntegerType
	ParameterRealType
	ParameterBoolType
	ParameterStringType
)

// Parameter represents a interface for the action parameter.
type Parameter interface {
	GetName() string
	GetType() ParameterType
	GetValue() interface{}
	Equals(other Parameter) bool
	IsInteger() bool
	IsReal() bool
	IsBool() bool
	IsString() bool
	GetInteger() (int64, error)
	GetReal() (float64, error)
	GetBool() (bool, error)
	GetString() (string, error)
	String() string
}
