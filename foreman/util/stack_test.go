// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"testing"
)

const (
	errorInvalid                = "Invalid parser object (%d != %d)"
	errorInvalidParserStackSize = "Invalid parser stack size (%d != %d)"
)

func TestNewStack(t *testing.T) {
	s := NewStack()

	if s.Size() != 0 {
		t.Errorf(errorInvalidParserStackSize, s.Size(), 0)
	}

	for n := 1; n <= 10; n++ {
		var value int
		value = n
		s.Push(value)
		if s.Size() != n {
			t.Errorf(errorInvalidParserStackSize, s.Size(), n)
		}
	}

	for n := 10; n >= 1; n-- {
		obj := s.Pop()
		value, _ := obj.(int)
		if value != n {
			t.Errorf(errorInvalid, value, n)
		}
		if s.Size() != (n - 1) {
			t.Errorf(errorInvalidParserStackSize, s.Size(), (n - 1))
		}
	}

	if s.Size() != 0 {
		t.Errorf(errorInvalidParserStackSize, s.Size(), 0)
	}
}
