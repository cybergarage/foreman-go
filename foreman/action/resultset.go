// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// ResultSet is an alias of Parameters
type ResultSet Parameters

// NewResultSet returns a new result set.
func NewResultSet() ResultSet {
	return NewParameters()
}
