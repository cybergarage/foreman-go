// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"testing"
)

func TestJSONMapDencode(t *testing.T) {
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
	err = d.Decode(encStr)
	if err != nil {
		t.Error(err)
		return
	}

	jsonAge, err := d.GetPathString("person/age")
	if err != nil {
		t.Error(err)
	}
	if jsonAge != age {
		t.Errorf("%s != %s", age, jsonAge)
	}

	jsonName, err := d.GetPathString("person/name")
	if err != nil {
		t.Error(err)
	}
	if jsonName != name {
		t.Errorf("%s != %s", name, jsonName)
	}
}
