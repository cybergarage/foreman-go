// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog
import "C"

import (
	"runtime"
	"unsafe"
)

// NewStore returns a new Store.
func newStoreWithCObject(cObject unsafe.Pointer) *CgoStore {
	store := &CgoStore{}
	store.cStore = cObject
	runtime.SetFinalizer(store, storeFinalizer)
	return store
}

func storeFinalizer(self *CgoStore) {
	if self.cStore != nil {
		if C.foreman_metric_store_delete(self.cStore) {
			self.cStore = nil
		}
	}
}

// NewSQLiteStore returns a new Store of SQLite.
func NewSQLiteStore() *CgoStore {
	store := newStoreWithCObject(C.foreman_metric_store_sqlite_create())
	return store
}

// NewStore returns a new Store.
func NewStore() *CgoStore {
	return NewSQLiteStore()
}
