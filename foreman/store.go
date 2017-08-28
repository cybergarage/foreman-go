// Copyright 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

// #include <foreman/foreman-c.h>
import "C"

// Store represents a Foreman Store.
type Store struct {
	cstore *C.ForemanStore
}

// NewStore returns a new Store.
func NewStore() *Store {
	store := &Store{}

	return store
}
