// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"github.com/cybergarage/foreman-go/foreman/errors"
)

// ScenarioExecutor represents a scenario executor.
type ScenarioExecutor interface {
	// Setup initializes the scenario.
	Setup() error
	// Execute runs the specified event.
	Execute(e *Event) (*Response, *errors.Error)
	// Cleanup closes the scenario.
	Cleanup() error
}

// ScenarioExecutionListener represents a scenario listener
type ScenarioExecutionListener interface {
	EventExecuted(*Event, *Response, *errors.Error)
}
