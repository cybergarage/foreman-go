// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// Value is a value of of FQL.
type Value = string

// NewValue returns a new value.
func NewValue() Value {
	var value Value
	return value
}
