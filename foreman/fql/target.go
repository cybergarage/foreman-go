// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// Target is a destination or source target of FQL.
type Target struct {
	Value string
}

// NewTargetWithString returns a new target with the specified string.
func NewTargetWithString(targetString string) *Target {
	t := &Target{
		Value: targetString,
	}
	return t
}

// NewTarget returns a new target.
func NewTarget() *Target {
	return NewTargetWithString("")
}

// String returns the target value.
func (t *Target) String() string {
	return t.Value
}
