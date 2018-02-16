// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

import (
	"strconv"

	"github.com/cybergarage/foreman-go/foreman/register"
)

// RegisterMetricListener represents a listener interface for the register metric.
type RegisterMetricListener interface {
	register.ObjectListener
}

// RegisterMetric represents a metric for the register.
type RegisterMetric struct {
	register.Object
}

// NewRegisterMetric returns a new metric.
func NewRegisterMetric() *RegisterMetric {
	m := &RegisterMetric{
		Object: register.NewObject(),
	}
	return m
}

// NewRegisterMetricWithMetric returns a new register metric with the metric.
func NewRegisterMetricWithMetric(m *Metric) *RegisterMetric {
	rm := NewRegisterMetric()
	rm.SetName(m.GetName())
	rm.SetValue(m.GetValue())
	return rm
}

// NewRegisterMetricWithObject returns a new register metric with the register object.
func NewRegisterMetricWithObject(obj register.Object) *RegisterMetric {
	rm := &RegisterMetric{
		Object: obj,
	}
	return rm
}

// SetValue sets the specified value.
func (rm *RegisterMetric) SetValue(value float64) error {
	return rm.SetData(strconv.FormatFloat(value, 'E', -1, 64))
}

// GetValue returns the stored value.
func (rm *RegisterMetric) GetValue() (float64, error) {
	strValue, err := rm.GetData()
	if err != nil {
		return 0.0, err
	}
	value, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		return 0.0, err
	}
	return value, nil
}

// String returns a string description of the instance
func (rm *RegisterMetric) String() string {
	return rm.Object.String()
}
