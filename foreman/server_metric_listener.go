// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
)

// MetricAdded is a listener for metric.Register
func (server *Server) MetricAdded(m *metric.Metric) {
}

// MetricUpdated is a listener for metric.Register
func (server *Server) MetricUpdated(m *metric.Metric) {
}
