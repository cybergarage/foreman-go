// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	errorEmptyEvent                 = "Empty event"
	errorInvalidEventParameterCount = "Invalid event parameter count (%d) : %s"
	errorInvalidEventCode           = "Invalid event code : %s"
)

// QueryEvent represents a scenario query event.
type QueryEvent struct {
	Query       string
	Expectation *QueryExpectation
}

// NewQueryEvent create an scenario query event
func NewQueryEvent() *QueryEvent {
	q := &QueryEvent{
		Expectation: NewQueryExpectation(),
	}
	return q
}

// NewQueryEventWithEvent create an scenario query event
func NewQueryEventWithEvent(e *Event) (*QueryEvent, error) {
	q := NewQueryEvent()
	err := q.ParseEvent(e)
	if err != nil {
		return nil, err
	}
	return q, nil
}

// ParseEvent parses the specified event.
func (q *QueryEvent) ParseEvent(e *Event) error {
	eventData := e.Data
	if len(eventData) <= 0 {
		return errors.New(errorEmptyEvent)
	}

	params := strings.Split(eventData, ";")
	if len(params) < 2 {
		return fmt.Errorf(errorInvalidEventParameterCount, len(params), e.Data)
	}

	q.Query = params[0]

	code, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return fmt.Errorf(errorInvalidEventCode, e.Data)
	}
	q.Expectation.StatusCode = int(code)

	if len(params) < 3 {
		return nil
	}
	q.Expectation.JSONPath = params[2]

	if len(params) < 4 {
		return nil
	}
	q.Expectation.JSONPathValue = params[3]

	return nil
}
