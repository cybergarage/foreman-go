// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"testing"
)

func TestNewRoute(t *testing.T) {
	NewRouteWithObjects(nil, nil)
	newRouteSourceWithName("")
	newRouteSourceWithScriptManagerAndName(nil, "")
}
