// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/go-graphite/net/graphite"
)

const (
	graphiteGoodRequestPrefix = "[GRAPHITE:OK]"
	graphiteBadRequestPrefix  = "[GRAPHITE:BAD]"
	graphiteInsertQuery       = "INSERT"
	graphiteFindQuery         = "FIND"
	graphiteRenderQuery       = "RENDER"
)

// InsertMetricRequestReceived is a listener for Graphite Carbon
func (server *Server) InsertMetricsRequestReceived(gm *graphite.Metrics, err error) {
	// Ignore error requests
	if err != nil {
		logging.Error("[GRAPHITE:BAD] : %s", err.Error())
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
			logging.Error("%s : %s %s %s %f", graphiteBadRequestPrefix, graphiteInsertQuery, fm.Name, fm.Timestamp.String(), fm.Value)
			continue
		}

		logging.Info("%s : %s %s %s %f", graphiteGoodRequestPrefix, graphiteInsertQuery, fm.Name, fm.Timestamp.String(), fm.Value)
	}
}

// FindMetricsRequestReceived is a listener for Graphite Render
func (server *Server) FindMetricsRequestReceived(gq *graphite.Query, err error) ([]*graphite.Metrics, error) {
	// Ignore error requests
	if err != nil {
		logging.Error("%s : %s %s", graphiteBadRequestPrefix, graphiteFindQuery, gq.Target)
		return nil, nil
	}

	// graphite.Query to foreman.Query
	fq := metric.NewMetricQuery()
	fq.Target = gq.Target

	rs, err := server.metricMgr.Query(fq)
	if err != nil {
		logging.Error("%s : %s %s", graphiteBadRequestPrefix, graphiteFindQuery, gq.Target)
		return nil, err
	}

	mCount := rs.GetMetricsCount()
	m := make([]*graphite.Metrics, mCount)

	ms := rs.GetFirstMetrics()
	for n := 0; n < mCount; n++ {
		m[n] = graphite.NewMetrics()
		if ms == nil {
			break
		}
		m[n].Name = ms.Name
		ms = rs.GetNextMetrics()
	}

	logging.Info("%s : %s %s", graphiteGoodRequestPrefix, graphiteFindQuery, gq.Target)

	return m, nil
}

// QueryMetricsRequestReceived is a listener for Graphite Render
func (server *Server) QueryMetricsRequestReceived(gq *graphite.Query, err error) ([]*graphite.Metrics, error) {
	// Ignore error requests
	if err != nil {
		logging.Error("%s : %s %s", graphiteBadRequestPrefix, graphiteRenderQuery, gq.Target)
		return nil, nil
	}

	// graphite.Query to foreman.Query
	fq := metric.NewDataQuery()
	fq.Target = gq.Target
	fq.From = gq.From
	fq.Until = gq.Until

	rs, err := server.metricMgr.Query(fq)
	if err != nil {
		logging.Error("%s : %s %s", graphiteBadRequestPrefix, graphiteRenderQuery, gq.Target)
		return nil, err
	}

	mCount := rs.GetMetricsCount()
	m := make([]*graphite.Metrics, mCount)

	ms := rs.GetFirstMetrics()
	for n := 0; n < mCount; n++ {
		m[n] = graphite.NewMetrics()
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

	logging.Info("%s : %s %s", graphiteGoodRequestPrefix, graphiteRenderQuery, gq.Target)

	return m, nil
}
