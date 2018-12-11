// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

// Scenario represents a parameter.
type Scenario struct {
	Name              string
	executor          ScenarioExecutor
	executionListener ScenarioExecutionListener
	Events            []*Event
	LastEvent         *Event
	LastResponse      *Response
	LastError         *errors.Error
}

// NewScenario create a new scenario.
func NewScenario() *Scenario {
	s := &Scenario{
		Name:              "",
		Events:            make([]*Event, 0),
		executor:          nil,
		executionListener: nil,
		LastEvent:         nil,
		LastResponse:      nil,
		LastError:         nil,
	}
	return s
}

// SetName sets the name.
func (s *Scenario) SetName(name string) {
	s.Name = name
}

// SetExecutor sets the event executor.
func (s *Scenario) SetExecutor(e ScenarioExecutor) {
	s.executor = e
}

// SetExecutor sets the event executor.
func (s *Scenario) SetExecutionListener(l ScenarioExecutionListener) {
	s.executionListener = l
}

// SetLastEvent sets a last executed event.
func (s *Scenario) SetLastEvent(e *Event) {
	s.LastEvent = e
}

// SetLastResponse sets a last executed response.
func (s *Scenario) SetLastResponse(r *Response) {
	s.LastResponse = r
}

// SetLastError sets a last error.
func (s *Scenario) SetLastError(e *errors.Error) {
	s.LastError = e
}

// GetName returns the name.
func (s *Scenario) GetName() string {
	return s.Name
}

// GetLastEvent gets the last executed event.
func (s *Scenario) GetLastEvent() *Event {
	return s.LastEvent
}

// GetLastResponse gets the last executed response.
func (s *Scenario) GetLastResponse() *Response {
	return s.LastResponse
}

// GetLastError gets the last error.
func (s *Scenario) GetLastError() *errors.Error {
	return s.LastError
}

// GetEvents returns all scenario events.
func (s *Scenario) GetEvents() []*Event {
	return s.Events
}

// ExecuteAll runs all scenario events.
func (s *Scenario) ExecuteAll(opt *ScenarioOption) *errors.Error {
	err := s.executor.Setup()
	if err != nil {
		return errors.NewErrorWithError(err)
	}

	for _, e := range s.Events {
		s.SetLastEvent(e)

		res, err := s.executor.Execute(e)
		s.SetLastResponse(res)
		s.SetLastError(err)

		if s.executionListener != nil {
			s.executionListener.EventExecuted(e, res, err)
		}

		if err != nil {
			return err
		}

		stepWait := opt.GetStepDuration()
		if 0 < stepWait {
			time.Sleep(stepWait)
		}
	}

	err = s.executor.Cleanup()
	if err != nil {
		return errors.NewErrorWithError(err)
	}

	return nil
}

// LoadFile loads a specified scenario file.
func (s *Scenario) LoadFile(filename string) *errors.Error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.NewErrorWithError(err)
	}

	s.SetName(filepath.Base(filename))

	for n, line := range strings.Split(string(content), "\n") {
		if len(line) <= 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}
		e := NewEvent((n + 1), line)
		if e == nil {
			continue
		}
		s.Events = append(s.Events, e)
	}

	return nil
}

// ExecuteFileWithOption loads a specified scenario file and runs it with an option
func (s *Scenario) ExecuteFileWithOption(filename string, opt *ScenarioOption) *errors.Error {
	err := s.LoadFile(filename)
	if err != nil {
		return err
	}
	return s.ExecuteAll(opt)
}

// ExecuteFile loads a specified scenario file and runs it
func (s *Scenario) ExecuteFile(filename string) *errors.Error {
	return s.ExecuteFileWithOption(filename, NewDefaultScenarioOption())
}
