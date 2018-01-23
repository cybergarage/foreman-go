// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

import (
	"testing"
)

const (
	errorInvalidListenerCount       = "Invalid listener count (%d != %d)"
	errorInvalidListenerUpdateCount = "Invalid listener update count (%d != %d)"
)

type testObjectListener struct {
	UpdateCount int
}

func newTestObjectListener() *testObjectListener {
	l := &testObjectListener{
		UpdateCount: 0,
	}
	return l
}

func (l *testObjectListener) ObjectUpdated(obj *Object) {
	l.UpdateCount++
}

func TestObjectListener(t *testing.T) {
	obj := NewObject()

	if len(obj.GetListeners()) != 0 {
		t.Errorf(errorInvalidListenerCount, len(obj.GetListeners()), 0)
	}

	l := newTestObjectListener()
	obj.AddListener(l)
	if len(obj.GetListeners()) != 1 {
		t.Errorf(errorInvalidListenerCount, len(obj.GetListeners()), 1)
	}

	// Adds same listener
	obj.AddListener(l)
	if len(obj.GetListeners()) != 1 {
		t.Errorf(errorInvalidListenerCount, len(obj.GetListeners()), 1)
	}

	updateLoopCount := 10
	for n := 0; n < updateLoopCount; n++ {
		obj.UpdateVersionAndTimestamp()
		if l.UpdateCount != (n + 1) {
			t.Errorf(errorInvalidListenerUpdateCount, l.UpdateCount, (n + 1))
		}
	}
}
