// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// #include <foreman/foreman-c.h>
import "C"

import "time"

// Store represents an abstract interface of metric store for Foreman.
type Store interface {
	SetRetentionInterval(value time.Duration) error
	GetRetentionInterval() (time.Duration, error)

	Open() error
	Close() error

	AddMetric(m *Metric) error
	Query(q *Query) (ResultSet, error)

	String() string
}
