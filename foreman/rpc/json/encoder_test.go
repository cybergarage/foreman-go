// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"testing"
)

func testObjectToJSON(t *testing.T, obj interface{}, jsonStr string) {
	e := NewEncorder()
	encodedStr, err := e.Encode(obj)
	if err != nil {
		t.Error(e)
		return
	}

	if jsonStr != encodedStr {
		t.Errorf("%s (%s != %s)", obj, jsonStr, encodedStr)
	}
}

func TestJSONArrayEncode(t *testing.T) {
	obj := []string{"milk", "bread", "eggs"}
	jsonStr := "[\"milk\",\"bread\",\"eggs\"]"
	testObjectToJSON(t, obj, jsonStr)
}

func TestJSONMapEncode(t *testing.T) {
	obj := map[string]string{
		"age":  "33",
		"name": "John Smith",
	}
	jsonStr := "{\"age\":\"33\",\"name\":\"John Smith\"}"
	testObjectToJSON(t, obj, jsonStr)
}
