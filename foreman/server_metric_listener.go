// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/foreman-go/foreman/qos"
)

// updateQoSMetric updates a metric in QoSs
func (server *Server) updateQoSMetricData(rm *metric.RegisterMetric) {
	name, err := rm.GetName()
	if err != nil {
		return
	}

	ver, ok := server.qosMgr.GetVariable(name)
	if !ok {
		return
	}

	value, err := rm.GetValue()
	if err != nil {
		return
	}

	ver.SetValue(value)
}

// updateRegisterMetricListeners updates all listeners
func (server *Server) updateRegisterMetricListeners(rm *metric.RegisterMetric) {
	name, err := rm.GetName()
	if err != nil {
		return
	}

	q := qos.NewQuery()
	q.Target = name

	rules, err := server.qosMgr.FindRelatedRules(q)
	if err != nil {
		return
	}

	for _, rule := range rules {
		rm.AddListener(rule)
		rule.SetListener(server)
	}
}

// RegisterMetricAdded is a listener for metric.Register
func (server *Server) RegisterMetricAdded(rm *metric.RegisterMetric) {
	// FIXME : Optimize the listener
	server.updateQoSMetricData(rm)
	server.updateRegisterMetricListeners(rm)
}

// RegisterMetricUpdated is a listener for metric.Register
func (server *Server) RegisterMetricUpdated(rm *metric.RegisterMetric) {
	// FIXME : QoS which is added after the metrics doesn't listen
	server.updateQoSMetricData(rm)
}
