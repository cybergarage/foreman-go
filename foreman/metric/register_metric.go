// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// RegisterMetricListener represents a listener for metric register.
type RegisterMetricListener interface {
	// RegisterMetricUpdated is called when the metric is updated.
	RegisterMetricUpdated(*RegisterMetric)
}

// RegisterMetric represents a metric for the register.
type RegisterMetric struct {
	*Metric
	listeners []RegisterMetricListener
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

// HasListener checks whether the specified listener is already added
func (rm *RegisterMetric) HasListener(listener RegisterMetricListener) bool {
	for _, l := range rm.listeners {
		if l == listener {
			return true
		}
	}
	return false
}

// AddListener adds a new listener
func (rm *RegisterMetric) AddListener(listener RegisterMetricListener) bool {
	if rm.HasListener(listener) {
		return false
	}
	rm.listeners = append(rm.listeners, listener)
	return true
}

// RemoveListener removes the specified listener
func (rm *RegisterMetric) RemoveListener(listener RegisterMetricListener) bool {
	for n, l := range rm.listeners {
		if l == listener {
			lastIndex := len(rm.listeners) - 1
			rm.listeners[lastIndex], rm.listeners[n] = rm.listeners[n], rm.listeners[lastIndex]
			rm.listeners = rm.listeners[:lastIndex]
			return true
		}
	}
	return false
}

// GetListeners returns the current listeners
func (rm *RegisterMetric) GetListeners() []RegisterMetricListener {
	return rm.listeners
}

// String returns a string description of the instance
func (rm *RegisterMetric) String() string {
	return rm.Metric.String()
}
