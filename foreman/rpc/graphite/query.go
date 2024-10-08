// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphite

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/go-graphite/net/graphite"
)

// NewGraphiteQueryWithMetricQuery returns a new query from metric.Query.
func NewGraphiteQueryWithMetricQuery(mq *metric.Query) *graphite.Query {
	gq := graphite.NewQuery()
	gq.Target = mq.Target
	gq.From = mq.From
	gq.Until = mq.Until
	return gq
}

// NewMetricQueryWithGraphiteQuery returns a new query from graphite.Query.
func NewMetricQueryWithGraphiteQuery(gq *graphite.Query) *metric.Query {
	mq := metric.NewDataQuery()
	mq.Target = gq.Target
	mq.From = gq.From
	mq.Until = gq.Until
	return mq
}
