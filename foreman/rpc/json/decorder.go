// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"encoding/json"
)

// Decorder represents a decorder for JSON
type Decorder struct {
	*Path
}

// NewDecorder return a JSON Decorder.
func NewDecorder() *Decorder {
	d := &Decorder{
		Path: newPath(),
	}
	return d
}

// Decode decordes the specified JSON string.
func (d *Decorder) Decode(jsonString string) (interface{}, error) {
	err := json.Unmarshal([]byte(jsonString), &d.Path.rootObject)
	if err != nil {
		return nil, err
	}
	return d.Path.rootObject, nil
}
