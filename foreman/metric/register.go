// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

import (
	"github.com/cybergarage/foreman-go/foreman/register"
)

// RegisterListener represents a listener for metric register.
type RegisterListener interface {
	// RegisterMetricAdded is called when a new metric is added
	RegisterMetricAdded(*RegisterMetric)
	// RegisterMetricUpdated is called when a metric which is already added is updated
	RegisterMetricUpdated(*RegisterMetric)
}

// Register represents an register store for metric.
type Register struct {
	StoreListener
	register.Store
	Listener RegisterListener
}

// NewRegisterWithStore returns a new store instance.
func NewRegisterWithStore(store register.Store) *Register {
	reg := &Register{
		Store:    store,
		Listener: nil,
	}
	return reg
}

// SetRegisterListener sets a listener.
func (rs *Register) SetRegisterListener(listener RegisterListener) {
	rs.Listener = listener
}

// GetMetric gets the specified metric.
func (rs *Register) GetMetric(key string) (*RegisterMetric, bool) {
	obj, ok := rs.GetObject(key)
	if !ok {
		return nil, false
	}

	rm := NewRegisterMetricWithObject(obj)

	return rm, ok
}

// UpdateMetricWithoutNotification updates the specified metric without notification.
func (rs *Register) UpdateMetricWithoutNotification(m *Metric) error {
	key := m.GetName()

	obj, _ := rs.GetObject(key)

	// Add new metric

	if obj == nil {
		rm := NewRegisterMetricWithMetric(m)
		err := rs.SetObject(rm)
		if err != nil {
			return err
		}

		return nil
	}

	// Update metric

	rm := NewRegisterMetricWithObject(obj)

	err := rm.SetValue(m.GetValue())
	if err != nil {
		return err
	}

	err = rs.SetObject(rm.Object)
	if err != nil {
		return err
	}

	return nil
}

// NotifyMetric notifies the metric to the listener.
func (rs *Register) NotifyMetric(m *Metric) error {
	key := m.GetName()
	obj, _ := rs.GetObject(key)

	if obj == nil {
		return nil
	}

	rm := NewRegisterMetricWithObject(obj)

	if rs.Listener != nil {
		rs.Listener.RegisterMetricUpdated(rm)
	}

	return nil
}

// UpdateMetric updates the specified metric with notification.
func (rs *Register) UpdateMetric(m *Metric) error {
	err := rs.UpdateMetricWithoutNotification(m)
	if err != nil {
		return err
	}
	return rs.NotifyMetric(m)
}
