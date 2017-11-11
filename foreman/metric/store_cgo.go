// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog
import "C"

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

// Store represents a metric store for Foreman.
type cgoStore struct {
	cStore unsafe.Pointer
}

// Open initializes the store.
func (self *cgoStore) Open() error {
	if self.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	if !C.foreman_metric_store_open(self.cStore) {
		return fmt.Errorf(errorStoreCouldNotOpen, self)
	}

	return nil
}

// Close closes the store.
func (self *cgoStore) Close() error {
	if self.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	if !C.foreman_metric_store_close(self.cStore) {
		return fmt.Errorf(errorStoreCouldNotClose, self)
	}

	return nil
}

// Clear remove all inserted data.
func (self *cgoStore) Clear() error {
	if self.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	if !C.foreman_metric_store_clear(self.cStore) {
		return fmt.Errorf(errorStoreCouldNotClose, self)
	}

	return nil
}

// SetRetentionInterval sets the retention duration.
func (self *cgoStore) SetRetentionInterval(value time.Duration) error {
	if self.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	C.foreman_metric_store_setretentioninterval(self.cStore, C.time_t(value.Seconds()))

	return nil
}

// GetRetentionInterval returns the retention duration.
func (self *cgoStore) GetRetentionInterval() (time.Duration, error) {
	if self.cStore == nil {
		return 0, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	durationSec := C.foreman_metric_store_getretentioninterval(self.cStore)
	duration := time.Second * time.Duration(durationSec)

	return duration, nil
}

// AddMetric adds a new metric.
func (self *cgoStore) AddMetric(m *Metric) error {
	if self.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cm, err := m.CMetric()
	if err != nil {
		return err
	}

	isSuccess, err := C.foreman_metric_store_addmetric(self.cStore, cm)
	if err != nil {
		return err
	}
	if !isSuccess {
		return fmt.Errorf(errorStoreCouldNotAddMetric, m.String())
	}

	return nil
}

// Query gets the specified metrics.
func (self *cgoStore) Query(q *Query) (ResultSet, error) {
	if self.cStore == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	duration, err := self.GetRetentionInterval()
	if err != nil {
		return nil, err
	}
	q.Interval = duration

	cq, err := q.CQuery()
	if err != nil {
		return nil, err
	}

	crs := C.foreman_metric_resultset_new()

	if !C.foreman_metric_store_query(self.cStore, cq, crs) {
		C.foreman_metric_resultset_delete(crs)
		return nil, fmt.Errorf(errorStoreCouldNotAddMetric, q.String())
	}

	return NewResultSetWithCObject(crs), nil
}

// String returns a string description of the instance
func (self *cgoStore) String() string {
	// FIXME : Not implemented
	return fmt.Sprintf("")
}
