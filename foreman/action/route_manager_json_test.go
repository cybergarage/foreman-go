// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"testing"
)

func TestRouteJSONExportImport(t *testing.T) {
	testRouteCount := 6
	testRouteJSON := "{" +
		"\"route01\" : { \"" + RouteColumnSource + "\" : \"src01\", \"" + RouteColumnDestination + "\" : \"dst01\"}," +
		"\"route02\" : { \"" + RouteColumnSource + "\" : \"src01\", \"" + RouteColumnDestination + "\" : \"dst01\"}," +
		"\"route03\" : { \"" + RouteColumnSource + "\" : \"src01\", \"" + RouteColumnDestination + "\" : \"dst01\"}," +
		"\"route04\" : { \"" + RouteColumnSource + "\" : \"src01\", \"" + RouteColumnDestination + "\" : \"dst01\"}," +
		"\"route05\" : { \"" + RouteColumnSource + "\" : \"src01\", \"" + RouteColumnDestination + "\" : \"dst01\"}," +
		"\"route06\" : { \"" + RouteColumnSource + "\" : \"src01\", \"" + RouteColumnDestination + "\" : \"dst01\"}" +
		"}"

	// Import (JSON String)

	mgr := NewRouteManager()

	err := mgr.importRouteJSONString(testRouteJSON)
	if err != nil {
		t.Error(err)
		return
	}

	routes := mgr.GetAllRoutes()
	if len(routes) != testRouteCount {
		t.Errorf("%d != %d", len(routes), testRouteCount)
	}

	// Export (JSON Object)

	routeJSONObj, err := mgr.exportRouteJSONObject()
	if err != nil {
		t.Error(err)
		return
	}

	// Import (JSON Object)

	mgr = NewRouteManager()

	err = mgr.ImportRouteJSONObject(routeJSONObj)
	if err != nil {
		t.Error(err)
		return
	}

	routes = mgr.GetAllRoutes()
	if len(routes) != testRouteCount {
		t.Errorf("%d != %d", len(routes), testRouteCount)
	}
}
