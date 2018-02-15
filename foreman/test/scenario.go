// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// ScenarioExecutor represents a scenario test.
type ScenarioExecutor interface {
	// Setup initializes the scenario.
	Setup() error
	// Execute runs the specified event.
	Execute(e *Event) error
	// Finalize closes the scenario.
	Finalize() error
}

// Scenario represents a parameter.
type Scenario struct {
	executor ScenarioExecutor
	Events   []*Event
}

// NewScenario create a new scenario.
func NewScenario() *Scenario {
	s := &Scenario{
		Events:   make([]*Event, 0),
		executor: nil,
	}
	return s
}

// SetExecutor set the event executor.
func (s *Scenario) SetExecutor(e ScenarioExecutor) {
	s.executor = e
}

// GetEvents returns all scenario events.
func (s *Scenario) GetEvents() []*Event {
	return s.Events
}

// ExecuteAll runs all scenario events.
func (s *Scenario) ExecuteAll() error {
	err := s.executor.Setup()
	if err != nil {
		return err
	}

	for n, e := range s.Events {
		err := s.executor.Execute(e)
		if err != nil {
			return fmt.Errorf("LINE [%d] %s : %s", n, e.Data, err.Error())
		}
	}

	err = s.executor.Finalize()
	if err != nil {
		return err
	}

	return nil
}

// LoadFile loads a specified scenario file.
func (s *Scenario) LoadFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(content), "\n") {
		if len(line) <= 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}
		e := NewEvent(line)
		if e == nil {
			continue
		}
		s.Events = append(s.Events, e)
	}

	return nil
}

// ExecuteFile loads a specified scenario file and runs it
func (s *Scenario) ExecuteFile(filename string) error {
	err := s.LoadFile(filename)
	if err != nil {
		return err
	}
	return s.ExecuteAll()
}
