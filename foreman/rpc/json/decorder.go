// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

// Decorder represents a decorder for JSON
type Decorder interface {
	Decode(json string) error
	GetRootObject() (interface{}, error)
	GetPathObject(path string) (interface{}, error)
	GetPathString(path string) (string, error)
	GetPathInteger(path string) (int, error)
	GetPathFloat(path string) (float64, error)
}
