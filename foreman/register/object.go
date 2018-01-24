// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

import (
	"time"
)

// ObjectListener represents an interface of the object in the register store.
type ObjectListener interface {
	ObjectUpdated()
}

// Object represents a meta object in the register store.
type Object interface {
	// SetName sets a specified name
	SetName(name string) error
	// GetName returns the object name
	GetName() string
	// GetVersion returns a version number of the instance
	GetVersion() int64
	// GetTimestamp returns a timestamp of the instance
	GetTimestamp() time.Time
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
