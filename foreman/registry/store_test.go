// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

import (
	"fmt"
	"testing"
)

const (
	testStoreLoopCount = 100
)

func testCreateRootObjects(t *testing.T, store Store) {
	err := store.Clear()
	if err != nil {
		t.Error(t)
	}

	q := NewQuery()
	q.ParentID = RootObjectID
	objs, err := store.Browse(q)
	if len(objs) != 0 {
		t.Errorf("%d != %d", len(objs), 0)
		return
	}

	// Insert

	for n := 0; n < testStoreLoopCount; n++ {
		name := fmt.Sprintf("name%d", n)
		data := fmt.Sprintf("data%d", n)

		obj := NewObject()
		obj.ParentID = RootObjectID
		obj.Name = name
		obj.Data = data
		err := store.CreateObject(obj)
		if err != nil {
			t.Error(t)
		}
	}

	// Get

	q = NewQuery()
	q.ParentID = RootObjectID
	objs, err = store.Browse(q)
	if len(objs) != testStoreLoopCount {
		t.Errorf("%d != %d", len(objs), testStoreLoopCount)
		return
	}

	for n := 0; n < testStoreLoopCount; n++ {
		obj := objs[n]
		if !obj.IsRootParentID() {
			t.Errorf("%s != %s", obj.ParentID, RootObjectID)
		}
	}

	// Delete

	q = NewQuery()
	q.ParentID = RootObjectID
	objs, err = store.Browse(q)
	if len(objs) != testStoreLoopCount {
		t.Errorf("%d != %d", len(objs), testStoreLoopCount)
		return
	}

	for n := 0; n < testStoreLoopCount; n++ {
		obj := objs[n]
		err := store.DeleteObject(obj.ID)
		if err != nil {
			t.Errorf("Could not remove [%s]", obj.ID)
		}
	}
}

func testStore(t *testing.T, store Store) {
	err := store.Open()
	if err != nil {
		t.Error(t)
	}

	testCreateRootObjects(t, store)

	err = store.Close()
	if err != nil {
		t.Error(t)
	}
}

func TestNewSQLiteStore(t *testing.T) {
	testStore(t, NewSQLiteStore())
}
