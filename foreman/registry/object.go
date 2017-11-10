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
	ID       string
	ParentID string
	Name     string
	Data     string
}

// NewObject returns a new object.
func NewObject() *Object {
	m := &Object{}
	return m
}

// IsRootParentID returns whether the object has a root parent ID.
func (self *Object) IsRootParentID() bool {
	if self.ParentID != RootObjectID {
		return false
	}
	return true
}

// String returns a string description of the instance
func (self *Object) String() string {
	return fmt.Sprintf("[%s] : %s", self.ID, self.Name)
}
