// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// baseQuery represents a parameter.
type baseQuery struct {
	*Target
	Values
	Conditions
}

// newQueryWithToken returns a new query with the specified string.
func newBaseQuery() *baseQuery {
	q := &baseQuery{
		Target:     nil,
		Values:     NewValues(),
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
func (q *baseQuery) SetTarget(target *Target) error {
	q.Target = target
	return nil
}

// GetTarget returns the target in the query.
func (q *baseQuery) GetTarget() (*Target, bool) {
	if q.Target == nil {
		return nil, false
	}
	return q.Target, true
}

// AddValue adds a specified value into the query.
func (q *baseQuery) AddValue(value *Value) error {
	q.Values = append(q.Values, value)
	return nil
}

// GetValues returns the values in the query.
func (q *baseQuery) GetValues() (Values, bool) {
	if len(q.Values) <= 0 {
		return nil, false
	}
	return q.Values, true
}
