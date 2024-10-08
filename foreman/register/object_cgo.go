// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

// #include <foreman/foreman-c.h>
import "C"

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

// cgoObject represents a meta object in the register store.
type cgoObject struct {
	Object
	cObject   unsafe.Pointer
	listeners []ObjectListener
}

// newObjectWithCObject returns a new object with the specified C++ object.
func newObjectWithCObject(cObject unsafe.Pointer) Object {
	obj := &cgoObject{
		cObject:   cObject,
		listeners: make([]ObjectListener, 0),
	}
	runtime.SetFinalizer(obj, registerObjectFinalizer)
	return obj
}

// registerObjectFinalizer deletes the C object when the Go object is deprecated
func registerObjectFinalizer(obj *cgoObject) {
	if obj.cObject != nil {
		if C.foreman_register_object_delete(obj.cObject) {
			obj.cObject = nil
		}
	}
}

// NewObject returns a new object.
func NewObject() Object {
	return newObjectWithCObject(C.foreman_register_object_new())
}

// GetCObject returns a stored C object
func (obj *cgoObject) GetCObject() unsafe.Pointer {
	return obj.cObject
}

// SetName sets a specified name
func (obj *cgoObject) SetName(name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	if !C.foreman_register_object_setkey(obj.cObject, cname) {
		return fmt.Errorf(errorInvalidObject, obj.String())
	}
	return nil
}

// GetName returns the object name
func (obj *cgoObject) GetName() (string, error) {
	var name *C.char
	if !C.foreman_register_object_getkey(obj.cObject, &name) {
		return "", fmt.Errorf(errorInvalidObject, obj.String())
	}
	return C.GoString(name), nil
}

// SetData sets a specified data
func (obj *cgoObject) SetData(data string) error {
	cdata := C.CString(data)
	defer C.free(unsafe.Pointer(cdata))
	if !C.foreman_register_object_setdata(obj.cObject, cdata) {
		return fmt.Errorf(errorInvalidObject, obj.String())
	}

	// Notify to the listeners
	for _, l := range obj.listeners {
		l.ObjectUpdated()
	}

	return nil
}

// GetData returns the object data
func (obj *cgoObject) GetData() (string, error) {
	var data *C.char
	if !C.foreman_register_object_getdata(obj.cObject, &data) {
		return "", fmt.Errorf(errorInvalidObject, obj.String())
	}
	return C.GoString(data), nil
}

// GetVersion returns a version number of the instance
func (obj *cgoObject) GetVersion() (int64, error) {
	var ver C.long
	if !C.foreman_register_object_getversion(obj.cObject, &ver) {
		return 0, fmt.Errorf(errorInvalidObject, obj.String())
	}
	return int64(ver), nil
}

// GetTimestamp returns a timestamp of the instance
func (obj *cgoObject) GetTimestamp() (time.Time, error) {
	var ts C.time_t
	if !C.foreman_register_object_gettimestamp(obj.cObject, &ts) {
		return time.Now(), fmt.Errorf(errorInvalidObject, obj.String())
	}
	return time.Unix(int64(ts), 0), nil
}

// UpdateVersion updates the internal timestamp and version, and sends the update messages to the listeners
func (obj *cgoObject) UpdateVersion() error {
	return nil
}

// HasListener checks whether the specified listener is already added
func (obj *cgoObject) HasListener(listener ObjectListener) bool {
	for _, l := range obj.listeners {
		if l == listener {
			return true
		}
	}
	return false
}

// AddListener adds a new listener
func (obj *cgoObject) AddListener(listener ObjectListener) bool {
	if obj.HasListener(listener) {
		return false
	}
	obj.listeners = append(obj.listeners, listener)
	return true
}

// RemoveListener removes the specified listener
func (obj *cgoObject) RemoveListener(listener ObjectListener) bool {
	for n, l := range obj.listeners {
		if l == listener {
			lastIndex := len(obj.listeners) - 1
			obj.listeners[lastIndex], obj.listeners[n] = obj.listeners[n], obj.listeners[lastIndex]
			obj.listeners = obj.listeners[:lastIndex]
			return true
		}
	}
	return false
}

// GetListeners returns the current listeners
func (obj *cgoObject) GetListeners() []ObjectListener {
	return obj.listeners
}

// String returns a string description of the instance
func (obj *cgoObject) String() string {
	name, err := obj.GetName()
	if err != nil {
		return ""
	}
	ver, err := obj.GetVersion()
	if err != nil {
		return ""
	}
	ts, err := obj.GetTimestamp()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("[%s] (%d:%d)", name, ver, ts.Unix())
}
