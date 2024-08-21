// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

// #include <foreman/foreman-c.h>
import "C"

// CgoStore represents an register store.
type CgoStore struct {
	Store
	cStore unsafe.Pointer
}

// newStoreWithCObject returns a new object with the C++ object.
func newStoreWithCObject(cObject unsafe.Pointer) *CgoStore {
	store := &CgoStore{
		cStore: cObject,
	}
	runtime.SetFinalizer(store, storeObjectFinalizer)
	return store
}

// storeObjectFinalizer deletes the C object when the Go object is deprecated
func storeObjectFinalizer(store *CgoStore) {
	if store.cStore != nil {
		if C.foreman_register_store_delete(store.cStore) {
			store.cStore = nil
		}
	}
}

// NewStore returns a new base object.
func NewStore() *CgoStore {
	return newStoreWithCObject(C.foreman_register_store_new())
}

// GetCObject returns the internal Clang object.
func (store *CgoStore) GetCObject() unsafe.Pointer {
	return store.cStore
}

// Open opens the store.
func (store *CgoStore) Open() error {
	if !C.foreman_register_store_open(store.cStore) {
		return fmt.Errorf(errorInvalidStore, store.String())
	}
	return nil
}

// Close closes the store.
func (store *CgoStore) Close() error {
	if !C.foreman_register_store_close(store.cStore) {
		return fmt.Errorf(errorInvalidStore, store.String())
	}
	return nil
}

// Clear clears all registers.
func (store *CgoStore) Clear() error {
	if !C.foreman_register_store_clear(store.cStore) {
		return fmt.Errorf(errorInvalidStore, store.String())
	}
	return nil
}

// SetObject sets a object into the store.
func (store *CgoStore) SetObject(obj Object) error {
	if obj == nil {
		return errorNilObject
	}
	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_register_store_setobject(store.cStore, obj.GetCObject(), cerr) {
		return errors.NewWithCObject(cerr).Error()
	}

	return nil
}

// GetObject gets the specified object.
func (store *CgoStore) GetObject(key string) (Object, bool) {
	cObj := C.foreman_register_object_new()

	cErr := C.foreman_error_new()
	defer C.foreman_error_delete(cErr)

	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	if !C.foreman_register_store_getobject(store.cStore, ckey, cObj, cErr) {
		C.foreman_register_object_delete(cObj)
		return nil, false
	}

	return newObjectWithCObject(cObj), true
}

// RemoveObject removes the specified key object.
func (store *CgoStore) RemoveObject(key string) bool {
	cErr := C.foreman_error_new()
	defer C.foreman_error_delete(cErr)

	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	if !C.foreman_register_store_removeobject(store.cStore, ckey, cErr) {
		return false
	}

	return true
}

// Size returns the object count.
func (store *CgoStore) Size() int64 {
	return int64(C.foreman_register_store_size(store.cStore))
}

// String returns a string description of the instance
func (store *CgoStore) String() string {
	return fmt.Sprintf("%s/%s", C.GoString(C.foreman_register_store_gettype(store.cStore)), C.GoString(C.foreman_register_store_getversion(store.cStore)))
}
