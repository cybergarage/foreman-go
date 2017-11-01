// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

// Object represents a meta object in the registry store.
type Object interface {
	GetID() string
	GetParentID() string
	GetName() string

	SetString(data string) error
	GetString() (string, error)

	GetProperty(name string) (string, error)
	GetAllProperties() ([]Property, error)

	String() string
}
