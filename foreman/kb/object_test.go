// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package kb

import (
	"testing"
)

type TestObject struct {
	Object
	Name  string
	Value float64
}

func (v *TestObject) GetName() string {
	return v.Name
}

func (v *TestObject) GetValue() (interface{}, error) {
	return v.Value, nil
}

func newTestObjectWithName(name string) *TestObject {
	return &TestObject{
		Name:  name,
		Value: 0.0,
	}
}

func newTestObject() *TestObject {
	return newTestObjectWithName("")
}

func TestNewObject(t *testing.T) {
	var v Object
	v = newTestObject()
	v.GetName()
}
