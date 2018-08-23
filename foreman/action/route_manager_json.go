// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

const (
	errorRouteInvalidJSONObject = "Invalid JSON Object : %v"
	errorRouteNotFound          = "Route (%s) Not Found"
)

func (mgr *RouteManager) importRouteJSONString(jsonStr string) error {
	jsonDecorder := json.NewDecorder()
	jsonObj, err := jsonDecorder.Decode(jsonStr)
	if err != nil {
		return err
	}
	return mgr.importRouteJSONObject(jsonObj)
}

func (mgr *RouteManager) importRouteJSONObject(jsonObj interface{}) error {
	routesMap, ok := jsonObj.(map[string]interface{})
	if !ok {
		return fmt.Errorf(errorRouteInvalidJSONObject, jsonObj)
	}

	for name, routeMapObj := range routesMap {
		routeMap := json.NewPathWithObject(routeMapObj)

		src, err := routeMap.GetPathString(RouteColumnSource)
		if err != nil {
			return fmt.Errorf(errorRouteInvalidJSONObject, routeMapObj)
		}

		dst, err := routeMap.GetPathString(RouteColumnDestination)
		if err != nil {
			return fmt.Errorf(errorRouteInvalidJSONObject, routeMapObj)
		}

		err = mgr.CreateRoute(name, src, dst)
		if err != nil {
			return err
		}
	}

	return nil
}

func (mgr *RouteManager) exportRouteJSONObjectWithName(name string) (interface{}, error) {
	hasName := false
	if 0 < len(name) {
		hasName = true
	}

	routes := map[string]interface{}{}
	for _, route := range mgr.GetAllRoutes() {
		if hasName {
			if route.Name != name {
				continue
			}
		}

		routeMap := map[string]string{}
		routeMap[RouteColumnSource] = route.GetSource()
		routeMap[RouteColumnDestination] = route.GetDestination()

		routes[route.GetName()] = routeMap
	}

	if hasName && (len(routes) <= 0) {
		return nil, fmt.Errorf(errorRouteNotFound, name)
	}

	return routes, nil
}

func (mgr *RouteManager) exportRouteJSONObject() (interface{}, error) {
	return mgr.exportRouteJSONObjectWithName("")
}
