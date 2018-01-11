// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"github.com/cybergarage/foreman-go/foreman/kb"
	"github.com/cybergarage/foreman-go/foreman/metric"
)

// Metric represents an metric instance for Knowledge base.
type Metric struct {
	kb.Variable
	*metric.Metric
}

// NewMetric returns a new metric instance.
func NewMetric() *Metric {
	m := &Metric{
		Metric: metric.NewMetric(),
	}
	return m
}

// NewMetricWithName returns a new metric instance with the specified name.
func NewMetricWithName(name string) *Metric {
	m := NewMetric()
	m.Name = name
	return m
}

// GetName is an interface method of kb.Variable
func (m *Metric) GetName() string {
	return m.Name
}

// GetValue is an interface method of kb.Variable
func (m *Metric) GetValue() interface{} {
	return m.Value
}
