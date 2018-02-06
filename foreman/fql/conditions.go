// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// Conditions represents a condition array.
type Conditions []*Condition

// NewConditions returns a new parameter map.
func NewConditions() Conditions {
	c := make(Conditions, 0)
	return c
}
