// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// QueryExpectation represents a scenario Expectation data.
type QueryExpectation struct {
	StatusCode    int
	JSONPath      string
	JSONPathValue string
}

// NewQueryExpectation create an scenario assertion data.
func NewQueryExpectation() *QueryExpectation {
	q := &QueryExpectation{}
	return q
}

// GetStatusCode returns the expected status code.
func (q *QueryExpectation) GetStatusCode() int {
	return q.StatusCode
}

// GetJSONPath returns the expected JSON path.
func (q *QueryExpectation) GetJSONPath() string {
	return q.JSONPath
}

// GetJSONPathValue returns the expected JSON path value.
func (q *QueryExpectation) GetJSONPathValue() string {
	return q.JSONPathValue
}
