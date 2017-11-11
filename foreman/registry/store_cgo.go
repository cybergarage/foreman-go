// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

// Store represents a registry store for Foreman.
type cgoStore struct {
	cStore unsafe.Pointer
}

// Open initializes the store.
func (self *cgoStore) Open() error {
	if self.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	if !C.foreman_registry_store_open(self.cStore) {
		return fmt.Errorf(errorStoreCouldNotOpen, self)
	}

	return nil
}

// Close closes the store.
func (self *cgoStore) Close() error {
	if self.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	if !C.foreman_registry_store_close(self.cStore) {
		return fmt.Errorf(errorStoreCouldNotClose, self)
	}

	return nil
}

// Clear remove all registry data.
func (self *cgoStore) Clear() error {
	if self.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	if !C.foreman_registry_store_clear(self.cStore) {
		return fmt.Errorf(errorStoreCouldNotClear, self)
	}

	return nil
}

// CreateObject insert a new object
func (self *cgoStore) CreateObject(obj *Object) error {
	if self.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cobj, err := obj.CObject()
	if err != nil {
		return err
	}
	defer C.foreman_registry_object_delete(cobj)

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_registry_store_createobject(self.cStore, cobj, cerr) {
		return errors.NewWithCObject(cerr)
	}

	var cID *C.char
	if C.foreman_registry_object_getid(cobj, &cID) {
		obj.ID = C.GoString(cID)
	}

	return nil
}

// CreateObject insert a new object
func (self *cgoStore) UpdateObject(obj *Object) error {
	if self.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cobj, err := obj.CObject()
	if err != nil {
		return err
	}
	defer C.foreman_registry_object_delete(cobj)

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_registry_store_updateobject(self.cStore, cobj, cerr) {
		return errors.NewWithCObject(cerr)
	}

	return nil
}

// GetObject gets a specified object.
func (self *cgoStore) GetObject(objID string) (*Object, error) {
	if self.cStore == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cobj := C.foreman_registry_object_new()
	defer C.foreman_registry_object_delete(cobj)

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_registry_store_getobject(self.cStore, C.CString(objID), cobj, cerr) {
		return nil, errors.NewWithCObject(cerr)
	}

	return newObjectWithCObject(cobj), nil
}

// DeleteObject deletes a specified object.
func (self *cgoStore) DeleteObject(objID string) error {
	if self.cStore == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_registry_store_deleteobject(self.cStore, C.CString(objID), cerr) {
		return errors.NewWithCObject(cerr)
	}

	return nil
}

// Browse returns a child objects
func (self *cgoStore) Browse(q *Query) ([]*Object, error) {
	if self.cStore == nil {
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

	if !C.foreman_registry_store_browse(self.cStore, cq, cobjs, cerr) {
		return nil, errors.NewWithCObject(cerr)
	}

	return newObjectsWithCObjects(cobjs), nil
}

// Search finds a specified objects.
func (self *cgoStore) Search(q *Query) ([]*Object, error) {
	if self.cStore == nil {
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

	if !C.foreman_registry_store_search(self.cStore, cq, cobjs, cerr) {
		return nil, errors.NewWithCObject(cerr)
	}

	return newObjectsWithCObjects(cobjs), nil
}

// String returns a string description of the instance
func (self *cgoStore) String() string {
	if self.cStore == nil {
		return ""
	}
	return fmt.Sprintf("%s/%s", C.GoString(C.foreman_registry_store_gettype(self.cStore)), C.GoString(C.foreman_registry_store_getversion(self.cStore)))
}
