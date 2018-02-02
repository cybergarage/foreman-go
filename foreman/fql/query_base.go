// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// baseQuery represents a parameter.
type baseQuery struct {
	token string
	Parameters
}

// newQueryWithToken returns a new query with the specified string.
func newBaseQuery() *baseQuery {
	q := &baseQuery{
		Parameters: NewParameters(),
	}
	return q
}

// GetType returns a stored type in the query.
func (q *baseQuery) GetType() QueryType {
	return QueryTypeUnknown
}

// GetTarget returns the target in the query.
func (q *baseQuery) GetTarget() (Target, bool) {
	return q.GetParameterString(parameterTable)
}

// AddValue adds a specified value into the query.
func (q *baseQuery) AddValue(value Value) bool {
	var values Values
	param, ok := q.GetParameter(parameterValues)
	if ok {
		values, ok = param.GetValue().(Values)
		if !ok {
			values = NewValues()
		}
	} else {
		values = NewValues()
	}
	err := param.SetValue(values)
	if err != nil {
		return false
	}
	err = q.SetParameter(param)
	if err != nil {
		return false
	}
	return true
}

// GetValues returns the values in the query.
func (q *baseQuery) GetValues() (Values, bool) {
	param, ok := q.GetParameter(parameterValues)
	if !ok {
		return nil, false
	}
	values, ok := param.GetValue().(Values)
	if !ok {
		return nil, false
	}
	return values, true
}
