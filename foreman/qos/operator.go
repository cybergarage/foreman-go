// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

// A Operator represents an operator in a formula.
type Operator struct {
	Type int
}

// NewOperator returns a new operator with the specified type.
func NewOperator(t int) *Operator {
	op := &Operator{
		Type: t,
	}
	return op
}
