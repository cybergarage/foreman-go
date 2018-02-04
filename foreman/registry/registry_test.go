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
	testregLoopCount           = 100
	testObjectNameFormat       = "name%d"
	testObjectDataFormat       = "data%d"
	testObjectUpdateNameFormat = "new_name%d"
	testObjectUpdateDataFormat = "new_data%d"
)

func TestPathSpecialSplit(t *testing.T) {
	specialPaths := []string{"", "/", "/0"}
	specialPathCounts := []int{0, 0, 1}
	for n, path := range specialPaths {
		name, err := registryGetPathStrings(path)
		if err != nil {
			t.Error(err)
		}
		if len(name) != specialPathCounts[n] {
			t.Errorf("%d != %d", len(name), specialPathCounts[n])
		}
	}
}

func TestPathSplit(t *testing.T) {
	for n := 0; n < testregLoopCount; n++ {
		path := ""
		names := make([]string, n)
		for i := 0; i < n; i++ {
			name := fmt.Sprintf("name%d", i)
			names[i] = name
			path += fmt.Sprintf("%s%s", PathDelim, name)
		}

		parseNames, err := registryGetPathStrings(path)
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

func ceateRootObjects(t *testing.T, reg *Registry) {
	for n := 0; n < testregLoopCount; n++ {
		name := fmt.Sprintf(testObjectNameFormat, n)
		data := fmt.Sprintf(testObjectDataFormat, n)

		obj := NewObject()
		obj.ParentID = RootObjectID
		obj.Name = name
		obj.Data = data
		err := reg.CreateObject(obj)
		if err != nil {
			t.Error(err)
		}
	}
}

func testCreateRootObjects(t *testing.T, reg *Registry) {
	err := reg.Clear()
	if err != nil {
		t.Error(err)
		return
	}

	objs, err := reg.GetRootObjects()
	if len(objs) != 0 {
		t.Errorf("%d != %d", len(objs), 0)
		return
	}

	// Insert

	ceateRootObjects(t, reg)

	// Get

	objs, err = reg.GetRootObjects()
	if len(objs) != testregLoopCount {
		t.Errorf("%d != %d", len(objs), testregLoopCount)
		return
	}

	for n := 0; n < testregLoopCount; n++ {
		obj := objs[n]

		fetchObj, err := reg.GetObject(obj.ID)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(obj, fetchObj) {
			t.Errorf("%s != %s", obj, fetchObj)
		}
	}

	// Update

	objs, err = reg.GetRootObjects()
	if len(objs) != testregLoopCount {
		t.Errorf("%d != %d", len(objs), testregLoopCount)
		return
	}

	for n := 0; n < testregLoopCount; n++ {
		obj := objs[n]

		name := fmt.Sprintf(testObjectUpdateNameFormat, n)
		data := fmt.Sprintf(testObjectUpdateDataFormat, n)
		obj.Name = name
		obj.Data = data
		err := reg.UpdateObject(obj)
		if err != nil {
			t.Error(err)
		}

		fetchObj, err := reg.GetObject(obj.ID)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(obj, fetchObj) {
			t.Errorf("'%s' != '%s'", obj, fetchObj)
		}
	}

	// Delete

	objs, err = reg.GetRootObjects()
	if len(objs) != testregLoopCount {
		t.Errorf("%d != %d", len(objs), testregLoopCount)
		return
	}

	for n := 0; n < testregLoopCount; n++ {
		obj := objs[n]
		err := reg.DeleteObject(obj.ID)
		if err != nil {
			t.Errorf("Could not remove [%s]", obj.ID)
		}
	}

	objs, err = reg.GetRootObjects()
	if len(objs) != 0 {
		t.Errorf("%d != %d", len(objs), 0)
		return
	}
}

func createHierarchicalObjects(t *testing.T, reg *Registry) {
	parentID := RootObjectID
	for n := 0; n < testregLoopCount; n++ {
		obj := NewObject()
		obj.ParentID = parentID
		obj.Name = fmt.Sprintf(testObjectNameFormat, n)
		obj.Data = fmt.Sprintf(testObjectDataFormat, n)

		err := reg.CreateObject(obj)
		if err != nil {
			t.Error(err)
		}

		parentID = obj.ID
	}
}

func testHierarchicalObjects(t *testing.T, reg *Registry) {
	err := reg.Clear()
	if err != nil {
		t.Error(err)
	}

	objs, err := reg.GetRootObjects()
	if len(objs) != 0 {
		t.Errorf("%d != %d", len(objs), 0)
		return
	}

	// Insert

	createHierarchicalObjects(t, reg)

	// Get

	parentID := RootObjectID
	for n := 0; n < testregLoopCount; n++ {
		objs, err := reg.GetChildObjects(parentID)
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
	for n := 0; n < testregLoopCount; n++ {
		var objs []*Object
		var err error

		if n == 0 {
			objs, err = reg.GetRootObjects()
		} else {
			objs, err = reg.GetChildObjectsByPath(path)
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

	err = reg.Clear()
	if err != nil {
		t.Error(err)
	}
}

func testreg(t *testing.T, reg *Registry) {
	err := reg.Open()
	if err != nil {
		t.Error(err)
	}

	testCreateRootObjects(t, reg)
	testHierarchicalObjects(t, reg)

	err = reg.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestNewRgistry(t *testing.T) {
	testreg(t, NewRegistry())
}
