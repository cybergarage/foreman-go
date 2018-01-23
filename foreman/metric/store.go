// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

import "time"

// Storing represents an abstract interface of metric store
type Storing interface {
	SetStoreListener(StoreListener) error

	SetRetentionInterval(value time.Duration) error
	GetRetentionInterval() (time.Duration, error)

	Open() error
	Close() error
	Clear() error

	AddMetric(m *Metric) error
	Query(q *Query) (ResultSet, error)

	String() string
}

// Store represents an metric store.
type Store struct {
	Storing
}
