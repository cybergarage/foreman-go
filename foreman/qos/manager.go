// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"github.com/cybergarage/foreman-go/foreman/action"
	"github.com/cybergarage/foreman-go/foreman/fql"
)

// Manager represents a metric manager.
type Manager struct {
	*QoS
	action.RouteContainer
	fql.QueryExecutor
}

// NewManager returns a new metric manager.
func NewManager() *Manager {
	mgr := &Manager{
		QoS: NewQoS(),
	}
	return mgr
}

// FindRouteSource searches a source object with the specified regex name.
func (mgr *Manager) FindRouteSource(name string) action.RouteSource {
	rule, ok := mgr.GetRule(name)
	if !ok {
		return nil
	}
	return rule
}

// FindRouteDestination searches a destination object with the specified regex name.
func (mgr *Manager) FindRouteDestination(name string) action.RouteSource {
	return nil
}

// Start starts the manager.
func (mgr *Manager) Start() error {
	return nil
}

// Stop stops the manager.
func (mgr *Manager) Stop() error {
	return nil
}
