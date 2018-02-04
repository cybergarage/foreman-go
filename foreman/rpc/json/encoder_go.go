// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"encoding/json"
)

// goEncorder represents an goEncorder for JSON
type goEncorder struct {
	Encorder
}

// NewEncorder return a JSON encorder.
func NewEncorder() Encorder {
	e := &goEncorder{}
	return e
}

// Encode returns a encorded string of the specified object.
func (e *goEncorder) Encode(obj interface{}) (string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(b), err
}
