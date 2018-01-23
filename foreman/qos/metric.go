// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/kb"
	"github.com/cybergarage/foreman-go/foreman/metric"
)

// Metric represents an metric instance for Knowledge base.
type Metric struct {
	kb.Variable
	Name   string
	entity *metric.Metric
}

// NewMetricWithName returns a new metric instance with the specified name.
func NewMetricWithName(name string) *Metric {
	m := &Metric{
		Name:   name,
		entity: nil,
	}
	return m
}

// HasEntity returns whether the metirc has the entity
func (m *Metric) HasEntity() bool {
	if m.entity == nil {
		return false
	}
	return true
}

// SetEntity sets a entity for the QoS metric.
func (m *Metric) SetEntity(entity *metric.Metric) {
	m.entity = entity
}

// GetName is an interface method of kb.Variable
func (m *Metric) GetName() string {
	return m.Name
}

// GetValue is an interface method of kb.Variable
func (m *Metric) GetValue() (interface{}, error) {
	if m.entity == nil {
		return nil, fmt.Errorf(errorNullEntity, m.Name)
	}
	return m.entity.Value, nil
}
