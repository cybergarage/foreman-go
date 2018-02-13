// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package action provides scripting interfaces.
package action

import (
	"fmt"

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
func (mgr *Manager) CreateRoute(name string, srcName string, destName string) error {

	// Find a target method of the specified destination name.

	ok := mgr.ScriptManager.HasMethod(srcName)
	if !ok {
		return fmt.Errorf(errorRouteDestinationNotFound, srcName)
	}
	destMethod := newRouteSourceWithScriptManagerAndName(mgr.ScriptManager, destName)

	// Create a route object with the specified source name.

	srcObj := newRouteSourceWithName(srcName)

	// Added a new route

	route := NewRouteWithObjects(name, srcObj, destMethod)

	err := mgr.RouteManager.AddRoute(route)
	if err != nil {
		return err
	}

	return nil
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
