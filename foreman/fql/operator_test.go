// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"testing"
)

func TestNewOperators(t *testing.T) {
	testOpeStrings := []string{
		QueryConditionOperatorLt,
		QueryConditionOperatorLe,
		QueryConditionOperatorGt,
		QueryConditionOperatorGe,
		QueryConditionOperatorEq,
		QueryConditionOperatorNeq,
	}

	for _, opeStr := range testOpeStrings {
		ope := NewOperatorWithString(opeStr)
		if ope.GetType() == OperatorTypeUnknown {
			t.Errorf("%s == %d", opeStr, ope.GetType())
		}
	}
}
