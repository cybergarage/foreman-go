// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -llua -lpython
import "C"

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

// Store represents a metric store for Foreman.
type cgoStore struct {
	cStore   unsafe.Pointer
	listener StoreListener
}

// Open initializes the store.
func (store *cgoStore) Open() error {
	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	if !C.foreman_metric_store_open(store.cStore) {
		return fmt.Errorf(errorStoreCouldNotOpen, store)
	}

	return nil
}

// Close closes the store.
func (store *cgoStore) Close() error {
	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	if !C.foreman_metric_store_close(store.cStore) {
		return fmt.Errorf(errorStoreCouldNotClose, store)
	}

	return nil
}

// Clear remove all inserted data.
func (store *cgoStore) Clear() error {
	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	if !C.foreman_metric_store_clear(store.cStore) {
		return fmt.Errorf(errorStoreCouldNotClose, store)
	}

	return nil
}

// SetListener sets a listener.
func (store *cgoStore) SetListener(listener StoreListener) error {
	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	store.listener = listener

	return nil
}

// SetRetentionInterval sets the retention duration.
func (store *cgoStore) SetRetentionInterval(value time.Duration) error {
	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	C.foreman_metric_store_setretentioninterval(store.cStore, C.time_t(value.Seconds()))

	return nil
}

// GetRetentionInterval returns the retention duration.
func (store *cgoStore) GetRetentionInterval() (time.Duration, error) {
	if store.cStore == nil {
		return 0, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	durationSec := C.foreman_metric_store_getretentioninterval(store.cStore)
	duration := time.Second * time.Duration(durationSec)

	return duration, nil
}

// addMetric adds a new metric.
func (store *cgoStore) addMetric(m *Metric) error {
	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cm, err := m.CMetric()
	if err != nil {
		return err
	}

	isSuccess := C.foreman_metric_store_addmetric(store.cStore, cm)
	if !isSuccess {
		return fmt.Errorf(errorStoreCouldNotAddMetric, m.String())
	}

	return nil
}

// AddMetric adds a new metric.
func (store *cgoStore) AddMetric(m *Metric) error {
	err := store.addMetric(m)
	if store.listener != nil {
		store.listener.MetricAdded(m, err)
	}
	return err
}

// Query gets the specified metrics.
func (store *cgoStore) Query(q *Query) (ResultSet, error) {
	if store.cStore == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	duration, err := store.GetRetentionInterval()
	if err != nil {
		return nil, err
	}
	q.Interval = duration

	cq, err := q.CQuery()
	if err != nil {
		return nil, err
	}

	crs := C.foreman_metric_resultset_new()

	if !C.foreman_metric_store_query(store.cStore, cq, crs) {
		C.foreman_metric_resultset_delete(crs)
		return nil, fmt.Errorf(errorStoreCouldNotAddMetric, q.String())
	}

	return NewResultSetWithCObject(crs), nil
}

// String returns a string description of the instance
func (store *cgoStore) String() string {
	// FIXME : Not implemented
	return fmt.Sprintf("")
}
