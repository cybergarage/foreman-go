// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"github.com/cybergarage/foreman-go/foreman/fql"
)

// Manager represents a metric manager.
type Manager struct {
	*QoS
	fql.QueryExecutor
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

// CreateQoS creates a new QoS with the specified parameters,
func (mgr *Manager) CreateQoS(name string, formula string) error {
	rule, err := mgr.ParseQoSString(formula)
	if err != nil {
		return err
	}

	err = rule.SetName(name)
	if err != nil {
		return err
	}

	err = mgr.SetRule(rule)
	if err != nil {
		return err
	}

	return nil
}
