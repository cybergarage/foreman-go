// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/foreman-go/foreman/qos"
)

// RegisterMetricAdded is a listener for metric.Register
func (server *Server) RegisterMetricAdded(rm *metric.RegisterMetric) {
	q := qos.NewQuery()
	q.Target = rm.GetName()

	// Set a new metric into the related formulas

	formulas, err := server.qosMgr.FindRelatedFormulas(q)
	if err != nil {
		return
	}

	for _, formula := range formulas {
		formula.SetMetricEntity(rm.Metric)
	}

	// Set a new metric into the related rules

	rules, err := server.qosMgr.FindRelatedRules(q)
	if err != nil {
		return
	}

	for _, rule := range rules {
		rm.AddListener(rule)
		rule.SetListener(server)
	}
}

// RegisterMetricUpdated is a listener for metric.Register
func (server *Server) RegisterMetricUpdated(rm *metric.RegisterMetric) {
	// RegisterMetricUpdated does nothing because the metric change is watched by the related rules directly.
}
