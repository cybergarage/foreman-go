// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"fmt"
	"strconv"

	"github.com/cybergarage/foreman-go/foreman"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

const (
	errorQueryInvalidStatusCode         = "Expected status code %d != %d (%s)"
	errorQueryEmptyResponseObject       = "Empty '%s'"
	errorQueryInvalidResponseString     = "Expected '%s' %s != %s"
	errorQueryInvalidResponseInteger    = "Expected '%s' %d != %s"
	errorQueryInvalidResponseReal       = "Expected '%s' %f != %s"
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
		server:   foreman.NewServer(),
		client:   foreman.NewClient(),
	}

	s.SetExecutor(s)

	return s
}

// Setup initializes the scenario.
func (s *QueryScenario) Setup() error {
	err := s.server.Start()
	if err != nil {
		return err
	}

	s.client.SetHTTPPort(s.server.GetHTTPPort())

	return nil
}

// Execute runs the specified event.
func (s *QueryScenario) Execute(e *Event) error {
	q := NewQueryEvent()
	err := q.ParseEvent(e)
	if err != nil {
		return err
	}

	resObj, resCode, err := s.client.PostQuery(q.Query)
	if err != nil {
		return err
	}

	if resCode != q.Response.StatusCode {
		return fmt.Errorf(errorQueryInvalidStatusCode, resCode, q.Response.StatusCode, q.Query)
	}

	if !q.HasVerifyData() {
		return nil
	}

	err = s.verifyResponse(resObj, q.Response)
	if err != nil {
		return err
	}

	return nil
}

// verifyResponse verifies the JSON response with the verify data.
func (s *QueryScenario) verifyResponse(resObj interface{}, expectRes *QueryResponse) error {
	jsonPath := json.NewPathWithObject(resObj)
	pathObj, err := jsonPath.GetPathObject(expectRes.JSONPath)
	if err != nil {
		return fmt.Errorf(errorQueryEmptyResponseObject, expectRes.JSONPath)
	}

	switch pathObj.(type) {
	case string:
		resValue, _ := pathObj.(string)
		expectValue := expectRes.JSONValue
		if resValue != expectValue {
			return fmt.Errorf(errorQueryInvalidResponseString, expectRes.JSONPath, resValue, expectValue)
		}
	case int:
		resValue, _ := pathObj.(int)
		expectValue, err := strconv.ParseInt(expectRes.JSONValue, 10, 64)
		if (err != nil) || (resValue != int(expectValue)) {
			return fmt.Errorf(errorQueryInvalidResponseInteger, expectRes.JSONPath, resValue, expectRes.JSONValue)
		}
	case int32:
		resValue, _ := pathObj.(int32)
		expectValue, err := strconv.ParseInt(expectRes.JSONValue, 10, 64)
		if (err != nil) || (resValue != int32(expectValue)) {
			return fmt.Errorf(errorQueryInvalidResponseInteger, expectRes.JSONPath, resValue, expectRes.JSONValue)
		}
	case int64:
		resValue, _ := pathObj.(int64)
		expectValue, err := strconv.ParseInt(expectRes.JSONValue, 10, 64)
		if (err != nil) || (resValue != int64(expectValue)) {
			return fmt.Errorf(errorQueryInvalidResponseInteger, expectRes.JSONPath, resValue, expectRes.JSONValue)
		}
	case float32:
		resValue, _ := pathObj.(float32)
		expectValue, err := strconv.ParseFloat(expectRes.JSONValue, 64)
		if (err != nil) || (resValue != float32(expectValue)) {
			return fmt.Errorf(errorQueryInvalidResponseReal, expectRes.JSONPath, resValue, expectRes.JSONValue)
		}
	case float64:
		resValue, _ := pathObj.(float64)
		expectValue, err := strconv.ParseFloat(expectRes.JSONValue, 64)
		if (err != nil) || (resValue != float64(expectValue)) {
			return fmt.Errorf(errorQueryInvalidResponseReal, expectRes.JSONPath, resValue, expectRes.JSONValue)
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
