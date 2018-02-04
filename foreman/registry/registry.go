// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

import "fmt"

// Registry represents an registry.
type Registry struct {
	Store
}

// NewRegistry returns a new registry.
func NewRegistry() *Registry {
	reg := &Registry{
		Store: NewStore(),
	}
	return reg
}

// GetChildObjects returns all child objects.
func (store *Registry) GetChildObjects(parentID string) ([]*Object, error) {
	q := NewQuery()
	q.ParentID = parentID
	return store.Browse(q)
}

// GetRootObjects returns all root objects
func (store *Registry) GetRootObjects() ([]*Object, error) {
	return store.GetChildObjects(RootObjectID)
}

// GetChildObjectByName returns an object which has the specified name under the specified parent object.
func (store *Registry) GetChildObjectByName(parentID string, name string) (*Object, error) {
	objs, err := store.GetChildObjects(parentID)
	if err != nil {
		return nil, err
	}

	for _, obj := range objs {
		if obj.IsName(name) {
			return obj, nil
		}
	}

	return nil, fmt.Errorf(errorNotFoundChild, name)

}

// GetObjectByPath return an object by the specified name path.
func (store *Registry) GetObjectByPath(namePath string) (*Object, error) {
	names, err := registryGetPathStrings(namePath)
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

// GetChildObjectsByPath return all objects by the specified name path.
func (store *Registry) GetChildObjectsByPath(namePath string) ([]*Object, error) {
	pathObj, err := store.GetObjectByPath(namePath)
	if err != nil {
		return nil, err
	}
	if pathObj == nil {
		return []*Object{}, nil
	}
	return store.GetChildObjects(pathObj.ID)
}
