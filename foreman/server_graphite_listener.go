// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/go-graphite/net/graphite"
)

// MetricRequestReceived is a listener for Graphite Carbon
func (server *Server) MetricRequestReceived(gm *graphite.Metric, err error) {
	// Ignore error requests
	if err != nil {
		return
	}

	for _, dp := range gm.DataPoints {
		// graphite.Metric to foreman.Metric
		fm := metric.NewMetric()
		fm.Name = gm.Name
		fm.Timestamp = dp.Timestamp
		fm.Value = dp.Value

		err = server.metricMgr.AddMetric(fm)
		if err != nil {
			// TODO : Handle the error
		}
	}
}

// QueryRequestReceived is a listener for Graphite Render
func (server *Server) QueryRequestReceived(gq *graphite.Query, err error) ([]*graphite.Metric, error) {
	// Ignore error requests
	if err != nil {
		return nil, nil
	}

	// graphite.Query to foreman.Query
	fq := metric.NewQuery()
	fq.Target = gq.Target
	fq.From = gq.From
	fq.Until = gq.Until

	rs, err := server.metricMgr.Query(fq)
	if err != nil {
		return nil, err
	}

	mCount := rs.GetDataPointCount()
	m := make([]*graphite.Metric, mCount)

	ms := rs.GetFirstMetrics()
	for n := 0; n < mCount; n++ {
		m[n] = graphite.NewMetric()
		if ms == nil {
			break
		}
		m[n].Name = ms.Name
		dpCount := len(ms.Values)
		m[n].DataPoints = graphite.NewDataPoints(dpCount)
		for i := 0; i < dpCount; i++ {
			dp := graphite.NewDataPoint()
			dp.Timestamp = ms.Values[i].Timestamp
			dp.Value = ms.Values[i].Value
			m[n].DataPoints[i] = dp
		}

		ms = rs.GetNextMetrics()
	}

	return m, nil
}
