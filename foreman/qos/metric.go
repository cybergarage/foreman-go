// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"fmt"
	"strconv"

	"github.com/cybergarage/foreman-go/foreman/kb"
)

// Metric represents an metric instance for Knowledge base.
type Metric struct {
	kb.Variable
	Name  string
	Value float64
}

// NewMetricWithName returns a new metric instance with the specified name.
func NewMetricWithName(name string) *Metric {
	m := &Metric{
		Name:  name,
		Value: 0.0,
	}
	return m
}

// GetName is an interface method of kb.Variable
func (m *Metric) GetName() string {
	return m.Name
}

// SetValue is an interface method of kb.Variable
func (m *Metric) SetValue(obj interface{}) error {
	switch obj.(type) {
	case float64:
		value, _ := obj.(float64)
		m.Value = value
	case float32:
		value, _ := obj.(float32)
		m.Value = float64(value)
	case int:
		value, _ := obj.(int)
		m.Value = float64(value)
	case int32:
		value, _ := obj.(int32)
		m.Value = float64(value)
	case int64:
		value, _ := obj.(int64)
		m.Value = float64(value)
	case string:
		str, _ := obj.(string)
		value, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		m.Value = value
	default:
		return fmt.Errorf(errorInvalidVariable, obj)
	}
	return nil
}

// GetValue is an interface method of kb.Variable
func (m *Metric) GetValue() (interface{}, error) {
	return m.Value, nil
}

// Expression returns a expression string in the formula
func (m *Metric) Expression() string {
	return m.Name
}
