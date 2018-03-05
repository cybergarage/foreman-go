// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cybergarage/foreman-go/foreman/fql"
)

type QueryType int

const (
	QueryTypeUnknown QueryType = iota
	QueryTypeInsert
	QueryTypeSelect
	QueryTypeAnalyze
)

type QuerySourceType int

const (
	QuerySourceTypeUnknown QuerySourceType = iota
	QuerySourceTypeMetric
	QuerySourceDataType
)

// Query represents a query for the metric store.
type Query struct {
	Type     QueryType
	Source   QuerySourceType
	Target   string
	From     *time.Time
	Until    *time.Time
	Interval time.Duration
}

// NewMetricQuery returns a new metric query.
func NewMetricQuery() *Query {
	q := &Query{
		Type:     QueryTypeSelect,
		Source:   QuerySourceTypeMetric,
		Target:   "",
		From:     nil,
		Until:    nil,
		Interval: 0,
	}
	return q
}

// NewDataQuery returns a new metric query.
func NewDataQuery() *Query {
	now := time.Now()
	from := now.Add(QueryDefaultFromOffset)
	q := &Query{
		Type:     QueryTypeSelect,
		Source:   QuerySourceDataType,
		Target:   "",
		From:     &from,
		Until:    &now,
		Interval: 0,
	}
	return q
}

// NewAnalyzeQuery returns a new metric query.
func NewAnalyzeQuery() *Query {
	now := time.Now()
	from := now.Add(QueryDefaultFromOffset)
	q := &Query{
		Type:     QueryTypeAnalyze,
		Source:   QuerySourceDataType,
		Target:   "",
		From:     &from,
		Until:    &now,
		Interval: 0,
	}
	return q
}

// NewQueryWithQuery returns a new query of the specified query.
func NewQueryWithQuery(fq fql.Query) (*Query, error) {
	var q *Query

	switch fq.GetType() {
	case fql.QueryTypeInsert:
		// TODO : Not implemented yet
	case fql.QueryTypeSelect:
		if fq.HasOnlyColumn(fql.QueryColumnName) {
			q = NewMetricQuery()
		} else {
			q = NewDataQuery()
		}
	case fql.QueryTypeAnalyze:
		q = NewAnalyzeQuery()
	}

	// Unknown query

	if q == nil {
		return nil, fmt.Errorf(errorStoreInvalidQuery, fq.String())
	}

	// Column (only for ANALYZE)

	columns, ok := fq.GetColumns()
	if ok {
		for _, column := range columns {
			switch fq.GetType() {
			case fql.QueryTypeAnalyze:
				q.Target = column.String()
			}
		}
	}

	// Where

	wheres, ok := fq.GetConditions()
	if ok {
		for _, where := range wheres {
			switch where.GetColumn() {
			case fql.QueryColumnName:
				switch where.GetOperator().GetType() {
				case fql.OperatorTypeEQ:
					q.Target = where.GetOperand()
				default:
					return nil, fmt.Errorf(errorStoreInvalidQuery, fq.String())
				}
			case fql.QueryColumnTimestamp:
				ts, err := strconv.ParseInt(where.GetOperand(), 10, 64)
				if err != nil {
					return nil, err
				}
				switch where.GetOperator().GetType() {
				case fql.OperatorTypeGT, fql.OperatorTypeGE:
					t := time.Unix(ts, 0)
					q.From = &t
				case fql.OperatorTypeLT, fql.OperatorTypeLE:
					t := time.Unix(ts, 0)
					q.Until = &t
				default:
					return nil, fmt.Errorf(errorStoreInvalidQuery, fq.String())
				}
			default:
				return nil, fmt.Errorf(errorStoreInvalidQuery, fq.String())
			}
		}
	}

	return q, nil
}

// String returns a string description of the instance
func (q *Query) String() string {
	target := ""
	if 0 < len(q.Target) {
		target = q.Target
	}

	from := int64(0)
	if q.From != nil {
		from = q.From.Unix()
	}

	until := int64(0)
	if q.Until != nil {
		until = q.Until.Unix()
	}

	return fmt.Sprintf("%s FROM %s WHERE %s == %s AND %s >= %d AND %s <= %d",
		fql.QuerySelectString,
		fql.QueryTargetMetrics,
		fql.QueryColumnName,
		target,
		fql.QueryColumnTimestamp,
		from,
		fql.QueryColumnTimestamp,
		until)
}
