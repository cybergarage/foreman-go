// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
)

// RegisterMetricAdded is a listener for metric.Register
func (server *Server) RegisterMetricAdded(rm *metric.RegisterMetric) {
	name, err := rm.GetName()
	if err != nil {
		return
	}
	value, err := rm.GetValue()
	if err != nil {
		return
	}
	server.qosMgr.UpdateVariableValue(name, value)
}

// RegisterMetricUpdated is a listener for metric.Register
func (server *Server) RegisterMetricUpdated(rm *metric.RegisterMetric) {
	name, err := rm.GetName()
	if err != nil {
		return
	}
	value, err := rm.GetValue()
	if err != nil {
		return
	}
	server.qosMgr.UpdateVariableValue(name, value)
}
