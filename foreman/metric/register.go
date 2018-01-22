// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/register"
)

// RegisterListener represents a listener for metric register.
type RegisterListener interface {
	MetricUpdated(*Metric)
}

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

// SetListener sets a listener.
func (rs *Register) SetListener(listener RegisterListener) {
	rs.Listener = listener
}

// GetMetric gets the specified metric.
func (rs *Register) GetMetric(key string) (*Metric, bool) {
	obj, ok := rs.GetObject(key)
	if !ok {
		return nil, false
	}

	m, ok := obj.Data.(*Metric)
	if !ok {
		return nil, false
	}

	return m, ok
}

// UpdateMetric updates the specified metric.
func (rs *Register) UpdateMetric(m *Metric) error {
	key := m.Name

	obj, _ := rs.GetObject(key)
	if obj == nil {
		rm := NewMetric()
		rm.Name = key
		obj = register.NewObject()
		obj.Name = key
		obj.Data = rm
	}

	rm, ok := obj.Data.(*Metric)
	if !ok {
		return fmt.Errorf(errorInvalidMetric, rm)
	}

	err := obj.UpdateVersionAndTimestamp()
	if err != nil {
		return err
	}

	rm.Value = m.Value

	if rs.Listener != nil {
		rs.Listener.MetricUpdated(rm)
	}

	return nil
}
