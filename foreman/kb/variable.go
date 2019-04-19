// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// Variable represents a simple operand object which has a name and value.
type Variable interface {
	Object
	// SetName sets a specified name.
	SetName(name string)
	// SetValue sets a specified value.
	SetValue(value interface{}) error
}
