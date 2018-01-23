// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/register"
)

// Register represents an register store for metric.
type Register struct {
	*register.Store
	Listener RegisterListener
}

// NewRegister returns a new store instance.
func NewRegister() *Register {
	reg := &Register{
		Store:    register.NewStore(),
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

	rm, ok := obj.Data.(*RegisterMetric)
	if !ok {
		return nil, false
	}

	return rm, ok
}

// UpdateMetric updates the specified metric.
func (rs *Register) UpdateMetric(m *Metric) error {
	key := m.Name
	isNewMetricAdded := false

	obj, _ := rs.GetObject(key)
	if obj == nil {
		rm := NewRegisterMetric()
		rm.Name = key
		obj = register.NewObject()
		obj.Name = key
		obj.Data = rm
		isNewMetricAdded = true
	}

	rm, ok := obj.Data.(*RegisterMetric)
	if !ok {
		return fmt.Errorf(errorInvalidMetric, rm)
	}

	rm.Value = m.Value

	err := obj.UpdateVersionAndTimestamp()
	if err != nil {
		return err
	}

	if rs.Listener != nil {
		if isNewMetricAdded {
			rs.Listener.RegisterMetricAdded(rm)
		} else {
			rs.Listener.RegisterMetricUpdated(rm)
		}
	}

	return nil
}
