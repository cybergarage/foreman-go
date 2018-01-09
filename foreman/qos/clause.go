// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

// A Clause represents a clause in a QoS rule.
type Clause struct {
	qualities []Quality
}

// NewClause returns a new clause.
func NewClause() *Clause {
	c := &Clause{}
	return c
}
