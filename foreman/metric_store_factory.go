// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lstdc++ -lsqlite3 -lfolly -lgflags
import "C"

import (
	"runtime"
	"unsafe"
)

// NewMetricStore returns a new MetricStore.
func newMetricStoreWithCObject(cObject unsafe.Pointer) *MetricStore {
	store := &MetricStore{}
	store.cStore = cObject
	runtime.SetFinalizer(store, storeFinalizer)
	return store
}

func storeFinalizer(self *MetricStore) {
	if self.cStore != nil {
		if C.foreman_store_delete(self.cStore) {
			self.cStore = nil
		}
	}
}

// NewSQLiteStore returns a new MetricStore of SQLite.
func NewSQLiteMetricStore() *MetricStore {
	store := newMetricStoreWithCObject(C.foreman_store_sqlite_create())
	return store
}

// NewStore returns a new MetricStore.
func NewMetricStore() *MetricStore {
	return NewSQLiteMetricStore()
}
