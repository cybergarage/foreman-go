// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import (
	"fmt"
)

// BaseLiteral represents a simple operand object which has only a value without name.
type BaseLiteral struct {
	Operand
	Value interface{}
}

// NewLiteral returns a new BaseLiteral instance.
func NewLiteral() *BaseLiteral {
	return NewLiteralWithValue(nil)
}

// NewLiteralWithValue returns a new BaseLiteral instance with the specified value.
func NewLiteralWithValue(value interface{}) *BaseLiteral {
	th := &BaseLiteral{
		Value: value,
	}
	return th
}

// SetValue sets a specified value.
func (l *BaseLiteral) SetValue(value interface{}) error {
	l.Value = value
	return nil
}

// GetValue returns the stored value.
func (l *BaseLiteral) GetValue() (interface{}, error) {
	return l.Value, nil
}

// Expression returns a string for the formula expression.
func (l *BaseLiteral) Expression() string {
	return fmt.Sprintf("%v", l.Value)
}
