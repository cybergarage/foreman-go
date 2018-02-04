// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

import (
	"fmt"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -llua -lpython
import "C"

// cgoStore represents an register store.
type cgoStore struct {
	Store
	cStore unsafe.Pointer
}

// newStoreWithCObject returns a new object with the C++ object.
func newStoreWithCObject(cObject unsafe.Pointer) *cgoStore {
	store := &cgoStore{
		cStore: cObject,
	}
	return store
}

// NewStore returns a new base object.
func NewStore() *cgoStore {
	return newStoreWithCObject(C.foreman_register_store_new())
}

// Open opens the store.
func (store *cgoStore) Open() error {
	if !C.foreman_register_store_open(store.cStore) {
		return fmt.Errorf(errorInvalidStore, store.String())
	}
	return nil
}

// Close closes the store.
func (store *cgoStore) Close() error {
	if !C.foreman_register_store_close(store.cStore) {
		return fmt.Errorf(errorInvalidStore, store.String())
	}
	return nil
}

// Clear clears all registers.
func (store *cgoStore) Clear() error {
	if !C.foreman_register_store_clear(store.cStore) {
		return fmt.Errorf(errorInvalidStore, store.String())
	}
	return nil
}

// SetObject sets a object into the store.
func (store *cgoStore) SetObject(obj Object) error {
	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_register_store_setobject(store.cStore, obj.GetCObject(), cerr) {
		return errors.NewWithCObject(cerr).Error()
	}

	return nil
}

// GetObject gets the specified object.
func (store *cgoStore) GetObject(key string) (Object, bool) {
	cObj := C.foreman_register_object_new()

	cErr := C.foreman_error_new()
	defer C.foreman_error_delete(cErr)

	if !C.foreman_register_store_getobject(store.cStore, C.CString(key), cObj, cErr) {
		C.foreman_register_object_delete(cObj)
		return nil, false
	}

	return newObjectWithCObject(cObj), true
}

// Size returns the object count.
func (store *cgoStore) Size() int64 {
	return int64(C.foreman_register_store_size(store.cStore))
}

// String returns a string description of the instance
func (store *cgoStore) String() string {
	return fmt.Sprintf("%s/%s", C.GoString(C.foreman_register_store_gettype(store.cStore)), C.GoString(C.foreman_register_store_getversion(store.cStore)))
}
