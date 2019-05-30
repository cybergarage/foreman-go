// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"github.com/cybergarage/foreman-go/foreman/kb"
)

// Formula represents a formula.
type Formula struct {
	*kb.BaseFormula
}

// NewFormulaWithParams returns a new base formula with the specified parameters.
func NewFormulaWithParams(lop kb.Operand, op kb.Operator, rop kb.Operand) *Formula {
	formula := &Formula{
		BaseFormula: kb.NewFormulaWithParams(lop, op, rop),
	}
	return formula
}

// NewFormula returns a new base formula.
func NewFormula() *Formula {
	return NewFormulaWithParams(nil, nil, nil)
}
