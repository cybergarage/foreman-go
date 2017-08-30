// Copyright 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

// #include <foreman/foreman-c.h>
import "C"
import "unsafe"
import "fmt"

// Store represents a Foreman Store.
type Store struct {
	cStore unsafe.Pointer
}

// Open initializes the store.
func (self *Store) Open() error {
	if self.cStore == nil {
		return fmt.Errorf(errorClangObjectNotInitialized)
	}

	if !C.foreman_store_open(self.cStore) {
		return fmt.Errorf(errorStoreCouldNotOpen, self)
	}

	return nil
}

// Close closes the store.
func (self *Store) Close() error {
	if self.cStore == nil {
		return fmt.Errorf(errorClangObjectNotInitialized)
	}

	if !C.foreman_store_close(self.cStore) {
		return fmt.Errorf(errorStoreCouldNotClose, self)
	}

	return nil
}

// AddMetric adds a new metric.
func (self *Store) AddMetric(m *Metric) error {
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

// Query gets the specfied metrics.
func (self *Store) Query(q *Query) (*ResultSet, error) {
	if self.cStore == nil {
		return nil, fmt.Errorf(errorClangObjectNotInitialized)
	}

	cq, err := q.CQuery()
	if err != nil {
		return nil, err
	}

	crs := C.foreman_resultset_new()

	if !C.foreman_store_query(self.cStore, cq, crs) {
		C.foreman_resultset_delete(crs)
		return nil, fmt.Errorf(errorStoreCouldNotAddMetric, q.String())
	}

	return NewResultSetWithCObject(crs), nil
}

// String returns a string description of the instance
func (self *Store) String() string {
	// FIXME : Not implemented
	return fmt.Sprintf("")
}
