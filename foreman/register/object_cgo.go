// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -llua -lpython
import "C"

import (
	"fmt"
	"time"
	"unsafe"
)

// cgoObject represents a meta object in the register store.
type cgoObject struct {
	Object
	cObject   unsafe.Pointer
	listeners []ObjectListener
}

// NewObjectWithCObject returns a new object with the C++ object.
func newObjectWithCObject(cObject unsafe.Pointer) *cgoObject {
	obj := &cgoObject{
		cObject:   cObject,
		listeners: make([]ObjectListener, 0),
	}
	return obj
}

// NewObject returns a new base object.
func NewObject() *cgoObject {
	return newObjectWithCObject(C.foreman_register_object_new())
}

// SetName sets a specified name
func (obj *cgoObject) SetName(name string) error {
	ok := C.foreman_register_object_setkey(obj.cObject, C.CString(name))
	if !ok {
		return fmt.Errorf(errorInvalidObject, obj.String())
	}
	return nil
}

// GetName returns the object name
func (obj *cgoObject) GetName() (string, error) {
	var name *C.char
	ok := C.foreman_register_object_getkey(obj.cObject, &name)
	if !ok {
		return "", fmt.Errorf(errorInvalidObject, obj.String())
	}
	return C.GoString(name), nil
}

// GetVersion returns a version number of the instance
func (obj *cgoObject) GetVersion() (int64, error) {
	var ver C.long
	ok := C.foreman_register_object_getversion(obj.cObject, &ver)
	if !ok {
		return 0, fmt.Errorf(errorInvalidObject, obj.String())
	}
	return int64(ver), nil
}

// GetTimestamp returns a timestamp of the instance
func (obj *cgoObject) GetTimestamp() (time.Time, error) {
	var ts C.time_t
	ok := C.foreman_register_object_gettimestamp(obj.cObject, &ts)
	if !ok {
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
