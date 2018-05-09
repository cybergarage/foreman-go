// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphite

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/go-graphite/net/graphite"
)

// NewResultSetWithGraphiteMetrics returns a resultset from an arrray of graphite metrics.
func NewResultSetWithGraphiteMetrics(gms []*graphite.Metrics) (metric.ResultSet, error) {
	rs := metric.NewResultSet()
	for _, gm := range gms {
		ms, err := NewDataPointsWithGraphiteMetric(gm)
		if err != nil {
			return nil, err
		}
		err = rs.AddMetrics(ms)
		if err != nil {
			return nil, err
		}
	}
	return rs, nil
}

// NewGraphiteMetricsWithResultSet returns an arrray of graphite metrics from a resultset.
func NewGraphiteMetricsWithResultSet(rs metric.ResultSet) ([]*graphite.Metrics, error) {
	mCount := rs.GetMetricsCount()
	gms := make([]*graphite.Metrics, mCount)

	ms := rs.GetFirstMetrics()
	for n := 0; n < mCount; n++ {
		gms[n] = graphite.NewMetrics()
		if ms == nil {
			break
		}
		gms[n].Name = ms.Name
		dpCount := len(ms.Values)
		gms[n].DataPoints = graphite.NewDataPoints(dpCount)
		for i := 0; i < dpCount; i++ {
			dp := graphite.NewDataPoint()
			dp.Timestamp = ms.Values[i].Timestamp
			dp.Value = ms.Values[i].Value
			gms[n].DataPoints[i] = dp
		}

		ms = rs.GetNextMetrics()
	}

	return gms, nil
}
