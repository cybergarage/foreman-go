// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

// Store represents an abstract interface of registry.
type Store interface {
	Open() error
	Close() error
	Clear() error

	CreateObject(obj *Object) error
	UpdateObject(obj *Object) error
	GetObject(objID string) (*Object, error)
	DeleteObject(objID string) (*Object, error)

	Browse(q *Query) ([]*Object, error)
	Search(q *Query) ([]*Object, error)

	String() string
}
