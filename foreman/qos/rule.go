// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

// Rule represents a QoS rule.
type Rule struct {
	Clauses []Clause
}

// NewRule returns a new null rule.
func NewRule() *Rule {
	p := &Rule{}
	return p
}

// ParseString parses a specified string.
func (rule *Rule) ParseString(ruleString string) error {
	return nil
}
