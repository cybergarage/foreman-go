// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import (
	"fmt"
)

// Literal represents a simple operand object which has only a value without name.
type Literal struct {
	Operand
	Value interface{}
}

// NewLiteral returns a new Literal instance.
func NewLiteral() *Literal {
	return NewLiteralWithValue(nil)
}

// NewLiteralWithValue returns a new Literal instance with the specified value.
func NewLiteralWithValue(value interface{}) *Literal {
	th := &Literal{
		Value: value,
	}
	return th
}

// SetValue sets a specified value.
func (l *Literal) SetValue(value interface{}) error {
	l.Value = value
	return nil
}

// GetValue returns the stored value.
func (l *Literal) GetValue() (interface{}, error) {
	return l.Value, nil
}

// Expression returns a string for the formula expression.
func (l *Literal) Expression() string {
	return fmt.Sprintf("%v", l.Value)
}
