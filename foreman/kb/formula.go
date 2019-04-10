// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// A Formula represents a formula interface.
type Formula interface {
	GetVariable() Variable
	GetOperator() Operator
	GetObjective() Operand
	ParseString(Factory, string) error
	IsSatisfied() (bool, error)
	String() string
}
