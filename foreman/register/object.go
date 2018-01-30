// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

import (
	"time"
	"unsafe"
)

// ObjectListener represents an interface of the object in the register store.
type ObjectListener interface {
	ObjectUpdated()
}

// Object represents an register object in thestore.
type Object interface {
	// GetCObject returns the native object
	GetCObject() unsafe.Pointer
	// SetName sets a specified name
	SetName(name string) error
	// GetName returns the object name
	GetName() (string, error)
	// SetData sets a specified data
	SetData(data interface{}) error
	// GetData returns the object data
	GetData() (interface{}, error)
	// GetVersion returns a version number of the instance
	GetVersion() (int64, error)
	// GetTimestamp returns a timestamp of the instance
	GetTimestamp() (time.Time, error)
	// UpdateVersion updates the internal timestamp and version, and sends the update messages to the listeners
	UpdateVersion() error
	// AddListener adds a new listener
	AddListener(listener ObjectListener) bool
	// RemoveListener removes the specified listener
	RemoveListener(listener ObjectListener) bool
	// GetListeners returns the current listeners
	GetListeners() []ObjectListener
	// String returns a string description of the instance
	String() string
}
