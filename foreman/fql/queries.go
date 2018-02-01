// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// Queries represents a query list.
type Queries []Query

// NewQueries returns a new query list.
func NewQueries() Queries {
	queries := make([]Query, 0)
	return queries
}
