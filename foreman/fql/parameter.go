// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

const (
	parameterTable  = "table"
	parameterValues = "values"
)

// Parameter represents a parameter interface.
type Parameter interface {
	SetName(string) error
	SetValue(interface{}) error
	GetName() string
	GetValue() interface{}
	GetString() (string, bool)
}
