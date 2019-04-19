// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// Operand represents an objective interface in formulas.
type Operand interface {
	// GetValue returns the operand value.
	GetValue() (interface{}, error)
	// Expression returns a string for the formula expression.
	Expression() string
}
