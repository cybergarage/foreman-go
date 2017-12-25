// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

// Storing represents an abstract interface of register.
type Storing interface {
	Open() error
	Close() error
	Clear() error

	SetObject(obj *Object) error
	GetObject(objID string) (*Object, error)
}

// Store represents an register store.
type Store struct {
	Storing
}
