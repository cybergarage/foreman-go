// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package registry

import (
	"testing"
)

func TestNewManager(t *testing.T) {
	mgr := NewManager()

	err := mgr.Start()
	if err != nil {
		t.Error(err)
	}

	err = mgr.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestCreateCategory(t *testing.T) {
	mgr := NewManager()

	err := mgr.Start()
	if err != nil {
		t.Error(err)
	}

	categoryName := "test"

	err = mgr.CreateCategoryObject(categoryName)
	if err != nil {
		t.Error(err)
	}

	_, err = mgr.GetCategoryObject(categoryName)
	if err != nil {
		t.Error(err)
	}

	err = mgr.Stop()
	if err != nil {
		t.Error(err)
	}
}
