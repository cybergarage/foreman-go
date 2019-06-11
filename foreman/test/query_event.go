// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

const (
	errorQueryInvalidStatusCode         = "Expected status code %d != %d (%s)"
	errorQueryEmptyResponseObject       = "Empty path object : %s"
	errorQueryInvalidResponseString     = "Expected '%s' %s != %s"
	errorQueryInvalidResponseInteger    = "Expected '%s' %s != %d"
	errorQueryInvalidResponseReal       = "Expected '%s' %s != %f"
	errorQueryUnknownResponseObjectType = "Unknown path object type %s"
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

// VerifyResponse verifies the JSON response with the verify data.
func (q *QueryEvent) VerifyResponse(res *QueryResponse, expectRes *QueryExpectation) error {
	resCode := res.GetStatusCode()
	if resCode != expectRes.StatusCode {
		return fmt.Errorf(errorQueryInvalidStatusCode, resCode, expectRes.StatusCode, res.Query)
	}

	// Check JSON path object

	if len(expectRes.JSONPath) <= 0 {
		return nil
	}

	resObj := res.GetContent()
	jsonPath := json.NewPathWithObject(resObj)
	pathObj, err := jsonPath.GetPathObject(expectRes.JSONPath)
	if err != nil || pathObj == nil {
		return fmt.Errorf(errorQueryEmptyResponseObject, expectRes.JSONPath)
	}

	// Check JSON path value

	if len(expectRes.JSONPathValue) <= 0 {
		return nil
	}

	switch pathObj.(type) {
	case string:
		resValue, _ := pathObj.(string)
		// Parse as a float at first such as '1E+00'
		resFloatValue, err := strconv.ParseFloat(resValue, 64)
		if err == nil {
			expectValue, err := strconv.ParseFloat(expectRes.JSONPathValue, 64)
			if (err == nil) && (resFloatValue == float64(expectValue)) {
				return nil
			}
		}
		// Check as a string
		expectValue := expectRes.JSONPathValue
		if resValue != expectValue {
			return fmt.Errorf(errorQueryInvalidResponseString, expectRes.JSONPath, expectValue, resValue)
		}
	case int:
		resValue, _ := pathObj.(int)
		expectValue, err := strconv.ParseInt(expectRes.JSONPathValue, 10, 64)
		if (err != nil) || (resValue != int(expectValue)) {
			return fmt.Errorf(errorQueryInvalidResponseInteger, expectRes.JSONPath, expectRes.JSONPathValue, resValue)
		}
	case int32:
		resValue, _ := pathObj.(int32)
		expectValue, err := strconv.ParseInt(expectRes.JSONPathValue, 10, 64)
		if (err != nil) || (resValue != int32(expectValue)) {
			return fmt.Errorf(errorQueryInvalidResponseInteger, expectRes.JSONPath, expectRes.JSONPathValue, resValue)
		}
	case int64:
		resValue, _ := pathObj.(int64)
		expectValue, err := strconv.ParseInt(expectRes.JSONPathValue, 10, 64)
		if (err != nil) || (resValue != int64(expectValue)) {
			return fmt.Errorf(errorQueryInvalidResponseInteger, expectRes.JSONPath, expectRes.JSONPathValue, resValue)
		}
	case float32:
		resValue, _ := pathObj.(float32)
		expectValue, err := strconv.ParseFloat(expectRes.JSONPathValue, 64)
		if (err != nil) || (resValue != float32(expectValue)) {
			return fmt.Errorf(errorQueryInvalidResponseReal, expectRes.JSONPath, expectRes.JSONPathValue, resValue)
		}
	case float64:
		resValue, _ := pathObj.(float64)
		expectValue, err := strconv.ParseFloat(expectRes.JSONPathValue, 64)
		if (err != nil) || (resValue != float64(expectValue)) {
			return fmt.Errorf(errorQueryInvalidResponseReal, expectRes.JSONPath, expectRes.JSONPathValue, resValue)
		}
	default:
		return fmt.Errorf(errorQueryUnknownResponseObjectType, expectRes.JSONPath)
	}

	return nil
}
