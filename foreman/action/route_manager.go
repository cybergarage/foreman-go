// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

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

// CreateRoute tries to create a new route with the specified parameters.
func (mgr *RouteManager) CreateRoute(name string, src string, dst string) error {
	route := NewRouteWithStrings(name, src, dst)

	err := mgr.AddRoute(route)
	if err != nil {
		return err
	}

	return nil
}

// AddRoute adds a new route.
func (mgr RouteManager) AddRoute(route *Route) error {
	key := route.GetSource()
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
func (mgr RouteManager) RemoveAllRoutes() error {
	mgr.routeMap = make(map[string][]*Route)
	return nil
}
