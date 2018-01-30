// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

import (
	"fmt"
	"time"
)

// BaseObject represents a meta object in the register store.
type BaseObject struct {
	Object
	name      string
	version   int64
	timestamp time.Time
	listeners []ObjectListener
}

// NewObject returns a new base object.
func NewObject() *BaseObject {
	obj := &BaseObject{
		version:   0,
		timestamp: time.Now(),
		listeners: make([]ObjectListener, 0),
	}
	return obj
}

// SetName sets a specified name
func (obj *BaseObject) SetName(name string) error {
	obj.name = name
	return nil
}

// GetName returns the object name
func (obj *BaseObject) GetName() (string, error) {
	return obj.name, nil
}

// GetVersion returns a version number of the instance
func (obj *BaseObject) GetVersion() int64 {
	return obj.version
}

// GetTimestamp returns a timestamp of the instance
func (obj *BaseObject) GetTimestamp() time.Time {
	return obj.timestamp
}

// UpdateVersion updates the internal timestamp and version, and sends the update messages to the listeners
func (obj *BaseObject) UpdateVersion() error {
	obj.version++
	obj.timestamp = time.Now()

	for _, l := range obj.listeners {
		l.ObjectUpdated()
	}

	return nil
}

// HasListener checks whether the specified listener is already added
func (obj *BaseObject) HasListener(listener ObjectListener) bool {
	for _, l := range obj.listeners {
		if l == listener {
			return true
		}
	}
	return false
}

// AddListener adds a new listener
func (obj *BaseObject) AddListener(listener ObjectListener) bool {
	if obj.HasListener(listener) {
		return false
	}
	obj.listeners = append(obj.listeners, listener)
	return true
}

// RemoveListener removes the specified listener
func (obj *BaseObject) RemoveListener(listener ObjectListener) bool {
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
func (obj *BaseObject) GetListeners() []ObjectListener {
	return obj.listeners
}

// String returns a string description of the instance
func (obj *BaseObject) String() string {
	return fmt.Sprintf("[%s] (%d:%d)", obj.name, obj.version, obj.timestamp.Unix())
}
