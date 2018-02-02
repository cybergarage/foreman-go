// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"testing"
)

const (
	errorInvalidParserObject    = "Invalid parser object (%d != %d)"
	errorInvalidParserStackSize = "Invalid parser stack size (%d != %d)"
)

func TestNewParserObjectStack(t *testing.T) {
	s := NewParserObjectStack()

	if s.Size() != 0 {
		t.Errorf(errorInvalidParserStackSize, s.Size(), 0)
	}

	for n := 1; n <= 10; n++ {
		var value int
		value = n
		s.PushObject(value)
		if s.Size() != n {
			t.Errorf(errorInvalidParserStackSize, s.Size(), n)
		}
	}

	for n := 10; n >= 1; n-- {
		obj := s.PopObject()
		value, _ := obj.(int)
		if value != n {
			t.Errorf(errorInvalidParserObject, value, n)
		}
		if s.Size() != (n - 1) {
			t.Errorf(errorInvalidParserStackSize, s.Size(), (n - 1))
		}
	}

	if s.Size() != 0 {
		t.Errorf(errorInvalidParserStackSize, s.Size(), 0)
	}
}
