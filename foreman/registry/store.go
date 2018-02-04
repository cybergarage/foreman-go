// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

// Storing represents an abstract interface of registry.
type Storing interface {
	Open() error
	Close() error
	Clear() error

	// To insert a new object using CreateObject, you must set at least the name and parent ID.
	// You can set the ID explicitly too, but the ID is generated automatically when the ID is not specified.
	// The generated ID is stored in Object.ID when the object is inserted normally.
	// The automatic generated ID is guaranteed the uniqueness of the registry store.
	CreateObject(obj *Object) error

	// To update the inserted object using UpdateObject, you must set the name, data and ID.
	// UpdateObject can only update the name and data, and it can't update the ID and parent ID.
	UpdateObject(obj *Object) error

	GetObject(objID string) (*Object, error)
	DeleteObject(objID string) error

	Browse(q *Query) ([]*Object, error)
	Search(q *Query) ([]*Object, error)

	String() string
}

// Store represents an registry store.
type Store struct {
	Storing
}
