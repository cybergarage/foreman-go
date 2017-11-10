// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

import (
	"fmt"
)

// Query represents a Foreman Query.
type Query struct {
	ParentID string
}

// NewQuery returns a new Query.
func NewQuery() *Query {
	q := &Query{
		ParentID: "",
	}

	return q
}

// String returns a string description of the instance
func (self *Query) String() string {
	return fmt.Sprintf("%s", self.ParentID)
}
