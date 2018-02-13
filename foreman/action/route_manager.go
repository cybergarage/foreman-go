// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import "regexp"

// RouteManager represents a route map.
type RouteManager struct {
	routeMap map[string][]*Route
}

// NewRouteManager returns a null route map.
func NewRouteManager() *RouteManager {
	mgr := &RouteManager{
		routeMap: make(map[string][]*Route),
	}
	return mgr
}

// AddRoute adds a new route.
func (mgr RouteManager) AddRoute(route *Route) error {
	key := route.Source.GetName()
	routes, ok := mgr.routeMap[key]
	if !ok {
		routes = make([]*Route, 0)
	}
	mgr.routeMap[key] = append(routes, route)
	return nil
}

// GetRoutes returns a route lists by the specified key.
func (mgr RouteManager) GetRoutes(key string) ([]*Route, bool) {
	routes, ok := mgr.routeMap[key]
	return routes, ok
}

// GetAllRoutes returns all routes.
func (mgr RouteManager) GetAllRoutes() []*Route {
	allRoutes := make([]*Route, 0)
	for _, routes := range mgr.routeMap {
		allRoutes = append(allRoutes, routes...)
	}
	return allRoutes
}

// RemoveRoute deletes a route with the specified route name.
func (mgr RouteManager) RemoveRoute(name string) bool {
	for key, routes := range mgr.routeMap {
		for n, route := range routes {
			if route.Name == name {
				mgr.routeMap[key] = append(routes[:n], routes[n+1:]...)
				return true
			}
		}
	}
	return true
}

// RemoveAllRoutes deletes all routes.
func (mgr RouteManager) RemoveAllRoutes() bool {
	mgr.routeMap = make(map[string][]*Route)
	return true
}

// FindRoutesBySourceRegKey returns a route lists by the specified regex key.
func (mgr RouteManager) FindRoutesBySourceRegKey(regKey string) []*Route {
	r := regexp.MustCompile(regKey)
	foundRoutes := make([]*Route, 0)
	for key, routes := range mgr.routeMap {
		if !r.MatchString(key) {
			continue
		}
		foundRoutes = append(foundRoutes, routes...)
	}
	return foundRoutes
}

// FindRoutesBySourceObject returns the target routes for the specified source objects.
func (mgr RouteManager) FindRoutesBySourceObject(srcObject RootSource) []*Route {
	return mgr.FindRoutesBySourceRegKey(srcObject.GetName())
}
