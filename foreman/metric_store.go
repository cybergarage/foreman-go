// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

// #include <foreman/foreman-c.h>
import "C"
import "unsafe"
import "fmt"
import "time"

// MetricStore represents a metric store for Foreman.
type MetricStore struct {
	cStore unsafe.Pointer
}

// Open initializes the store.
func (self *MetricStore) Open() error {
	if self.cStore == nil {
		return fmt.Errorf(errorClangObjectNotInitialized)
	}

	if !C.foreman_store_open(self.cStore) {
		return fmt.Errorf(errorStoreCouldNotOpen, self)
	}

	return nil
}

// Close closes the store.
func (self *MetricStore) Close() error {
	if self.cStore == nil {
		return fmt.Errorf(errorClangObjectNotInitialized)
	}

	if !C.foreman_store_close(self.cStore) {
		return fmt.Errorf(errorStoreCouldNotClose, self)
	}

	return nil
}

// SetRetentionInterval sets the retention duration.
func (self *MetricStore) SetRetentionInterval(value time.Duration) error {
	if self.cStore == nil {
		return fmt.Errorf(errorClangObjectNotInitialized)
	}

	C.foreman_store_setretentioninterval(self.cStore, C.time_t(value.Seconds()))

	return nil
}

// GetRetentionInterval returns the retention duration.
func (self *MetricStore) GetRetentionInterval() (time.Duration, error) {
	if self.cStore == nil {
		return 0, fmt.Errorf(errorClangObjectNotInitialized)
	}

	durationSec := C.foreman_store_getretentioninterval(self.cStore)
	duration := time.Second * time.Duration(durationSec)

	return duration, nil
}

// AddMetric adds a new metric.
func (self *MetricStore) AddMetric(m *Metric) error {
	if self.cStore == nil {
		return fmt.Errorf(errorClangObjectNotInitialized)
	}

	cm, err := m.CMetric()
	if err != nil {
		return err
	}

	isSuccess, err := C.foreman_store_addmetric(self.cStore, cm)
	if err != nil {
		return err
	}
	if !isSuccess {
		return fmt.Errorf(errorStoreCouldNotAddMetric, m.String())
	}

	return nil
}

// Query gets the specified metrics.
func (self *MetricStore) Query(q *MetricQuery) (*MetricResultSet, error) {
	if self.cStore == nil {
		return nil, fmt.Errorf(errorClangObjectNotInitialized)
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

	crs := C.foreman_resultset_new()

	if !C.foreman_store_query(self.cStore, cq, crs) {
		C.foreman_resultset_delete(crs)
		return nil, fmt.Errorf(errorStoreCouldNotAddMetric, q.String())
	}

	return NewMetricResultSetWithCObject(crs), nil
}

// String returns a string description of the instance
func (self *MetricStore) String() string {
	// FIXME : Not implemented
	return fmt.Sprintf("")
}
