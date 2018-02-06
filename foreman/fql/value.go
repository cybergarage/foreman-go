// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// Value is a value of of FQL.
type Value struct {
	Value string
}

// NewValueWithString returns a new value with the specified string.
func NewValueWithString(valString string) *Value {
	v := &Value{
		Value: valString,
	}
	return v
}

// GetValue returns the value.
func (v *Value) GetValue() string {
	return v.Value
}
