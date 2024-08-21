// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/foreman-go/foreman/rpc/graphite"

	go_graphite "github.com/cybergarage/go-graphite/net/graphite"
)

const (
	graphitePrefix      = "GRAPHITE"
	graphiteInsertQuery = "INSERT"
	graphiteFindQuery   = "FIND"
	graphiteRenderQuery = "RENDER"
)

// newMetricFromGraphiteMetric returns a metric from the specified metrics datapoint of Graphite
func newMetricFromGraphiteMetric(gm *go_graphite.Metrics, dp *go_graphite.DataPoint) *metric.Metric {
	fm := metric.NewMetric()
	fm.Name = gm.Name
	fm.Timestamp = dp.Timestamp
	fm.Value = dp.Value
	return fm
}

// InsertMetricsRequestReceived is a listener for Graphite Carbon
func (server *Server) InsertMetricsRequestReceived(gms []*go_graphite.Metrics, err error) {
	// Ignore error requests
	if err != nil {
		logging.Error("%s %s", graphitePrefix, err.Error())
		return
	}

	if gms == nil {
		return
	}

	if len(gms) == 0 {
		return
	}

	ms := make([]*metric.Metric, 0)
	for _, gm := range gms {
		for _, dp := range gm.DataPoints {
			m := metric.NewMetric()
			m.Name = gm.Name
			m.Timestamp = dp.Timestamp
			m.Value = dp.Value
			ms = append(ms, m)
		}
	}

	for _, m := range ms {
		err = server.metricMgr.AddMetricWithoutNotification(m)
		if err != nil {
			logging.Error("%s %s %s %s %f (%s)", graphitePrefix, graphiteInsertQuery, m.Name, m.Timestamp.String(), m.Value, err.Error())
		}
	}

	for _, m := range ms {
		err = server.metricMgr.NotifyMetric(m)
		if err != nil {
			logging.Error("%s %s %s %s %f (%s)", graphitePrefix, graphiteInsertQuery, m.Name, m.Timestamp.String(), m.Value, err.Error())
		}
	}
}

// FindMetricsRequestReceived is a listener for Graphite Render
func (server *Server) FindMetricsRequestReceived(gq *go_graphite.Query, err error) ([]*go_graphite.Metrics, error) {
	logging.Trace("%s %s %s", graphitePrefix, graphiteFindQuery, gq.Target)

	// Ignore error requests
	if err != nil {
		logging.Error("%s %s %s (%s)", graphitePrefix, graphiteFindQuery, gq.Target, err.Error())
		return nil, nil
	}

	// graphite.Query to foreman.Query
	fq := metric.NewMetricQuery()
	fq.Target = gq.Target

	rs, err := server.metricMgr.Query(fq)
	if err != nil {
		logging.Error("%s %s %s (%s)", graphitePrefix, graphiteFindQuery, gq.Target, err.Error())
		return nil, err
	}

	mCount := rs.GetMetricsCount()
	m := make([]*go_graphite.Metrics, mCount)

	ms := rs.GetFirstMetrics()
	for n := 0; n < mCount; n++ {
		m[n] = go_graphite.NewMetrics()
		if ms == nil {
			break
		}
		m[n].Name = ms.Name
		ms = rs.GetNextMetrics()
	}

	return m, nil
}

// queryMetricsRequest requests a query into a node
func (server *Server) queryMetricsRequest(node Node, gq *go_graphite.Query) ([]*go_graphite.Metrics, error) {
	mq := graphite.NewMetricQueryWithGraphiteQuery(gq)

	var rs metric.ResultSet
	var err error

	if NodeEqual(node, server) { // For local node
		rs, err = server.metricMgr.Query(mq)
	} else { // For remote node
		gc := NewClientWithNode(node)
		rs, err = gc.GetMetrics(mq)
	}

	if err != nil {
		logging.Error("%s %s %s (%s)", graphitePrefix, graphiteRenderQuery, gq.Target, err.Error())
		return nil, err
	}

	gms, err := graphite.NewGraphiteMetricsWithResultSet(rs)
	if err != nil {
		logging.Error("%s %s %s (%s)", graphitePrefix, graphiteRenderQuery, gq.Target, err.Error())
		return nil, err
	}

	return gms, nil
}

// queryFederatedMetricsRequest request a query into appropriate nodes
func (server *Server) queryFederatedMetricsRequest(gq *go_graphite.Query) ([]*go_graphite.Metrics, error) {
	// Is a regular expression query ?

	re := NewRegexp()
	err := re.CompileGraphite(gq.Target)
	if err != nil {
		return server.queryMetricsRequest(server, gq)
	}

	// Is a regular expression query ?

	hasMatchNodes := false
	gmsAll := make([]*go_graphite.Metrics, 0)

	for _, node := range server.GetAllClusterNodes() {
		if !re.MatchNode(node) {
			continue
		}

		expandTarget, ok := re.ExpandNode(node)
		if !ok {
			continue
		}

		egq := go_graphite.NewQueryWithQuery(gq)
		egq.Target = expandTarget
		gms, err := server.queryMetricsRequest(node, egq)
		if (err != nil) || (len(gms) <= 0) {
			continue
		}
		gmsAll = append(gmsAll, gms...)
	}

	// Has any match nodes for the regular expression query ?

	if hasMatchNodes {
		return gmsAll, nil
	}

	// Handle in the local node

	return server.queryMetricsRequest(server, gq)
}

// QueryMetricsRequestReceived is a listener for Graphite Render
func (server *Server) QueryMetricsRequestReceived(gq *go_graphite.Query, err error) ([]*go_graphite.Metrics, error) {
	logging.Trace("%s %s %s %s %s", graphitePrefix, graphiteRenderQuery, gq.Target, gq.From.String(), gq.Until.String())

	// Ignore error requests
	if err != nil {
		logging.Error("%s %s %s (%s)", graphitePrefix, graphiteRenderQuery, gq.Target, err.Error())
		return nil, nil
	}

	gms, err := server.queryFederatedMetricsRequest(gq)
	if err != nil {
		logging.Error("%s %s %s (%s)", graphitePrefix, graphiteRenderQuery, gq.Target, err.Error())
		return nil, err
	}

	return gms, nil
}
