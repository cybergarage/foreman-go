// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package meta provides interfaces for MetaStore of Foreman C++.
package meta

// #include <foreman/foreman-c.h>
import "C"
import "unsafe"
import "fmt"

// Store represents a metric store for Foreman.
type Store struct {
	cStore unsafe.Pointer
}

// Open initializes the store.
func (self *Store) Open() error {
	if self.cStore == nil {
		return fmt.Errorf(errorClangObjectNotInitialized)
	}

	return nil
}

// Close closes the store.
func (self *Store) Close() error {
	if self.cStore == nil {
		return fmt.Errorf(errorClangObjectNotInitialized)
	}

	return nil
}

// String returns a string description of the instance
func (self *Store) String() string {
	// FIXME : Not implemented
	return fmt.Sprintf("")
}
