// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// Values is a value array of of FQL.
type Values []*Value

// NewValues returns a new value array.
func NewValues() Values {
	return make(Values, 0)
}
