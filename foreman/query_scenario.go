// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/test"
)

// QueryScenario represents a scenario test.
type QueryScenario struct {
	*test.Scenario
	server *Server
	client *Client
}

// NewQueryScenarioWithServer returns create a query scenario with the specified server.
func NewQueryScenarioWithServer(server *Server) *QueryScenario {
	s := &QueryScenario{
		Scenario: test.NewScenario(),
		server:   server,
		client:   NewClient(),
	}

	s.SetExecutor(s)
	s.SetExecutionListener(s)

	return s
}

// NewQueryScenario returns create a query scenario.
func NewQueryScenario() *QueryScenario {
	return NewQueryScenarioWithServer(NewServer())
}

// Setup initializes the scenario.
func (s *QueryScenario) Setup() error {
	err := s.server.Start()
	if err != nil {
		return err
	}

	s.client.SetNode(s.server)

	return nil
}

// Execute runs the specified event.
func (s *QueryScenario) Execute(e *test.Event) (*test.Response, *errors.Error) {
	var q test.Executor
	var err error

	if e.IsShellEvent() {
		q, err = test.NewShellEventWithEvent(e)
	} else if e.IsSleepEvent() {
		q, err = test.NewSleepEventWithEvent(e)
	} else {
		q, err = NewQueryEventWithEventAndClient(e, s.client)
	}

	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return q.Execute()
}

// EventExecuted are received the execution result.
func (s *QueryScenario) EventExecuted(*test.Event, *test.Response, *errors.Error) {
}

// Cleanup closes the scenario.
func (s *QueryScenario) Cleanup() error {
	err := s.server.Stop()
	if err != nil {
		return err
	}

	return nil
}
