// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphite

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/go-graphite/net/graphite"
)

// NewDataPointWithGraphiteDataPoint returns a new data point from the specified Graphite data point.
func NewDataPointWithGraphiteDataPoint(gdp *graphite.DataPoint) (*metric.DataPoint, error) {
	dp := metric.NewDataPoint()
	dp.Timestamp = gdp.Timestamp
	dp.Value = gdp.Value
	return dp, nil
}
