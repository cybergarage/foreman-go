// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

import (
	"fmt"
	"time"
)

// Metric represents a Foreman metric.
type Metric struct {
	Name      string
	Value     float64
	Timestamp time.Time
}

// NewMetric returns a new metric.
func NewMetric() *Metric {
	m := &Metric{}
	return m
}

// String returns a string description of the instance
func (self *Metric) String() string {
	return fmt.Sprintf("%s : %f (%d)", self.Name, self.Value, self.Timestamp.Unix())
}
