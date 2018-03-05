// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"fmt"
	"strconv"

	"github.com/cybergarage/foreman-go/foreman"
	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

const (
	errorQueryInvalidStatusCode         = "Expected status code %d != %d (%s)"
	errorQueryEmptyResponseObject       = "Empty path object : %s"
	errorQueryInvalidResponseString     = "Expected '%s' %v != %v"
	errorQueryInvalidResponseInteger    = "Expected '%s' %v != %v"
	errorQueryInvalidResponseReal       = "Expected '%s' %v != %v"
	errorQueryUnknownResponseObjectType = "Unknown path object type %s"
)

// QueryScenario represents a scenario test.
type QueryScenario struct {
	*Scenario
	server *foreman.Server
	client *foreman.Client
}

// NewQueryScenario returns create a query scenario.
func NewQueryScenario() *QueryScenario {
	s := &QueryScenario{
		Scenario: NewScenario(),
		server:   nil,
		client:   foreman.NewClient(),
	}

	s.SetExecutor(s)

	return s
}

// Setup initializes the scenario.
func (s *QueryScenario) Setup() error {
	s.server = foreman.NewServer()

	err := s.server.Start()
	if err != nil {
		return err
	}

	s.client.SetHTTPPort(s.server.GetHTTPPort())

	return nil
}

// Execute runs the specified event.
func (s *QueryScenario) Execute(e *Event) (*Response, *errors.Error) {
	q := NewQueryEvent()
	err := q.ParseEvent(e)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	resObj, resCode, err := s.client.PostQuery(q.Query)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	res := NewQueryResponse()
	res.Query = q.Query
	res.StatusCode = resCode
	res.Content = resObj

	err = s.verifyResponse(res, q.Expectation)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return res.Response, nil
}

// verifyResponse verifies the JSON response with the verify data.
func (s *QueryScenario) verifyResponse(res *QueryResponse, expectRes *QueryExpectation) error {
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
		resFloatValue, err := strconv.ParseFloat(expectRes.JSONPathValue, 64)
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

// Cleanup closes the scenario.
func (s *QueryScenario) Cleanup() error {
	err := s.server.Stop()
	if err != nil {
		return err
	}

	s.server = nil

	return nil
}
