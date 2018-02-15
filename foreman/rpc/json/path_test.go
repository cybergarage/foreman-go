// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"testing"
)

func testJSONArray(t *testing.T) {
	testArray := []interface{}{
		[]string{"1", "2", "3"},
		[]int{1, 2, 3},
		[]float64{1, 2, 3},
	}

	e := NewEncorder()
	encStr, err := e.Encode(testArray)
	if err != nil {
		t.Error(err)
		return
	}

	d := NewDecorder()
	rootObj, err := d.Decode(encStr)
	if err != nil {
		t.Error(err)
		return
	}

	path := NewPathWithObject(rootObj)
	strVal, err := path.GetPathString("[0]/[0]")
	if err != nil {
		t.Error(err)
		return
	}
	if strVal != "1" {
		t.Errorf("%s != %s", strVal, "1")
	}

	path = NewPathWithObject(rootObj)
	intVal, err := path.GetPathInteger("[1]/[1]")
	if err != nil {
		t.Error(err)
		return
	}
	if intVal != 2 {
		t.Errorf("%d != %d", intVal, 2)
	}

	path = NewPathWithObject(rootObj)
	realVal, err := path.GetPathFloat("[2]/[2]")
	if err != nil {
		t.Error(err)
		return
	}
	if realVal != 3 {
		t.Errorf("%f != %f", realVal, 3)
	}
}

func testJSONDictionary(t *testing.T) {
	age := "33"
	name := "John Smith"

	personal0 := map[string]string{
		"age":  age,
		"name": name,
	}

	personals := map[string]interface{}{
		"person": personal0,
	}

	e := NewEncorder()
	encStr, err := e.Encode(personals)
	if err != nil {
		t.Error(err)
		return
	}

	d := NewDecorder()
	rootObj, err := d.Decode(encStr)
	if err != nil {
		t.Error(err)
		return
	}

	path := NewPathWithObject(rootObj)
	jsonAge, err := path.GetPathString("person/age")
	if err != nil {
		t.Error(err)
	}
	if jsonAge != age {
		t.Errorf("%s != %s", age, jsonAge)
	}

	jsonName, err := path.GetPathString("person/name")
	if err != nil {
		t.Error(err)
	}
	if jsonName != name {
		t.Errorf("%s != %s", name, jsonName)
	}
}

func TestJSONPath(t *testing.T) {
	testJSONDictionary(t)
	testJSONArray(t)
}
