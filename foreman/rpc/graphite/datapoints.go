// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphite

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/go-graphite/net/graphite"
)

// NewDataPointsWithGraphiteMetric returns a new data points from the specified graphite.Metric.
func NewDataPointsWithGraphiteMetric(gm *graphite.Metric) (*metric.DataPoints, error) {
	dpSize := len(gm.DataPoints)
	dps := metric.NewDataPoints(dpSize)
	dps.Name = gm.Name
	for n, gdp := range gm.DataPoints {
		var err error
		dps.Values[n], err = NewDataPointWithGraphiteDataPoint(gdp)
		if err != nil {
			return nil, err
		}
	}
	return dps, nil
}
