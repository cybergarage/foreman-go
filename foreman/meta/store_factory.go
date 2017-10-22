// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package meta provides interfaces for MetaStore of Foreman C++.
package meta

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lstdc++ -lsqlite3 -lfolly -lgflags
import "C"

// NewStore returns a new Store.
func NewStore() *Store {
	return nil
}
