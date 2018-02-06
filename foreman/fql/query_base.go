// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// baseQuery represents a parameter.
type baseQuery struct {
	token string
	Parameters
	Conditions
}

// newQueryWithToken returns a new query with the specified string.
func newBaseQuery() *baseQuery {
	q := &baseQuery{
		token:      "",
		Parameters: NewParameters(),
		Conditions: NewConditions(),
	}
	return q
}

// GetType returns a stored type in the query.
func (q *baseQuery) GetType() QueryType {
	return QueryTypeUnknown
}

// AddCondition adds a new condition.
func (q *baseQuery) AddCondition(c *Condition) error {
	q.Conditions = append(q.Conditions, c)
	return nil
}

// GetConditions retusn all conditions.
func (q *baseQuery) GetConditions() Conditions {
	return q.Conditions
}

// SetTarget sets a specified target into the query.
func (q *baseQuery) SetTarget(target Target) error {
	param := NewParameterWithObject(parameterTarget, target)
	return q.SetParameter(param)
}

// GetTarget returns the target in the query.
func (q *baseQuery) GetTarget() (Target, bool) {
	return q.GetParameterString(parameterTarget)
}

// SetValues sets a specified values intp the query.
func (q *baseQuery) SetValues(values Values) error {
	param := NewParameterWithObject(parameterValues, values)
	return q.SetParameter(param)
}

// AddValue adds a specified value into the query.
func (q *baseQuery) AddValue(value Value) error {
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
	values = append(values, value)
	return q.SetValues(values)
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
