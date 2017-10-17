// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import
// A Literal represents a literal in a SAT problem.
type Literal struct {
	Var  *Var
	Sign bool
}

// NewLiteral returns a new literal.
func NewLiteral(v *Var, sign bool) *Literal {
	l := &Literal{v, sign}
	return l
}
