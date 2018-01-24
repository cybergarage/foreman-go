// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// RegisterMetric represents a metric for the register.
type RegisterMetric struct {
	*Metric
}

// NewRegisterMetric returns a new metric.
func NewRegisterMetric() *RegisterMetric {
	m := &RegisterMetric{
		Metric:    NewMetric(),
		listeners: make([]RegisterMetricListener, 0),
	}
	return m
}

// NewRegisterMetricWithName returns a new register metric with the specified name.
func NewRegisterMetricWithName(name string) *RegisterMetric {
	m := NewRegisterMetric()
	m.Name = name
	return m
}

// String returns a string description of the instance
func (rm *RegisterMetric) String() string {
	return rm.Metric.String()
}
