// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package kb

import (
	"fmt"
	"testing"
)

type TestOperand struct {
	Value float64
}

func (obj *TestOperand) GetValue() interface{} {
	return obj.Value
}

func (obj *TestOperand) String() string {
	return fmt.Sprintf("%f", obj.Value)
}

func newTestOperandWithString(value string) *TestOperand {
	return &TestOperand{
		Value: 0.0,
	}
}

func newTestOperand() *TestOperand {
	return newTestOperandWithString("")
}

func TestNewOperand(t *testing.T) {
	newTestOperand()
}
