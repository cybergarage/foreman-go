// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

import (
	"fmt"
	"time"
)

// ObjectListener represents an interface of the meta object.
type ObjectListener interface {
	ObjectUpdated()
}

// Object represents a meta object in the register store.
type Object struct {
	name      string
	version   int64
	timestamp time.Time
	listeners []ObjectListener
}

// NewObject returns a new object.
func NewObject() *Object {
	m := &Object{
		version:   0,
		listeners: make([]ObjectListener, 0),
	}
	return m
}

// SetName sets a specified name
func (obj *Object) SetName(name string) error {
	obj.name = name
	return nil
}

// GetName returns the object name
func (obj *Object) GetName() string {
	return obj.name
}

// GetVersion returns a version number of the instance
func (obj *Object) GetVersion() int64 {
	return obj.version
}

// GetTimestamp returns a timestamp of the instance
func (obj *Object) GetTimestamp() time.Time {
	return obj.timestamp
}

// UpdateVersion updates the internal timestamp and version, and sends the update messages to the listeners
func (obj *Object) UpdateVersion() error {
	obj.version++
	obj.timestamp = time.Now()

	for _, l := range obj.listeners {
		l.ObjectUpdated()
	}

	return nil
}

// HasListener checks whether the specified listener is already added
func (obj *Object) HasListener(listener ObjectListener) bool {
	for _, l := range obj.listeners {
		if l == listener {
			return true
		}
	}
	return false
}

// AddListener adds a new listener
func (obj *Object) AddListener(listener ObjectListener) bool {
	if obj.HasListener(listener) {
		return false
	}
	obj.listeners = append(obj.listeners, listener)
	return true
}

// RemoveListener removes the specified listener
func (obj *Object) RemoveListener(listener ObjectListener) bool {
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
func (obj *Object) GetListeners() []ObjectListener {
	return obj.listeners
}

// String returns a string description of the instance
func (obj *Object) String() string {
	return fmt.Sprintf("[%s] (%d:%d)", obj.name, obj.version, obj.timestamp.Unix())
}
