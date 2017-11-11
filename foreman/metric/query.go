// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

import (
	"fmt"
	"time"
)

// Query represents a Foreman Query.
type Query struct {
	Target   string
	From     *time.Time
	Until    *time.Time
	Interval time.Duration
}

// NewQuery returns a new Query.
func NewQuery() *Query {
	q := &Query{
		From:     nil,
		Until:    nil,
		Interval: 0,
	}

	return q
}

// String returns a string description of the instance
func (q *Query) String() string {
	return fmt.Sprintf("%s [%s - %s]", q.Target, q.From.String(), q.Until.String())
}
