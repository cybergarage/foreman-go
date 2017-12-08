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
