// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package kb

import (
	"testing"
)

type TestOperator struct {
}

func (v *TestOperator) IsSatisfied(value1 Operand, value2 Operand) (bool, error) {
	return true, nil
}

func (v *TestOperator) Expression() string {
	return ""
}

func newTestOperatorWithString(opeStr string) *TestOperator {
	return &TestOperator{}
}

func newTestOperator() *TestOperator {
	return newTestOperatorWithString("")
}

func TestNewOperator(t *testing.T) {
	newTestOperator()
}
