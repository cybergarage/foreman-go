// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"fmt"
)

// RouteObject represents an abstract interface for the route object.
type RouteObject interface {
}

// RootSource represents an abstract interface for the route source object.
type RootSource interface {
	RouteObject
	EventSource
}

// RootDestination represents an abstract interface for the route destination object.
type RootDestination interface {
	RouteObject
	ActionObject
}

// Route represents a route.
type Route struct {
	Source      RootSource
	Destination RootDestination
}

// NewRouteWithObjects returns a new boolean parameter.
func NewRouteWithObjects(srcObj RootSource, destObj RootDestination) *Route {
	route := &Route{
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
	return fmt.Sprintf("%s TO %s", route.Source.String(), route.Destination.String())
}
