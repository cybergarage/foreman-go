// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/foreman-go/foreman/qos"
)

// RegisterMetricAdded is a listener for metric.Register
func (server *Server) RegisterMetricAdded(rm *metric.RegisterMetric) {
	q := qos.NewQuery()
	q.Target = rm.GetName()

	formula, err := server.qosMgr.FindRelatedFormulas(q)
	if err != nil {
		return
	}

}

// RegisterMetricUpdated is a listener for metric.Register
func (server *Server) RegisterMetricUpdated(rm *metric.RegisterMetric) {
}
