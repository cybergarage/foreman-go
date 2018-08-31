// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphite

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/go-graphite/net/graphite"
)

// NewMetricsWithGraphiteMetric returns new metrics from Graphite metric.
func NewMetricsWithGraphiteMetric(gm *graphite.Metrics) []*metric.Metric {
	ms := make([]*metric.Metric, len(gm.DataPoints))
	for n, dp := range gm.DataPoints {
		ms[n] = metric.NewMetric()
		ms[n].Name = gm.Name
		ms[n].Value = dp.Value
		ms[n].Timestamp = dp.Timestamp
	}
	return ms
}

// NewGraphiteMetricsWithMetric returns a new metric from Graphite metric.
func NewGraphiteMetricsWithMetric(m *metric.Metric) (*graphite.Metrics, error) {
	gm := graphite.NewMetrics()
	gm.Name = m.Name

	dp := graphite.NewDataPoint()
	dp.Value = m.Value
	dp.Timestamp = m.Timestamp
	err := gm.AddDataPoint(dp)
	if err != nil {
		return nil, err
	}

	return gm, nil
}
