// Copyright 2017 Satoshi Konno. All rights reserved.
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

// NewStore returns a new Store.
func newStoreWithCObject(cObject unsafe.Pointer) *Store {
	store := &Store{}
	store.cStore = cObject
	runtime.SetFinalizer(store, storeFinalizer)
	return store
}

func storeFinalizer(self *Store) {
	if self.cStore != nil {
		if C.foreman_store_delete(self.cStore) {
			self.cStore = nil
		}
	}
}

// NewSQLiteStore returns a new Store.
func NewSQLiteStore() *Store {
	store := newStoreWithCObject(C.foreman_store_tsmap_create())
	return store
}

// NewStore returns a new Store.
func NewStore() *Store {
	return NewSQLiteStore()
}
