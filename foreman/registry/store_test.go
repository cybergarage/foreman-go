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
	testStoreLoopCount         = 100
	testObjectNameFormat       = "name%d"
	testObjectDataFormat       = "data%d"
	testObjectUpdateNameFormat = "new_name%d"
	testObjectUpdateDataFormat = "new_data%d"
)

func TestPathSpecialSplit(t *testing.T) {
	specialPaths := []string{"", "/", "/0"}
	specialPathCounts := []int{0, 0, 1}
	for n, path := range specialPaths {
		name, err := getStorePathStrings(path)
		if err != nil {
			t.Error(err)
		}
		if len(name) != specialPathCounts[n] {
			t.Errorf("%d != %d", len(name), specialPathCounts[n])
		}
	}
}

func TestPathSplit(t *testing.T) {
	for n := 0; n < testStoreLoopCount; n++ {
		path := ""
		names := make([]string, n)
		for i := 0; i < n; i++ {
			name := fmt.Sprintf("name%d", i)
			names[i] = name
			path += fmt.Sprintf("%s%s", PathDelim, name)
		}

		parseNames, err := getStorePathStrings(path)
		if err != nil {
			t.Error(err)
		}
		if len(parseNames) != n {
			t.Errorf("%d != %d", len(parseNames), n)
		}
		for i := 0; i < n; i++ {
			name := fmt.Sprintf("name%d", i)
			if parseNames[i] != name {
				t.Errorf("%s != %s", parseNames[i], name)
			}
		}
	}
}

func ceateRootObjects(t *testing.T, store *Store) {
	for n := 0; n < testStoreLoopCount; n++ {
		name := fmt.Sprintf(testObjectNameFormat, n)
		data := fmt.Sprintf(testObjectDataFormat, n)

		obj := NewObject()
		obj.ParentID = RootObjectID
		obj.Name = name
		obj.Data = data
		err := store.CreateObject(obj)
		if err != nil {
			t.Error(err)
		}
	}
}

func testCreateRootObjects(t *testing.T, store *Store) {
	err := store.Clear()
	if err != nil {
		t.Error(err)
		return
	}

	objs, err := store.GetRootObjects()
	if len(objs) != 0 {
		t.Errorf("%d != %d", len(objs), 0)
		return
	}

	// Insert

	ceateRootObjects(t, store)

	// Get

	objs, err = store.GetRootObjects()
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

	// Update

	objs, err = store.GetRootObjects()
	if len(objs) != testStoreLoopCount {
		t.Errorf("%d != %d", len(objs), testStoreLoopCount)
		return
	}

	for n := 0; n < testStoreLoopCount; n++ {
		obj := objs[n]

		name := fmt.Sprintf(testObjectUpdateNameFormat, n)
		data := fmt.Sprintf(testObjectUpdateDataFormat, n)
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
			t.Errorf("'%s' != '%s'", obj, fetchObj)
		}
	}

	// Delete

	objs, err = store.GetRootObjects()
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

	objs, err = store.GetRootObjects()
	if len(objs) != 0 {
		t.Errorf("%d != %d", len(objs), 0)
		return
	}
}

func createHierarchicalObjects(t *testing.T, store *Store) {
	parentID := RootObjectID
	for n := 0; n < testStoreLoopCount; n++ {
		obj := NewObject()
		obj.ParentID = parentID
		obj.Name = fmt.Sprintf(testObjectNameFormat, n)
		obj.Data = fmt.Sprintf(testObjectDataFormat, n)

		err := store.CreateObject(obj)
		if err != nil {
			t.Error(err)
		}

		parentID = obj.ID
	}
}

func testHierarchicalObjects(t *testing.T, store *Store) {
	err := store.Clear()
	if err != nil {
		t.Error(err)
	}

	objs, err := store.GetRootObjects()
	if len(objs) != 0 {
		t.Errorf("%d != %d", len(objs), 0)
		return
	}

	// Insert

	createHierarchicalObjects(t, store)

	// Get

	parentID := RootObjectID
	for n := 0; n < testStoreLoopCount; n++ {
		objs, err := store.GetChildObjects(parentID)
		if err != nil {
			t.Error(err)
		}
		if len(objs) != 1 {
			t.Errorf("%d != %d", len(objs), 1)
		}

		obj := objs[0]
		objName := fmt.Sprintf(testObjectNameFormat, n)
		if !obj.IsName(objName) {
			t.Errorf("%s != %s", obj.Name, objName)
		}

		parentID = obj.ID
	}

	// Get (Path)

	path := ""
	for n := 0; n < testStoreLoopCount; n++ {
		var objs []*Object
		var err error

		if n == 0 {
			objs, err = store.GetRootObjects()
		} else {
			objs, err = store.GetChildObjectsByNamePath(path)
		}

		if err != nil {
			t.Error(err)
		}
		if len(objs) != 1 {
			t.Errorf("%d != %d", len(objs), 1)
		}

		obj := objs[0]
		objName := fmt.Sprintf(testObjectNameFormat, n)
		if !obj.IsName(objName) {
			t.Errorf("%s != %s", obj.Name, objName)
		}

		path += fmt.Sprintf("%s%s", PathDelim, obj.Name)
	}

	// Clean

	err = store.Clear()
	if err != nil {
		t.Error(err)
	}
}

func testStore(t *testing.T, store *Store) {
	err := store.Open()
	if err != nil {
		t.Error(err)
	}

	testCreateRootObjects(t, store)
	testHierarchicalObjects(t, store)

	err = store.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestNewSQLiteStore(t *testing.T) {
	testStore(t, NewSQLiteStore())
}
