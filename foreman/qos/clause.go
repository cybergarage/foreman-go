// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

// A Clause represents a clause in a QoS rule.
type Clause struct {
	Qualities []*Quality
}

// NewClause returns a new clause.
func NewClause() *Clause {
	clause := &Clause{
		Qualities: make([]*Quality, 0),
	}
	return clause
}

// AddQuality adds a nre quality.
func (clause *Clause) AddQuality(quality *Quality) error {
	clause.Qualities = append(clause.Qualities, quality)
	return nil
}
