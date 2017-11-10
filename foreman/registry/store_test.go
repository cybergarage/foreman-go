// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

import (
	"fmt"
	"reflect"
	"testing"
)

const (
	testStoreLoopCount = 100
)

func testCreateRootObjects(t *testing.T, store *Store) {
	err := store.Clear()
	if err != nil {
		t.Error(err)
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
			t.Error(err)
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

		fetchObj, err := store.GetObject(obj.ID)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(obj, fetchObj) {
			t.Errorf("%s != %s", obj, fetchObj)
		}
	}

	// Updated

	q = NewQuery()
	q.ParentID = RootObjectID
	objs, err = store.Browse(q)
	if len(objs) != testStoreLoopCount {
		t.Errorf("%d != %d", len(objs), testStoreLoopCount)
		return
	}

	for n := 0; n < testStoreLoopCount; n++ {
		obj := objs[n]

		name := fmt.Sprintf("name%d", n)
		data := fmt.Sprintf("data%d", n)
		obj.Name = name
		obj.Data = data
		err := store.UpdateObject(obj)
		if err != nil {
			t.Error(err)
		}

		fetchObj, err := store.GetObject(obj.ID)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(obj, fetchObj) {
			t.Errorf("%s != %s", obj, fetchObj)
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

func testStore(t *testing.T, store *Store) {
	err := store.Open()
	if err != nil {
		t.Error(err)
	}

	testCreateRootObjects(t, store)

	err = store.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestNewSQLiteStore(t *testing.T) {
	testStore(t, NewSQLiteStore())
}
