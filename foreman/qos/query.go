// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"fmt"
)

// Query represents a query.
type Query struct {
	Target string
}

// NewQuery returns a new query.
func NewQuery() *Query {
	q := &Query{
		Target: "",
	}

	return q
}

// String returns a string description of the instance
func (q *Query) String() string {
	return fmt.Sprintf("%s", q.Target)
}
