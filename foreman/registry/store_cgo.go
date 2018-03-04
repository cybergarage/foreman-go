// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

// #include <foreman/foreman-c.h>
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

// CgoStore represents a registry store for Foreman.
type CgoStore struct {
	cStore unsafe.Pointer
}

// GetCObject returns the internal Clang object.
func (store *CgoStore) GetCObject() unsafe.Pointer {
	return store.cStore
}

// Open initializes the store.
func (store *CgoStore) Open() error {
	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	if !C.foreman_registry_store_open(store.cStore) {
		return fmt.Errorf(errorStoreCouldNotOpen, store)
	}

	return nil
}

// Close closes the store.
func (store *CgoStore) Close() error {
	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	if !C.foreman_registry_store_close(store.cStore) {
		return fmt.Errorf(errorStoreCouldNotClose, store)
	}

	return nil
}

// Clear remove all registry data.
func (store *CgoStore) Clear() error {
	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	if !C.foreman_registry_store_clear(store.cStore) {
		return fmt.Errorf(errorStoreCouldNotClear, store)
	}

	return nil
}

// CreateObject insert a new object
func (store *CgoStore) CreateObject(obj *Object) error {
	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cobj, err := obj.CObject()
	if err != nil {
		return err
	}
	defer C.foreman_registry_object_delete(cobj)

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_registry_store_createobject(store.cStore, cobj, cerr) {
		return errors.NewWithCObject(cerr).Error()
	}

	var cID *C.char
	if C.foreman_registry_object_getid(cobj, &cID) {
		obj.ID = C.GoString(cID)
	}

	return nil
}

// CreateObject insert a new object
func (store *CgoStore) UpdateObject(obj *Object) error {
	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cobj, err := obj.CObject()
	if err != nil {
		return err
	}
	defer C.foreman_registry_object_delete(cobj)

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_registry_store_updateobject(store.cStore, cobj, cerr) {
		return errors.NewWithCObject(cerr).Error()
	}

	return nil
}

// GetObject gets a specified object.
func (store *CgoStore) GetObject(objID string) (*Object, error) {
	if store.cStore == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cobj := C.foreman_registry_object_new()
	defer C.foreman_registry_object_delete(cobj)

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_registry_store_getobject(store.cStore, C.CString(objID), cobj, cerr) {
		return nil, errors.NewWithCObject(cerr).Error()
	}

	return newObjectWithCObject(cobj), nil
}

// DeleteObject deletes a specified object.
func (store *CgoStore) DeleteObject(objID string) error {
	if store.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_registry_store_deleteobject(store.cStore, C.CString(objID), cerr) {
		return errors.NewWithCObject(cerr).Error()
	}

	return nil
}

// Browse returns a child objects
func (store *CgoStore) Browse(q *Query) ([]*Object, error) {
	if store.cStore == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cq, err := q.CObject()
	if err != nil {
		return nil, err
	}
	defer C.foreman_registry_query_delete(cq)

	cobjs := C.foreman_registry_objects_new()
	defer C.foreman_registry_objects_delete(cobjs)

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_registry_store_browse(store.cStore, cq, cobjs, cerr) {
		return nil, errors.NewWithCObject(cerr).Error()
	}

	return newObjectsWithCObjects(cobjs), nil
}

// Search finds a specified objects.
func (store *CgoStore) Search(q *Query) ([]*Object, error) {
	if store.cStore == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cq, err := q.CObject()
	if err != nil {
		return nil, err
	}
	defer C.foreman_registry_query_delete(cq)

	cobjs := C.foreman_registry_objects_new()
	defer C.foreman_registry_objects_delete(cobjs)

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_registry_store_search(store.cStore, cq, cobjs, cerr) {
		return nil, errors.NewWithCObject(cerr).Error()
	}

	return newObjectsWithCObjects(cobjs), nil
}

// String returns a string description of the instance
func (store *CgoStore) String() string {
	if store.cStore == nil {
		return ""
	}
	return fmt.Sprintf("%s/%s", C.GoString(C.foreman_registry_store_gettype(store.cStore)), C.GoString(C.foreman_registry_store_getversion(store.cStore)))
}
