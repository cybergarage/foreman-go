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

// ActionObject represents an abstract interface for the action.
type ActionObject interface {
	ProcessEvent(e *Event) (ResultSet, error)
}

// RouteSource represents an abstract interface for the route source object.
type RouteSource interface {
	RouteObject
}

// RouteDestination represents an abstract interface for the route destination object.
type RouteDestination interface {
	RouteObject
	ActionObject
}

// Route represents a route.
type Route struct {
	Name        string
	Source      string
	Destination string
}

// NewRouteWithStrings returns a new route with the specified string parameters.
func NewRouteWithStrings(name string, src string, dest string) *Route {
	route := &Route{
		Name:        name,
		Source:      src,
		Destination: dest,
	}
	return route
}

// GetSource returns the source name
func (route *Route) GetSource() string {
	return route.Source
}

// GetDestination returns the destination name
func (route *Route) GetDestination() string {
	return route.Destination
}

// GetName returns the name
func (route *Route) GetName() string {
	return route.Name
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
	return fmt.Sprintf("%s : %s TO %s", route.Name, route.Source, route.Destination)
}
