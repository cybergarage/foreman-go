// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package kb

import (
	"fmt"
	"testing"
)

type TestObjective struct {
	Value float64
}

func (obj *TestObjective) GetValue() interface{} {
	return obj.Value
}

func (obj *TestObjective) String() string {
	return fmt.Sprintf("%f", obj.Value)
}

func newTestObjectiveWithString(value string) *TestObjective {
	return &TestObjective{
		Value: 0.0,
	}
}

func newTestObjective() *TestObjective {
	return newTestObjectiveWithString("")
}

func TestNewObjective(t *testing.T) {
	newTestObjective()
}
