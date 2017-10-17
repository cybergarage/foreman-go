// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

// VarI represents a interface for any variable.
type VarI interface {
	isTrue() bool
}

// A Var represents a variable in a SAT problem.
type Var struct {
	Name string
	VarI
}

// NewVar returns a new variable.
func NewVar() *Var {
	v := &Var{}
	return v
}
