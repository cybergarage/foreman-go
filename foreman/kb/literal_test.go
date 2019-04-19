// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package kb

import (
	"testing"
)

type TestLiteral struct {
	Literal
	Value float64
}

func (v *TestLiteral) GetValue() (interface{}, error) {
	return v.Value, nil
}

func newTestLiteral() *TestLiteral {
	return &TestLiteral{
		Value: 0.0,
	}
}

func TestNewLiteral(t *testing.T) {
	newTestLiteral()
}
