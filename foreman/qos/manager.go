// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

// Manager represents a metric manager.
type Manager struct {
	*QoS
}

// NewManager returns a new metric manager.
func NewManager() *Manager {
	mgr := &Manager{
		QoS: NewQoS(),
	}
	return mgr
}

// Start starts the manager.
func (mgr *Manager) Start() error {
	return nil
}

// Stop stops the manager.
func (mgr *Manager) Stop() error {
	return nil
}
