// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

import (
	"github.com/cybergarage/foreman-go/foreman/register"
)

// RegisterMetricListener represents a listener interface for the register metric.
type RegisterMetricListener interface {
	register.ObjectListener
}

// RegisterMetric represents a metric for the register.
type RegisterMetric struct {
	*Metric
	*register.BaseObject
}

// NewRegisterMetric returns a new metric.
func NewRegisterMetric() *RegisterMetric {
	m := &RegisterMetric{
		Metric:     NewMetric(),
		BaseObject: register.NewObject(),
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
