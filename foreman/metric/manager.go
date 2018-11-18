// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/register"
)

// Manager represents a metric manager.
type Manager struct {
	*Store
	*Register
	fql.QueryExecutor
}

// NewManagerWithStore returns a new metric manager with the specified store.
func NewManagerWithStore(store *Store) *Manager {
	mgr := &Manager{
		Store:    store,
		Register: nil,
	}
	return mgr
}

// NewManagerWithStoreName returns a new metric manager with the specified store name.
func NewManagerWithStoreName(name string) (*Manager, error) {
	store, err := NewStoreWithName(name)
	if err != nil {
		return nil, err
	}
	return NewManagerWithStore(store), nil
}

// SetRegisterStore sets a raw register store.
func (mgr *Manager) SetRegisterStore(store register.Store) error {
	mgr.Register = NewRegisterWithStore(store)
	mgr.Store.SetStoreListener(mgr.Register)
	return nil
}

// Start starts the manager.
func (mgr *Manager) Start() error {
	if mgr.Register == nil {
		return fmt.Errorf(errorNullRegisterStore)
	}

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

// String return the store infomation.

func (mgr *Manager) String() string {
	return mgr.Store.String()
}
