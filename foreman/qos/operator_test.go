// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package qos

import (
	"testing"
)

type TestOperator struct {
}

func (v *TestOperator) IsSatisfied(value1 interface{}, value2 interface{}) (bool, error) {
	return true, nil
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
