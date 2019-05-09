// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

// #include <foreman/foreman-c.h>
import "C"

import (
	"fmt"
	"sync"
	"time"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/register"
)

const (
	// FIXME : Support auto vacuum
	cgoStoreVacuumInterval = 1000
)

// Store represents a metric store for Foreman.
type cgoStore struct {
	cStore   unsafe.Pointer
	listener StoreListener
	*Register
	*sync.Mutex
	// FIXME : Support auto vacuum
	vacuumCounter uint
}

// SetRegisterStore sets a register store.
func (store *cgoStore) SetRegisterStore(regStore register.Store) error {
	store.Register = NewRegisterWithStore(regStore)
	return nil
}

// SetRetentionInterval sets the specified retention duration.
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

// SetRetentionPeriod sets the specified retention period.
func (store *cgoStore) SetRetentionPeriod(value time.Duration) error {
	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	C.foreman_metric_store_setretentionperiod(store.cStore, C.time_t(value.Seconds()))

	return nil
}

// GetRetentionPeriod returns the retention period.
func (store *cgoStore) GetRetentionPeriod() (time.Duration, error) {
	if store.cStore == nil {
		return 0, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	durationSec := C.foreman_metric_store_getretentionperiod(store.cStore)
	duration := time.Second * time.Duration(durationSec)

	return duration, nil
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

	store.Lock()
	defer store.Unlock()

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

// AddMetricWithoutNotification adds a new metric without notification.
func (store *cgoStore) AddMetricWithoutNotification(m *Metric) error {

	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cm, err := m.CObject()
	if err != nil {
		return err
	}
	defer C.foreman_metric_delete(cm)

	var lastErr error

	store.Lock()

	isSuccess := C.foreman_metric_store_addmetric(store.cStore, cm)
	if !isSuccess {
		lastErr = fmt.Errorf(errorStoreCouldNotAddMetric, m.String())
	}

	store.Unlock()

	// Update register

	if store.Register != nil {
		store.Register.UpdateMetricWithoutNotification(m)
	}

	// FIXME : Support auto vacuum

	ri, err := store.GetRetentionPeriod()
	if (err == nil) && (0 < ri) {
		store.vacuumCounter++
		if cgoStoreVacuumInterval < store.vacuumCounter {
			err = store.Vacuum()
			if err == nil {
				store.vacuumCounter = 0
			} else {
				lastErr = err
			}
		}
	}

	return lastErr
}

// NotifyMetric notifies a new metric to the listeners.
func (store *cgoStore) NotifyMetric(m *Metric) error {
	// Notify to register

	if store.Register != nil {
		store.Register.NotifyMetric(m)
	}

	// Notify to listener

	if store.listener == nil {
		return nil
	}
	return store.listener.StoreMetricAdded(m)
}

// AddMetric adds a new metric.
func (store *cgoStore) AddMetric(m *Metric) error {
	err := store.AddMetricWithoutNotification(m)
	if err != nil {
		return err
	}
	return store.NotifyMetric(m)
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

	cq, err := q.CObject()
	if err != nil {
		return nil, err
	}
	defer C.foreman_metric_query_delete(cq)

	if q.Source == QuerySourceTypeUnknown {
		return nil, fmt.Errorf(errorStoreInvalidQuery, q.String())
	}

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	crs := C.foreman_metric_resultset_new()

	store.Lock()

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

	store.Unlock()

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

	store.Lock()

	C.foreman_metric_store_query_delete_expired_metrics(store.cStore)

	store.Unlock()

	return nil
}

// String returns a string description of the instance
func (store *cgoStore) String() string {
	name := ""
	ver := ""

	var cname *C.char
	if C.foreman_metric_store_getname(store.cStore, &cname) {
		name = C.GoString(cname)
	}

	var cver *C.char
	if C.foreman_metric_store_getversion(store.cStore, &cver) {
		ver = C.GoString(cver)
	}

	return fmt.Sprintf("%s/%s", name, ver)
}
