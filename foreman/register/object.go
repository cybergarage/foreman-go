// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

import (
	"fmt"
	"time"
)

// Object represents a meta object in the register store.
type Object struct {
	Name      string
	Data      interface{}
	version   int64
	timestamp time.Time
}

// NewObject returns a new object.
func NewObject() *Object {
	m := &Object{
		version: 0,
	}
	return m
}

// GetName returns the object name
func (obj *Object) GetName() string {
	return obj.Name
}

// GetData returns the object data
func (obj *Object) GetData() interface{} {
	return obj.Data
}

// GetVersion returns a version number of the instance
func (obj *Object) GetVersion() int64 {
	return obj.version
}

// GetTimestamp returns a timestamp of the instance
func (obj *Object) GetTimestamp() time.Time {
	return obj.timestamp
}

// incrementVersion updates the current version number.
func (obj *Object) incrementVersion() error {
	obj.version++
	return nil
}

// updateTimestamp updates the current time stamp.
func (obj *Object) updateTimestamp() error {
	obj.timestamp = time.Now()
	return nil
}

// String returns a string description of the instance
func (obj *Object) String() string {
	return fmt.Sprintf("[%s] : %s", obj.Name, obj.Data)
}
