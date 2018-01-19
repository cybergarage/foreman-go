// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphite

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/go-graphite/net/graphite"
)

// NewMetricsWithGraphiteMetric returns new metrics from Graphite metric.
func NewMetricsWithGraphiteMetric(gm *graphite.Metric) []*metric.Metric {
	m := make([]*metric.Metric, len(gm.DataPoints))

	for n, dp := range gm.DataPoints {
		m[n] = metric.NewMetric()
		m[n].Name = gm.Name
		m[n].Value = dp.Value
		m[n].Timestamp = dp.Timestamp
	}

	return m
}

// NewGraphiteMetricsWithMetric returns a new metric from Graphite metric.
func NewGraphiteMetricsWithMetric(m *metric.Metric) (*graphite.Metric, error) {
	gm := graphite.NewMetric()
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
