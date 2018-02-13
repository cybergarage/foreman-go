// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"fmt"
)

// RouteObject represents an abstract interface for the route object.
type RouteObject interface {
	GetName() string
	String() string
}

// RouteSource represents an abstract interface for the route source object.
type RouteSource interface {
	RouteObject
}

// RouteDestination represents an abstract interface for the route destination object.
type RouteDestination interface {
	RouteObject
	ProcessEvent(e *Event) (ResultSet, error)
}

// Route represents a route.
type Route struct {
	Name        string
	Source      RouteSource
	Destination RouteDestination
}

// NewRouteWithObjects returns a new boolean parameter.
func NewRouteWithObjects(name string, srcObj RouteSource, destObj RouteDestination) *Route {
	route := &Route{
		Name:        name,
		Source:      srcObj,
		Destination: destObj,
	}
	return route
}

// Equals returns true when the specified routes aresame, otherwise false.
func (route *Route) Equals(other *Route) bool {
	if route.String() != other.String() {
		return false
	}
	return true
}

// String returns a string description
func (route *Route) String() string {
	return fmt.Sprintf("%s : %s TO %s", route.Name, route.Source.String(), route.Destination.String())
}
