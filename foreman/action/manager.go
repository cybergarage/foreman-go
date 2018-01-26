// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package action provides scripting interfaces.
package action

// Manager represents an action manager.
type Manager struct {
	*ScriptManager
	*RouteManager
}

// NewManager returns a new action manager.
func NewManager() *Manager {
	mgr := &Manager{
		ScriptManager: NewScriptManager(),
		RouteManager:  NewRouteManager(),
	}
	return mgr
}

// PostEvent posts an event.
func (mgr *Manager) PostEvent(e *Event) error {
	routes := mgr.FindRoutesBySourceObject(e.GetSource())

	// TODO : Update to call parallel.
	for _, route := range routes {
		route.Destination.ProcessEvent(e)
	}

	return nil
}
