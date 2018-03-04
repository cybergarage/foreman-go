// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package registry

import (
	"github.com/cybergarage/foreman-go/foreman/fql"
)

// Manager represents a metric manager.
type Manager struct {
	*Registry
	fql.QueryExecutor
}

// NewManager returns a new metric manager.
func NewManager() *Manager {
	mgr := &Manager{
		Registry: NewRegistry(),
	}

	return mgr
}

// GetStore returns an internal store
func (mgr *Manager) GetStore() *Store {
	return mgr.Registry.Store
}

// CreateCategoryObject create the specified category object under the root.
func (mgr *Manager) CreateCategoryObject(category string) error {
	categoryObj := NewObject()
	categoryObj.ParentID = RootObjectID
	categoryObj.Name = category

	err := mgr.CreateObject(categoryObj)
	if err != nil {
		return err
	}

	return nil
}

// GetCategoryObject returns a specified category object.
func (mgr *Manager) GetCategoryObject(category string) (*Object, error) {
	return mgr.GetChildObjectByName(RootObjectID, category)
}

// CreateCategoryKeyObject create the specified category key object under the category object.
func (mgr *Manager) CreateCategoryKeyObject(category string, key string, value string) error {
	categoryObj, err := mgr.GetCategoryObject(category)
	if err != nil {
		return err
	}

	keyObj := NewObject()
	keyObj.ParentID = categoryObj.ParentID
	keyObj.Name = key
	keyObj.Data = value

	err = mgr.CreateObject(keyObj)
	if err != nil {
		return err
	}

	return nil
}

// GetCategoryKeyObject returns a specified category key object.
func (mgr *Manager) GetCategoryKeyObject(category string, key string) (*Object, error) {
	categoryObj, err := mgr.GetCategoryObject(category)
	if err != nil {
		return nil, err
	}
	return mgr.GetChildObjectByName(categoryObj.ID, key)
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
