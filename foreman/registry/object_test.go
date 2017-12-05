// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

import (
	"fmt"
	"testing"
)

func TestNewObject(t *testing.T) {
}

func TestObjectProperty(t *testing.T) {
	obj := NewObject()

	// Initialize

	propName := make([]string, defaultTestLoopCount)
	propData := make([]string, defaultTestLoopCount)
	for n := 0; n < defaultTestLoopCount; n++ {
		propName[n] = fmt.Sprintf("prop_name%d", n)
		propData[n] = fmt.Sprintf("prop_data%d", n)
	}

	// Insert

	for n := 0; n < defaultTestLoopCount; n++ {
		prop := NewProperty()
		prop.Name = propName[n]
		prop.Data = propData[n]
		obj.SetProperty(prop)
	}

	// Get

	props, err := obj.GetProperties()
	if err != nil {
		t.Error(err)
		return
	}

	if len(props) != defaultTestLoopCount {
		t.Errorf("%d != %d", len(props), defaultTestLoopCount)
	}

	for n := 0; n < defaultTestLoopCount; n++ {
		prop, ok := props.GetPropertyByName(propName[n])
		if !ok {
			t.Errorf("%s is not found", propName[n])
		}
		if prop.Data != propData[n] {
			t.Errorf("%s != %s", prop.Data, propData[n])
		}
	}

	// Insert (Dupulicated)

	for n := 0; n < defaultTestLoopCount; n++ {
		propData[n] = fmt.Sprintf("new_prop_data%d", n)
	}

	for n := 0; n < defaultTestLoopCount; n++ {
		prop := NewProperty()
		prop.Name = propName[n]
		prop.Data = propData[n]
		obj.SetProperty(prop)
	}

	// Get

	props, err = obj.GetProperties()
	if err != nil {
		t.Error(err)
		return
	}

	if len(props) != defaultTestLoopCount {
		t.Errorf("%d != %d", len(props), defaultTestLoopCount)
	}

	for n := 0; n < defaultTestLoopCount; n++ {
		prop, ok := props.GetPropertyByName(propName[n])
		if !ok {
			t.Errorf("%s is not found", propName[n])
		}
		if prop.Data != propData[n] {
			t.Errorf("%s != %s", prop.Data, propData[n])
		}
	}
}
