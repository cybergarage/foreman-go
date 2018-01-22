// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphite

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/go-graphite/net/graphite"
)

// NewResultSetWithGraphiteMetrics returns a new resultset from graphite.Metric.
func NewResultSetWithGraphiteMetrics(gms []*graphite.Metric) (metric.ResultSet, error) {
	rs := metric.NewResultSet()
	for _, gm := range gms {
		dps, err := NewDataPointsWithGraphiteMetric(gm)
		if err != nil {
			return nil, err
		}
		err = rs.AddDataPoints(dps)
		if err != nil {
			return nil, err
		}
	}
	return rs, nil
}
