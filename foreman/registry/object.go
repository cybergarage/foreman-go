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
func (obj *Object) isID(id string) bool {
	if obj.ID != id {
		return false
	}
	return true
}

// IsParentID returns whether the object has a specified ID.
func (obj *Object) IsParentID(id string) bool {
	if obj.ParentID != id {
		return false
	}
	return true
}

// IsRootParentID returns whether the object has a root parent ID.
func (obj *Object) IsRootParentID() bool {
	return obj.IsParentID(RootObjectID)
}

// IsName returns whether the object has a specified name.
func (obj *Object) IsName(name string) bool {
	if obj.Name != name {
		return false
	}
	return true
}

// IsData returns whether the object has a specified data.
func (obj *Object) IsData(data string) bool {
	if obj.Data != data {
		return false
	}
	return true
}

// String returns a string description of the instance
func (obj *Object) String() string {
	return fmt.Sprintf("[%s] : %s", obj.ID, obj.Name)
}
