// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"fmt"
)

// baseQuery represents a parameter.
type baseQuery struct {
	*Target
	Columns
	Values
	Conditions
}

// newQueryWithToken returns a new query with the specified string.
func newBaseQuery() *baseQuery {
	q := &baseQuery{
		Target:     nil,
		Columns:    NewColumns(),
		Values:     NewValues(),
		Conditions: NewConditions(),
	}
	return q
}

// GetType returns a stored type in the query.
func (q *baseQuery) GetType() QueryType {
	return QueryTypeUnknown
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

// AddColumn adds a specified column into the query.
func (q *baseQuery) AddColumn(column *Column) error {
	q.Columns = append(q.Columns, column)
	return nil
}

// GetColumns returns the columns in the query.
func (q *baseQuery) GetColumns() (Columns, bool) {
	if len(q.Columns) <= 0 {
		return q.Columns, false
	}
	return q.Columns, true
}

// HasColumn returns the columns in the query.
func (q *baseQuery) HasColumn(name string) bool {
	for _, c := range q.Columns {
		if c.String() == name {
			return true
		}
	}
	return false
}

// HasOnlyColumn returns the columns in the query.
func (q *baseQuery) HasOnlyColumn(name string) bool {
	if len(q.Columns) != 1 {
		return false
	}
	if q.Columns[0].String() != name {
		return false
	}
	return true
}

// AddValue adds a specified value into the query.
func (q *baseQuery) AddValue(value *Value) error {
	q.Values = append(q.Values, value)
	return nil
}

// GetValues returns the values in the query.
func (q *baseQuery) GetValues() (Values, bool) {
	if len(q.Values) <= 0 {
		return q.Values, false
	}
	return q.Values, true
}

// AddCondition adds a new condition.
func (q *baseQuery) AddCondition(c *Condition) error {
	q.Conditions = append(q.Conditions, c)
	return nil
}

// GetConditions retusn all conditions.
func (q *baseQuery) GetConditions() (Conditions, bool) {
	if len(q.Conditions) <= 0 {
		return q.Conditions, false
	}
	return q.Conditions, true
}

// GetConditionByColumnName retusn the operator and right operand of the specified left operator.
func (q *baseQuery) GetConditionByColumn(leftOpe string) (*Operator, string, bool) {
	for _, c := range q.Conditions {
		if c.GetColumn() == leftOpe {
			return c.GetOperator(), c.GetOperand(), true
		}
	}
	return nil, "", false
}

// String returns a string description of the instance
func (q *baseQuery) String() string {

	var queryString string

	queryTypeStrings := map[QueryType]string{
		QueryTypeInsert:  QueryInsertString,
		QueryTypeSelect:  QuerySelectString,
		QueryTypeDelete:  QueryDeleteString,
		QueryTypeAnalyze: QueryAnalyzeString,
		QueryTypeExecute: QueryExecuteString,
	}
	queryTypeString, ok := queryTypeStrings[q.GetType()]
	if ok {
		queryString += fmt.Sprintf("%s", queryTypeString)
	}

	// Columns (Select)

	if q.GetType() == QueryTypeSelect {
		columns, ok := q.GetColumns()
		if ok {
			queryString += "("
			for n, column := range columns {
				if 0 < n {
					queryString += ", "
				}
				queryString += column.String()
			}
			queryString += ")"
		}
	}

	// Target

	target, ok := q.GetTarget()
	if ok {
		switch q.GetType() {
		case QueryTypeInsert:
			queryString += fmt.Sprintf("INTO %s", target)
		case QueryTypeSelect, QueryTypeDelete:
			queryString += fmt.Sprintf("FROM %s", target)
		}
	}

	// Columns (Insert)

	if q.GetType() == QueryTypeInsert {
		columns, ok := q.GetColumns()
		if ok {
			queryString += "("
			for n, column := range columns {
				if 0 < n {
					queryString += ", "
				}
				queryString += column.String()
			}
			queryString += ")"
		}
	}

	// Values (Insert)

	if q.GetType() == QueryTypeInsert {
		values, ok := q.GetValues()
		if ok {
			queryString += "VALUES ("
			for n, value := range values {
				if 0 < n {
					queryString += ", "
				}
				queryString += value.String()
			}
			queryString += ")"
		}
	}

	// Conditions

	conditions, ok := q.GetConditions()
	if ok {
		queryString += "WHERE "
		for _, condition := range conditions {
			queryString += fmt.Sprintf("%s ", condition.String())
		}
	}

	return queryString
}
