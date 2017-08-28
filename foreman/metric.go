// Copyright 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

// Metric represents a Foreman Metric.
type Metric struct {
}

// NewMetric returns a new Metric.
func NewMetric() *Metric {
	m := &Metric{}

	return m
}
