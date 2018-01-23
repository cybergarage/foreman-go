// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

// Manager represents a metric manager.
type Manager struct {
	*Store
	*Register
}

// NewManager returns a new metric manager.
func NewManager() *Manager {
	mgr := &Manager{
		Store:    NewStore(),
		Register: NewRegister(),
	}

	mgr.Store.SetListener(mgr)

	return mgr
}

// SetRegisterListener sets a register listener.
func (mgr *Manager) SetRegisterListener(listener RegisterListener) {
	mgr.Register.Listener = listener
}

// MetricAdded is a listener for Store
func (mgr *Manager) MetricAdded(m *Metric) {
	mgr.Register.UpdateMetric(m)
}

// Start starts the manager.
func (mgr *Manager) Start() error {
	err := mgr.Store.Open()
	if err != nil {
		mgr.Stop()
		return err
	}

	err = mgr.Register.Open()
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

	err = mgr.Register.Close()
	if err != nil {
		return err
	}

	return nil
}
