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

type QuerySourceType int

const (
	QuerySourceUnknownType QuerySourceType = iota
	QuerySourceMetricType
	QuerySourceDataType
)

// Query represents a query for the metric store.
type Query struct {
	Source   QuerySourceType
	Target   string
	From     *time.Time
	Until    *time.Time
	Interval time.Duration
}

// NewQueryWithSource returns a new query of the specified type.
func NewQueryWithSource(srcType QuerySourceType) *Query {
	q := &Query{
		Source:   srcType,
		Target:   "",
		From:     nil,
		Until:    nil,
		Interval: 0,
	}
	return q
}

// NewMetricQuery returns a new metric query.
func NewMetricQuery() *Query {
	return NewQueryWithSource(QuerySourceMetricType)
}

// NewDataQuery returns a new metric query.
func NewDataQuery() *Query {
	return NewQueryWithSource(QuerySourceDataType)
}

// NewQueryWithQuery returns a new query of the specified query.
func NewQueryWithQuery(fq fql.Query) (*Query, error) {
	if fq.GetType() != fql.QueryTypeSelect {
		return nil, fmt.Errorf(errorStoreInvalidQuery, fq.String())
	}

	var q *Query = nil
	if fq.HasOnlyColumn(fql.QueryColumnId) {
		q = NewMetricQuery()
	} else {
		q = NewDataQuery()
	}

	// Where

	wheres, ok := fq.GetConditions()
	if ok {
		for _, where := range wheres {
			switch where.GetColumn() {
			case fql.QueryColumnId:
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
	return fmt.Sprintf("%s [%s - %s]", q.Target, q.From.String(), q.Until.String())
}
