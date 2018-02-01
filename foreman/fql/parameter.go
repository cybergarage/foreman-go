// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// Parameter represents a parameter interface.
type Parameter interface {
	GetName() string
	GetValue() interface{}
	GetString() (string, bool)
	String() string
}
