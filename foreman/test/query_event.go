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

// QueryResponse represents a scenario query response.
type QueryResponse struct {
	StatusCode     int
	OnlyStatusCode bool
	JSONPath       string
	JSONValue      string
}

// QueryEvent represents a scenario query event.
type QueryEvent struct {
	Query    string
	Response *QueryResponse
}

// NewQueryEvent create an scenario query event
func NewQueryEvent() *QueryEvent {
	q := &QueryEvent{
		Response: &QueryResponse{},
	}
	return q
}

// HasVerifyData returns whether verify data is specified.
func (q *QueryEvent) HasVerifyData() bool {
	return !q.Response.OnlyStatusCode
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
	q.Response.StatusCode = int(code)

	if len(params) < 4 {
		q.Response.OnlyStatusCode = true
		return nil
	}

	q.Response.OnlyStatusCode = false
	q.Response.JSONPath = params[2]
	q.Response.JSONValue = params[3]

	if len(q.Response.JSONPath) <= 0 {
		q.Response.OnlyStatusCode = true
	}

	return nil
}
