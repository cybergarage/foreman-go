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

// InsertMetricsRequestReceived is a listener for Graphite Carbon
func (server *Server) InsertMetricsRequestReceived(gm *go_graphite.Metrics, err error) {
	// Ignore error requests
	if err != nil {
		logging.Error("[GRAPHITE:BAD] %s", err.Error())
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
			logging.Error("%s %s %s %s %f", graphitePrefix, graphiteInsertQuery, fm.Name, fm.Timestamp.String(), fm.Value)
			continue
		}

		logging.Trace("%s %s %s %s %f", graphitePrefix, graphiteInsertQuery, fm.Name, fm.Timestamp.String(), fm.Value)
	}
}

// FindMetricsRequestReceived is a listener for Graphite Render
func (server *Server) FindMetricsRequestReceived(gq *go_graphite.Query, err error) ([]*go_graphite.Metrics, error) {
	// Ignore error requests
	if err != nil {
		logging.Error("%s %s %s", graphitePrefix, graphiteFindQuery, gq.Target)
		return nil, nil
	}

	// graphite.Query to foreman.Query
	fq := metric.NewMetricQuery()
	fq.Target = gq.Target

	rs, err := server.metricMgr.Query(fq)
	if err != nil {
		logging.Error("%s %s %s", graphitePrefix, graphiteFindQuery, gq.Target)
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

	logging.Info("%s %s %s", graphitePrefix, graphiteFindQuery, gq.Target)

	return m, nil
}

// queryMetricsRequest request a query into a node
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
		logging.Error("%s %s %s", graphitePrefix, graphiteRenderQuery, gq.Target)
		return nil, err
	}

	gms, err := graphite.NewGraphiteMetricsWithResultSet(rs)
	if err != nil {
		logging.Error("%s %s %s", graphitePrefix, graphiteRenderQuery, gq.Target)
		return nil, err
	}

	return gms, nil
}

// queryFederatedMetricsRequest request a query into appropriate nodes
func (server *Server) queryFederatedMetricsRequest(gq *go_graphite.Query) ([]*go_graphite.Metrics, error) {
	re := NewRegexp()
	err := re.CompileGraphite(gq.Target)
	if err != nil {
		return nil, err
	}

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

	return gmsAll, nil
}

// QueryMetricsRequestReceived is a listener for Graphite Render
func (server *Server) QueryMetricsRequestReceived(gq *go_graphite.Query, err error) ([]*go_graphite.Metrics, error) {
	// Ignore error requests
	if err != nil {
		logging.Error("%s %s %s", graphitePrefix, graphiteRenderQuery, gq.Target)
		return nil, nil
	}

	gms, err := server.queryFederatedMetricsRequest(gq)
	if err != nil {
		logging.Error("%s %s %s", graphitePrefix, graphiteRenderQuery, gq.Target)
		return nil, err
	}

	logging.Info("%s %s %s %s %s", graphitePrefix, graphiteRenderQuery, gq.Target, gq.From.String(), gq.Until.String())

	return gms, nil
}
