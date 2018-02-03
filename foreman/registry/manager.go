// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package registry

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/fql"
)

// Manager represents a metric manager.
type Manager struct {
	*Store
	fql.QueryExecutor
}

// NewManager returns a new metric manager.
func NewManager() *Manager {
	mgr := &Manager{
		Store: NewStore(),
	}

	return mgr
}

// CreateCategoryObject create the specified category object under the root.
func (mgr *Manager) CreateCategoryObject(name string) error {
	obj := NewObject()
	obj.ParentID = RootObjectID
	obj.Name = name

	err := mgr.CreateObject(obj)
	if err != nil {
		return err
	}

	return nil
}

// GetCategoryObject returns a specified category object.
func (mgr *Manager) GetCategoryObject(name string) (*Object, error) {
	q := NewQuery()
	q.ParentID = RootObjectID

	objs, err := mgr.Browse(q)
	if err != nil {
		return nil, err
	}

	for _, obj := range objs {
		if obj.IsName(name) {
			return obj, nil
		}
	}

	return nil, fmt.Errorf(errorCategoryNotFound, name)
}

// Start starts the manager.
func (mgr *Manager) Start() error {
	err := mgr.Store.Open()
	if err != nil {
		mgr.Stop()
		return err
	}

	return nil
}

// Stop stops the manager.
func (mgr *Manager) Stop() error {
	err := mgr.Store.Close()
	if err != nil {
		return err
	}

	return nil
}
