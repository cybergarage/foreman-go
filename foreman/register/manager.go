// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

import (
	"github.com/cybergarage/foreman-go/foreman/fql"
)

// Manager represents a metric manager.
type Manager struct {
	Store
	fql.QueryExecutor
}

// NewManager returns a new metric manager.
func NewManager() *Manager {
	mgr := &Manager{
		Store: NewStore(),
	}

	return mgr
}

// GetStore returns an internal store
func (mgr *Manager) GetStore() Store {
	return mgr.Store
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
