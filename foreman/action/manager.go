// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package action provides scripting interfaces.
package action

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/logging"
)

const (
	managerPostEventRouteMessageFormat = "EXECUTE ROUTE (%s) FROM %s TO %s"
)

const (
	errorInvalidRouteNoDestination = "Invalid Route (%s) : destination object is not found"
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

// executeEventRoute executes only the specified event route.
func (mgr *Manager) executeEventRoute(e *Event, route *Route) error {
	dest := mgr.findRouteDestination(route.GetDestination())
	if dest == nil {
		err := fmt.Errorf(errorInvalidRouteNoDestination, route.GetName())
		logging.Error(err.Error())
		return err
	}

	src := e.GetSource()
	logging.Info(managerPostEventRouteMessageFormat, route.GetName(), src.GetName(), dest.GetName())

	results, err := dest.ProcessEvent(e)
	if err != nil {
		logging.Error(err.Error())
		return err
	}

	// Executes the child routes recursively when the route was executed normally

	destEvent := NewEventWithSource(dest)
	destEvent.AddResultSet(results)
	mgr.PostEvent(destEvent)

	return nil
}

// PostEvent execute an posted event.
func (mgr *Manager) PostEvent(e *Event) error {
	src := e.GetSource()
	routes, ok := mgr.GetRoutes(src.GetName())
	if !ok {
		return nil
	}

	for _, route := range routes {
		go mgr.executeEventRoute(e, route)
	}

	return nil
}

// Start starts the manager
func (mgr *Manager) Start() error {
	return nil
}

// Stop stops the manager
func (mgr *Manager) Stop() error {
	return nil
}

// Clear removes all actions and routes int the manager
func (mgr *Manager) Clear() error {
	var lastError error

	err := mgr.ScriptManager.RemoveAllMethods()
	if err != nil {
		lastError = err
	}

	err = mgr.RouteManager.RemoveAllRoutes()
	if err != nil {
		lastError = err
	}

	return lastError
}
