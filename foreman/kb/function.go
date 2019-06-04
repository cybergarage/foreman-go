// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// Function represents a function operand object.
type Function interface {
	Object
	// SetName sets a specified name.
	SetName(name string)
	// SetParameters sets a specified parameters.
	SetParameters(params []interface{})
	// Execute returns the operand value with the specified parameters.
	Execute([]interface{}) (interface{}, error)
}