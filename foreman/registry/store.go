// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

import (
	"strings"
)

// Storing represents an abstract interface of registry.
type Storing interface {
	Open() error
	Close() error
	Clear() error

	// To insert a new object using CreateObject, you must set at least the name and parent ID.
	// You can set the ID explicitly too, but the ID is generated automatically when the ID is not specified.
	// The generated ID is stored in Object.ID when the object is inserted normally.
	// The automatic generated ID is guaranteed the uniqueness of the registry store.
	CreateObject(obj *Object) error

	// To update the inserted object using UpdateObject, you must set the name, data and ID.
	// UpdateObject can only update the name and data, and it can't update the ID and parent ID.
	UpdateObject(obj *Object) error

	GetObject(objID string) (*Object, error)
	DeleteObject(objID string) error

	Browse(q *Query) ([]*Object, error)
	Search(q *Query) ([]*Object, error)

	String() string
}

// Store represents an registry store.
type Store struct {
	Storing
}

// getStorePathStrings splits the specified path string into an string array.
func getStorePathStrings(path string) ([]string, error) {
	allNames := strings.Split(path, PathDelim)
	names := make([]string, 0)
	for _, name := range allNames {
		if len(name) <= 0 {
			continue
		}
		names = append(names, name)
	}
	return names, nil
}

// getStoreObjectByName return a object which has the specified name in a specified object array.
func getStoreObjectByName(objs []*Object, name string) (*Object, error) {
	for _, obj := range objs {
		if obj.IsName(name) {
			return obj, nil
		}
	}
	return nil, nil
}

// GetChildObjects returns all child objects.
func (store *Store) GetChildObjects(parentID string) ([]*Object, error) {
	q := NewQuery()
	q.ParentID = parentID
	return store.Browse(q)
}

// GetRootObjects returns all root objects
func (store *Store) GetRootObjects() ([]*Object, error) {
	return store.GetChildObjects(RootObjectID)
}

// GetChildObjectByName returns an object which has the specified name under the specified parent object.
func (store *Store) GetChildObjectByName(parentID string, name string) (*Object, error) {
	objs, err := store.GetChildObjects(parentID)
	if err != nil {
		return nil, err
	}
	return getStoreObjectByName(objs, name)
}

// GetObjectByNamePath return an object by the specified name path.
func (store *Store) GetObjectByNamePath(namePath string) (*Object, error) {
	names, err := getStorePathStrings(namePath)
	if err != nil {
		return nil, err
	}

	var lastObj *Object
	for n, name := range names {
		if n == 0 {
			lastObj, err = store.GetChildObjectByName(RootObjectID, name)
		} else {
			lastObj, err = store.GetChildObjectByName(lastObj.ID, name)
		}
		if err != nil {
			return nil, err
		}
		if lastObj == nil {
			return nil, nil
		}
	}

	return lastObj, nil
}

// GetChildObjectsByNamePath return all objects by the specified name path.
func (store *Store) GetChildObjectsByNamePath(namePath string) ([]*Object, error) {
	pathObj, err := store.GetObjectByNamePath(namePath)
	if err != nil {
		return nil, err
	}
	if pathObj == nil {
		return []*Object{}, nil
	}
	return store.GetChildObjects(pathObj.ID)
}
