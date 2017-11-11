// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

import (
	"fmt"
)

// Object represents a meta object in the registry store.
type Object struct {
	ID       string // ID is guaranteed the uniqueness in the registry store basically. "0" is a special ID, and it is reserved for the root object.
	ParentID string
	Name     string
	Data     string
}

// NewObject returns a new object.
func NewObject() *Object {
	m := &Object{}
	return m
}

// isID returns whether the object has a specified ID.
func (self *Object) isID(id string) bool {
	if self.ID != id {
		return false
	}
	return true
}

// IsRootParentID returns whether the object has a specified ID.
func (self *Object) IsParentID(id string) bool {
	if self.ParentID != id {
		return false
	}
	return true
}

// IsRootParentID returns whether the object has a root parent ID.
func (self *Object) IsRootParentID() bool {
	return self.IsParentID(RootObjectID)
}

// IsName returns whether the object has a specified name.
func (self *Object) IsName(name string) bool {
	if self.Name != name {
		return false
	}
	return true
}

// IsData returns whether the object has a specified data.
func (self *Object) IsData(data string) bool {
	if self.Data != data {
		return false
	}
	return true
}

// String returns a string description of the instance
func (self *Object) String() string {
	return fmt.Sprintf("[%s] : %s", self.ID, self.Name)
}
