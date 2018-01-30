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

type testObject struct {
	Object
	Data        string
	UpdateCount int
}

func newTestObject() *testObject {
	obj := &testObject{
		Object:      NewObject(),
		Data:        "",
		UpdateCount: 0,
	}
	return obj
}

func (obj *testObject) ObjectUpdated() {
	obj.UpdateCount++
}

func TestObjectListener(t *testing.T) {
	obj := newTestObject()

	if len(obj.GetListeners()) != 0 {
		t.Errorf(errorInvalidListenerCount, len(obj.GetListeners()), 0)
	}

	obj.AddListener(obj)
	if len(obj.GetListeners()) != 1 {
		t.Errorf(errorInvalidListenerCount, len(obj.GetListeners()), 1)
	}

	// Adds same listener
	obj.AddListener(obj)
	if len(obj.GetListeners()) != 1 {
		t.Errorf(errorInvalidListenerCount, len(obj.GetListeners()), 1)
	}

	updateLoopCount := 10
	for n := 0; n < updateLoopCount; n++ {
		obj.UpdateVersion()
		if obj.UpdateCount != (n + 1) {
			t.Errorf(errorInvalidListenerUpdateCount, obj.UpdateCount, (n + 1))
		}
	}
}
