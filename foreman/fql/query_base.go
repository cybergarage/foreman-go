// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"fmt"
	"strings"
)

// baseQuery represents a parameter.
type baseQuery struct {
	Type QueryType
	*Target
	Columns
	Values
	Conditions
	forwardingFlag bool
}

// newQueryWithToken returns a new query with the specified string.
func newBaseQueryWithType(queryType QueryType) *baseQuery {
	q := &baseQuery{
		Type:           queryType,
		Target:         nil,
		Columns:        NewColumns(),
		Values:         NewValues(),
		Conditions:     NewConditions(),
		forwardingFlag: false,
	}
	return q
}

// GetType returns a query type.
func (q *baseQuery) GetType() QueryType {
	return q.Type
}

// IsStateChangeQuery returns whether state change query
func (q *baseQuery) IsStateChangeQuery() bool {
	if (q.Type == QueryTypeInsert) || (q.Type == QueryTypeDelete) {
		return true
	}
	return false
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

// HasColumn returns true when the query has the specified column.
func (q *baseQuery) HasColumn(name string) bool {
	for _, c := range q.Columns {
		if c.String() == name {
			return true
		}
	}
	return false
}

// HasOnlyColumn returns true when the query has only the specified column.
func (q *baseQuery) HasOnlyColumn(name string) bool {
	if len(q.Columns) != 1 {
		return false
	}
	if q.Columns[0].String() != name {
		return false
	}
	return true
}

// IsAllColumn returns true when the query has only "*" column.
func (q *baseQuery) IsAllColumn() bool {
	if len(q.Columns) == 0 {
		return true
	}
	return q.HasColumn(QueryColumnAll)
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

// SetRetransmissionFlag sets a retransmission flag.
func (q *baseQuery) SetRetransmissionFlag(flag bool) {
	q.forwardingFlag = flag
}

// IsRetransmissionQuery returns whether retransmission query
func (q *baseQuery) IsRetransmissionQuery() bool {
	return q.forwardingFlag
}

// String returns a string description of the instance
func (q *baseQuery) String() string {

	// Query Type

	queryType := QueryTypeUnknown
	qt, ok := (interface{}(q)).(Query)
	if ok {
		queryType = qt.GetType()
	}

	// Query String
	var queryString string

	queryTypeStrings := map[QueryType]string{
		QueryTypeInsert:  QueryInsertString,
		QueryTypeSelect:  QuerySelectString,
		QueryTypeUpdate:  QueryUpdateString,
		QueryTypeDelete:  QueryDeleteString,
		QueryTypeAnalyze: QueryAnalyzeString,
		QueryTypeExecute: QueryExecuteString,
	}
	queryTypeString, ok := queryTypeStrings[queryType]
	if ok {
		queryString += fmt.Sprintf("%s", queryTypeString)
	}

	// Columns (Only Select)

	switch queryType {
	case QueryTypeSelect:
		if q.IsAllColumn() {
			queryString += fmt.Sprintf(" %s", QueryColumnAll)
		} else {
			columns, ok := q.GetColumns()
			if ok {
				queryString += " ("
				for n, column := range columns {
					if 0 < n {
						queryString += ", "
					}
					queryString += column.String()
				}
				queryString += ")"
			}
		}
	}

	// Target

	target, ok := q.GetTarget()
	if ok {
		switch queryType {
		case QueryTypeInsert:
			queryString += fmt.Sprintf(" INTO %s", target)
		case QueryTypeSelect, QueryTypeDelete:
			queryString += fmt.Sprintf(" FROM %s", target)
		case QueryTypeUpdate:
			queryString += fmt.Sprintf(" %s", target)
		}
	}

	// Columns and Values (Only Insert and Update)

	switch queryType {
	case QueryTypeInsert:
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

		values, ok := q.GetValues()
		if ok {
			queryString += " VALUES ("
			for n, value := range values {
				if 0 < n {
					queryString += ", "
				}
				valueString := value.String()
				if 0 <= strings.IndexAny(valueString, " ") {
					valueString = fmt.Sprintf("\"%s\"", valueString)
				}
				queryString += valueString
			}
			queryString += ")"
		}
	case QueryTypeUpdate:
		columns, _ := q.GetColumns()
		values, _ := q.GetValues()
		if len(columns) == len(values) {
			if 0 < len(columns) {
				queryString += " SET "
				for n, column := range columns {
					if 0 < n {
						queryString += ", "
					}
					value := values[n]
					queryString += fmt.Sprintf("%s = %s", column.String(), value.String())
				}
			}
		}
	}

	// Conditions

	conditions, ok := q.GetConditions()
	if ok {
		queryString += " WHERE "
		for _, condition := range conditions {
			queryString += fmt.Sprintf("%s ", condition.String())
		}
	}

	return queryString
}
