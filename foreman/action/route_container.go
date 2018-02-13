// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// RouteContainer represents an abstract interface for the route container.
type RouteContainer interface {
	// FindRouteSource searches a source object with the specified regex name.
	FindRouteSource(name string) RouteSource
	// FindRouteDestination searches a destination object with the specified regex name.
	FindRouteDestination(name string) RouteSource
}
