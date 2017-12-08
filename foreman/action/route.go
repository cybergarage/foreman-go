// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"fmt"
)

// RouteObject represents an abstract interface for the route objects.
type RouteObject interface {
	GetName() string
	ProcessEvent(e *Event) error
	String() string
}

// Route represents a route.
type Route struct {
	Destination RouteObject
	Source      RouteObject
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
