// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

// SetStoreListener sets a listener.
func (store *cgoStore) SetStoreListener(listener StoreListener) error {
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

// AddMetric adds a new metric.
func (store *cgoStore) AddMetric(m *Metric) error {
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

	if store.listener != nil {
		store.listener.StoreMetricAdded(m)
	}

	return nil
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

	if q.Source == QuerySourceTypeUnknown {
		return nil, fmt.Errorf(errorStoreInvalidQuery, q.String())
	}

	crs := C.foreman_metric_resultset_new()

	executed := false
	switch q.Source {
	case QuerySourceTypeMetric:
		executed = bool(C.foreman_metric_store_querymetric(store.cStore, cq, crs))
	case QuerySourceTypeData:
		executed = bool(C.foreman_metric_store_querydata(store.cStore, cq, crs))
	}

	if !executed {
		C.foreman_metric_resultset_delete(crs)
		return nil, fmt.Errorf(errorStoreInvalidQuery, q.String())
	}

	return NewResultSetWithCObject(crs), nil
}

// String returns a string description of the instance
func (store *cgoStore) String() string {
	// FIXME : Not implemented
	return fmt.Sprintf("")
}
