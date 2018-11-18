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
		"\"route02\" : { \"" + RouteColumnSource + "\" : \"src02\", \"" + RouteColumnDestination + "\" : \"dst02\"}," +
		"\"route03\" : { \"" + RouteColumnSource + "\" : \"src03\", \"" + RouteColumnDestination + "\" : \"dst03\"}," +
		"\"route04\" : { \"" + RouteColumnSource + "\" : \"src04\", \"" + RouteColumnDestination + "\" : \"dst04\"}," +
		"\"route05\" : { \"" + RouteColumnSource + "\" : \"src05\", \"" + RouteColumnDestination + "\" : \"dst05\"}," +
		"\"route06\" : { \"" + RouteColumnSource + "\" : \"src06\", \"" + RouteColumnDestination + "\" : \"dst06\"}" +
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
