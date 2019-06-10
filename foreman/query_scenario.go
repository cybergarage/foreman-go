// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"strconv"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
	"github.com/cybergarage/foreman-go/foreman/test"
)

const (
	errorQueryInvalidStatusCode         = "Expected status code %d != %d (%s)"
	errorQueryEmptyResponseObject       = "Empty path object : %s"
	errorQueryInvalidResponseString     = "Expected '%s' %s != %s"
	errorQueryInvalidResponseInteger    = "Expected '%s' %s != %d"
	errorQueryInvalidResponseReal       = "Expected '%s' %s != %f"
	errorQueryUnknownResponseObjectType = "Unknown path object type %s"
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

// executeQuery runs the specified query event.
func (s *QueryScenario) executeQuery(e *test.Event) (*test.Response, *errors.Error) {
	q, err := test.NewQueryEventWithEvent(e)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	resObj, resCode, err := s.client.PostQueryOverHTTP(q.Query)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	res := test.NewQueryResponse()
	res.Query = q.Query
	res.StatusCode = resCode
	res.Content = resObj

	err = s.verifyResponse(res, q.Expectation)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return res.Response, nil
}

// executeCommand runs the specified shell event.
func (s *QueryScenario) executeCommand(e *test.Event) (*test.Response, *errors.Error) {
	q, err := test.NewShellEventWithEvent(e)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}
	return q.Execute()
}

// Execute runs the specified event.
func (s *QueryScenario) Execute(e *test.Event) (*test.Response, *errors.Error) {
	if e.HasDataPrefix("SHELL") {
		return s.executeCommand(e)
	}
	return s.executeQuery(e)
}

// EventExecuted are received the execution result.
func (s *QueryScenario) EventExecuted(*test.Event, *test.Response, *errors.Error) {

}

// verifyResponse verifies the JSON response with the verify data.
func (s *QueryScenario) verifyResponse(res *test.QueryResponse, expectRes *test.QueryExpectation) error {
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

// Cleanup closes the scenario.
func (s *QueryScenario) Cleanup() error {
	err := s.server.Stop()
	if err != nil {
		return err
	}

	return nil
}
