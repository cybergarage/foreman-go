// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// Columns is a column array of of FQL.
type Columns = []*Column

// NewColumns returns a new Column array.
func NewColumns() Columns {
	return make(Columns, 0)
}
