// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

import (
	"strconv"
	"time"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
)

func (mgr *Manager) executeInsertQuery(q fql.Query) (interface{}, *errors.Error) {
	values, ok := q.GetValues()
	if !ok || len(values) < 3 {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidValues)
	}

	name := values[0].String()
	value, valueErr := strconv.ParseFloat(values[1].String(), 64)
	ts, tsErr := strconv.ParseInt(values[2].String(), 10, 64)
	if len(name) < 0 || valueErr != nil || tsErr != nil {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidValues)
	}

	m := NewMetricWithName(name)
	m.Value = value
	m.Timestamp = time.Unix(ts, 0)

	err := mgr.AddMetric(m)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return nil, nil
}

func (mgr *Manager) executeSelectQuery(fq fql.Query) (interface{}, *errors.Error) {
	q, err := NewQueryWithQuery(fq)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	rs, err := mgr.Query(q)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	ms := rs.GetFirstMetrics()
	for ms != nil {
		ms = rs.GetNextMetrics()
	}

	return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
}

// ExecuteQuery must return the result as a standard array or map.
func (mgr *Manager) ExecuteQuery(q fql.Query) (interface{}, *errors.Error) {
	switch q.GetType() {
	case fql.QueryTypeInsert:
		return mgr.executeInsertQuery(q)
	case fql.QueryTypeSelect:
		return mgr.executeSelectQuery(q)
	}

	return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
}
