// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

// Property represents a property in the registry object.
type Property interface {
	SetName(name string)
	GetName() string

	SetString(value string) error
	GetString() (string, error)

	SetInteger(value int) error
	GetInteger() (int, error)

	SetReal(value float64) error
	GetRealInteger() (float64, error)
}
