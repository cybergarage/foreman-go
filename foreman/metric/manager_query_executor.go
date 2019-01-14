// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

import (
	"math"
	"strconv"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

// ExecuteQuery must return the result as a standard array or map.
func (mgr *Manager) ExecuteQuery(q fql.Query) (interface{}, *errors.Error) {
	switch q.GetType() {
	case fql.QueryTypeInsert:
		return mgr.executeInsertQuery(q)
	case fql.QueryTypeSelect:
		return mgr.executeSelectQuery(q)
	case fql.QueryTypeAnalyze:
		return mgr.executeAnalyzeQuery(q)
	}

	return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
}

func (mgr *Manager) executeInsertQuery(q fql.Query) (interface{}, *errors.Error) {
	values, ok := q.GetValues()
	if !ok || len(values) < 3 {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidValues)
	}

	name := values[0].String()

	strValue := values[1].String()
	value, valueErr := strconv.ParseFloat(strValue, 64)
	if valueErr != nil {
		if (strValue == "None") || (strValue == "NULL") || (strValue == "nil") {
			value = math.NaN()
			valueErr = nil
		}
	}

	ts, tsErr := fql.AbsouleteTimeStringToTime(values[2].String())

	if len(name) < 0 || valueErr != nil || tsErr != nil {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidValues)
	}

	m := NewMetricWithName(name)
	m.Value = value
	m.Timestamp = *ts

	err := mgr.AddMetric(m)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return nil, nil
}

func (mgr *Manager) executeSearchMetricsQuery(q *Query) (interface{}, *errors.Error) {
	rs, err := mgr.Query(q)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	// See : Graphite HTTP API (http://graphite-api.readthedocs.io/en/latest/api.html)

	msArray := []string{}

	ms := rs.GetFirstMetrics()
	for ms != nil {
		msArray = append(msArray, ms.Name)
		ms = rs.GetNextMetrics()
	}

	return msArray, nil
}

func (mgr *Manager) executeSelectMetricsDataQuery(q *Query) (interface{}, *errors.Error) {
	rs, err := mgr.Query(q)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	// See : Graphite Render URL API (http://graphite.readthedocs.io/en/latest/render_api.html)

	rootContainer := []interface{}{}

	ms := rs.GetFirstMetrics()
	for ms != nil {
		msMap := map[string]interface{}{}
		msMap[json.GraphiteResponseTarget] = ms.Name

		dps := []interface{}{}
		for _, v := range ms.Values {
			if v == nil {
				continue
			}
			if math.IsNaN(v.Value) {
				continue
			}
			dp := []interface{}{}
			dp = append(dp, v.Value)
			dp = append(dp, v.Timestamp.Unix())
			dps = append(dps, dp)
		}
		msMap[json.GraphiteResponseDatapoints] = dps

		rootContainer = append(rootContainer, msMap)

		ms = rs.GetNextMetrics()
	}

	return rootContainer, nil
}

func (mgr *Manager) executeSelectQuery(fq fql.Query) (interface{}, *errors.Error) {
	q, err := NewQueryWithQuery(fq)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	switch q.Source {
	case QuerySourceTypeMetric:
		return mgr.executeSearchMetricsQuery(q)
	case QuerySourceDataType:
		return mgr.executeSelectMetricsDataQuery(q)
	}

	return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
}

func (mgr *Manager) executeAnalyzeQuery(fq fql.Query) (interface{}, *errors.Error) {
	q, err := NewQueryWithQuery(fq)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	rs, err := mgr.Query(q)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	rootContainer := map[string]interface{}{}

	ms := rs.GetFirstMetrics()
	for ms != nil {
		results := []float64{}
		for n, v := range ms.Values {
			switch n {
			case 0: // Correlation
				results = append(results, v.Value)
			case 1: // Max Value and Timestamp
				results = append(results, v.Value)
				results = append(results, float64(v.Timestamp.Unix()))
			case 2: // Min Value and Timestamp
				results = append(results, v.Value)
				results = append(results, float64(v.Timestamp.Unix()))
			}
		}
		rootContainer[ms.Name] = results

		ms = rs.GetNextMetrics()
	}

	return rootContainer, nil
}
