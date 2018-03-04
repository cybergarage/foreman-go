// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

// #include <foreman/foreman-c.h>
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
	storeImp := &CgoStore{}
	storeImp.cStore = cObject
	runtime.SetFinalizer(storeImp, storeFinalizer)
	return newStoreWithInterface(storeImp)
}

func storeFinalizer(self *CgoStore) {
	if self.cStore != nil {
		if C.foreman_registry_store_delete(self.cStore) {
			self.cStore = nil
		}
	}
}

// NewSQLiteStore returns a new SQLite store.
func NewSQLiteStore() *Store {
	store := newStoreWithCObject(C.foreman_registry_sqlite_store_new())
	return store
}

// NewStore returns a new store.
func NewStore() *Store {
	return NewSQLiteStore()
}
