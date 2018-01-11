// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"github.com/cybergarage/foreman-go/foreman/kb"
)

// Threshold represents an objective instance for Knowledge base.
type Threshold struct {
	kb.Objective
	Value float64
}

// NewThresholdWithValue returns a new threshold instance.
func NewThreshold() *Threshold {
	return NewThresholdWithValue(0.0)
}

// NewThresholdWithValue returns a new threshold instance with the specified value.
func NewThresholdWithValue(value float64) *Threshold {
	th := &Threshold{
		Value: value,
	}
	return th
}

// GetValue returns the stored value.
func (th *Threshold) GetValue() interface{} {
	return th.Value
}
