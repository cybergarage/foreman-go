// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// #include <foreman/foreman-c.h>
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

const (
	errorMetricStoreNotFound = "Store (%s) Not Found"
)

func newStoreWithInterface(storeImpl Storing) *Store {
	store := &Store{
		Storing: storeImpl,
	}
	return store
}

func newStoreWithCObject(cObject unsafe.Pointer) *Store {
	storeImp := &cgoStore{
		cStore:   cObject,
		listener: nil,
	}
	storeImp.SetRetentionInterval(DefaultRetentionInterval)
	runtime.SetFinalizer(storeImp, storeFinalizer)
	return newStoreWithInterface(storeImp)
}

func storeFinalizer(self *cgoStore) {
	if self.cStore != nil {
		if C.foreman_metric_store_delete(self.cStore) {
			self.cStore = nil
		}
	}
}

// NewSQLiteStore returns a new store of SQLite.
func NewSQLiteStore() *Store {
	store := newStoreWithCObject(C.foreman_metric_store_sqlite_create())
	return store
}

// NewEmptyStore returns a new empty store.
func NewEmptyStore() *Store {
	store := newStoreWithCObject(C.foreman_metric_store_empty_create())
	return store
}

// NewGorillaStore returns a new store of Facebook's Gorilla.
func NewGorillaStore() *Store {
	store := newStoreWithCObject(C.foreman_metric_store_tsmap_create())
	return store
}

// NewRingMapStore returns a new store of RingMap.
func NewRingMapStore() *Store {
	store := newStoreWithCObject(C.foreman_metric_store_ringmap_create())
	return store
}

// NewDefaultStore returns a new default store.
func NewDefaultStore() *Store {
	store := NewEmptyStore()
	return store
}

// newDefaultTestStore returns a new default store for testing.
func newDefaultTestStore() *Store {
	store := NewSQLiteStore()
	return store
}

// NewStoreWithName returns a new store with specified name.
func NewStoreWithName(name string) (*Store, error) {
	switch name {
	case MetricStoreEmpty:
		return NewEmptyStore(), nil
	case MetricStoreSqlite:
		return NewEmptyStore(), nil
	case MetricStoreTsmap:
		return NewGorillaStore(), nil
	case MetricStoreRingmap:
		return NewRingMapStore(), nil
	}
	return nil, fmt.Errorf(errorMetricStoreNotFound, name)
}
