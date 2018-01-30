// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

// Store represents an abstract interface of register.
type Store interface {
	Open() error
	Close() error
	Clear() error

	SetObject(obj Object) error
	GetObject(key string) (Object, bool)

	Size() int64

	String() string
}
