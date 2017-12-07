// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -llua -lpython
import "C"

import (
	"runtime"
	"unsafe"
)

func newStoreWithInterface(storeImpl Storing) *Store {
	store := &Store{
		Storing: storeImpl,
	}
	return store
}

func newStoreWithCObject(cObject unsafe.Pointer) *Store {
	storeImp := &cgoStore{}
	storeImp.cStore = cObject
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

// NewSQLiteStore returns a new Store of SQLite.
func NewSQLiteStore() *Store {
	store := newStoreWithCObject(C.foreman_metric_store_sqlite_create())
	return store
}

// NewStore returns a new Store.
func NewStore() *Store {
	return NewSQLiteStore()
}
