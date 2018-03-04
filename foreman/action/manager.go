// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package action provides scripting interfaces.
package action

import (
	"github.com/cybergarage/foreman-go/foreman/fql"
)

// Manager represents an action manager.
type Manager struct {
	*ScriptManager
	*RouteManager
	routeContainers []RouteContainer
	fql.QueryExecutor
}

// NewManager returns a new action manager.
func NewManager() *Manager {
	mgr := &Manager{
		ScriptManager:   NewScriptManager(),
		RouteManager:    NewRouteManager(),
		routeContainers: make([]RouteContainer, 0),
	}

	mgr.AddRouteContainer(mgr.ScriptManager)

	return mgr
}

// AddRouteContainer adds a route container.
func (mgr *Manager) AddRouteContainer(c RouteContainer) error {
	mgr.routeContainers = append(mgr.routeContainers, c)
	return nil
}

// CreateRoute tries to creat a new route with the specified route names.
func (mgr *Manager) CreateRoute(name string, src string, dest string) error {
	// Added a new route

	route := NewRouteWithStrings(name, src, dest)

	err := mgr.RouteManager.AddRoute(route)
	if err != nil {
		return err
	}

	return nil
}

// findRouteDestination returns a destination in the all route containers.
func (mgr *Manager) findRouteDestination(name string) RouteDestination {
	for _, c := range mgr.routeContainers {
		dest := c.FindRouteDestination(name)
		if dest != nil {
			return dest
		}
	}
	return nil
}

// PostEvent execute an posted event.
func (mgr *Manager) PostEvent(e *Event) error {
	routes, ok := mgr.GetRoutes(e.GetSource().GetName())
	if !ok {
		return nil
	}

	// TODO : Update to execute parallel.
	for _, route := range routes {
		dest := mgr.findRouteDestination(route.GetDestination())
		if dest == nil {
			continue
		}
		dest.ProcessEvent(e)
	}

	return nil
}
