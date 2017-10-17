// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package qos

import (
	"testing"
)

type TrueVar struct {
	*Var
}

func (v *TrueVar) isTrue() bool {
	return true
}

func NewTestVar() *TrueVar {
	return &TrueVar{
		Var: NewVar(),
	}
}
func TestNewVar(t *testing.T) {
	NewVar()
	//v.isTrue()
}
