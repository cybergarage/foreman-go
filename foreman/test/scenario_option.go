// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"time"
)

// ScenarioOption represents a scenario option.
type ScenarioOption struct {
	stepDuration time.Duration
}

// NewDefaultScenarioOption create a default scenario option.
func NewDefaultScenarioOption() *ScenarioOption {
	opt := &ScenarioOption{
		stepDuration: 0,
	}
	return opt
}

// SetStepDuration sets a sleep duration each step.
func (opt *ScenarioOption) SetStepDuration(d time.Duration) {
	opt.stepDuration = d
}

// GetStepDuration return the sleep duration each step.
func (opt *ScenarioOption) GetStepDuration() time.Duration {
	return opt.stepDuration
}
