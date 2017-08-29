// Copyright 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"testing"
)

func testStore(t *testing.T, store *Store) {
	err := store.Open()
	if err != nil {
		t.Error(t)
	}
	err = store.Close()
	if err != nil {
		t.Error(t)
	}
}

func TestNewSQLiteStore(t *testing.T) {
	testStore(t, NewSQLiteStore())
}
