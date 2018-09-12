// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

// #include <foreman/foreman-c.h>
import "C"

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

const (
	// FIXME : Support auto vacuum
	cgoStoreVacuumInterval = 1000
)

// Store represents a metric store for Foreman.
type cgoStore struct {
	cStore   unsafe.Pointer
	listener StoreListener
	// FIXME : Support auto vacuum
	vacuumCounter uint
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

	// FIXME : Support auto vacuum
	store.vacuumCounter++
	if cgoStoreVacuumInterval < store.vacuumCounter {
		err = store.Vacuum()
		if err != nil {
			return err
		}
		store.vacuumCounter = 0
	}

	return nil
}

// Query gets the specified metrics.
func (store *cgoStore) Query(q *Query) (ResultSet, error) {
	if store.cStore == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	if q.Interval == 0 {
		duration, err := store.GetRetentionInterval()
		if err != nil {
			return nil, err
		}
		q.Interval = duration
	}

	cq, err := q.CQuery()
	if err != nil {
		return nil, err
	}

	if q.Source == QuerySourceTypeUnknown {
		return nil, fmt.Errorf(errorStoreInvalidQuery, q.String())
	}

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	crs := C.foreman_metric_resultset_new()

	executed := false
	switch q.Type {
	case QueryTypeSelect:
		switch q.Source {
		case QuerySourceTypeMetric:
			executed = bool(C.foreman_metric_store_querymetric(store.cStore, cq, crs))
		case QuerySourceDataType:
			executed = bool(C.foreman_metric_store_querydata(store.cStore, cq, crs))
		}
	case QueryTypeAnalyze:
		switch q.Source {
		case QuerySourceTypeMetric, QuerySourceDataType:
			executed = bool(C.foreman_metric_store_analyzedata(store.cStore, cq, crs, cerr))
		}
	}

	if !executed {
		C.foreman_metric_resultset_delete(crs)

		errMsg := fmt.Sprintf(errorStoreInvalidQuery, q.String())
		queryErr := errors.NewWithCObject(cerr)
		if 0 < len(queryErr.Message) {
			errMsg = fmt.Sprintf("%s (%s)", errMsg, queryErr.Message)
		}

		return nil, fmt.Errorf(errMsg)
	}

	return NewResultSetWithCObject(crs), nil
}

// Query gets the specified metrics.
func (store *cgoStore) Vacuum() error {
	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	C.foreman_metric_store_query_delete_expired_metrics(store.cStore)

	return nil
}

// String returns a string description of the instance
func (store *cgoStore) String() string {
	// FIXME : Not implemented
	return fmt.Sprintf("")
}
