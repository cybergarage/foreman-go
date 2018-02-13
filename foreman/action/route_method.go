// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

type routeMethod struct {
	RouteDestination
	*ScriptManager
	Name string
}

func newRouteSourceWithScriptManagerAndName(scriptMgr *ScriptManager, name string) *routeMethod {
	rm := &routeMethod{
		ScriptManager: scriptMgr,
		Name:          name,
	}
	return rm
}

func (rm *routeMethod) GetName() string {
	return rm.Name
}

func (rm *routeMethod) String() string {
	return rm.Name
}

func (rm *routeMethod) ProcessEvent(e *Event) (ResultSet, error) {
	rs, err := rm.ScriptManager.ExecMethod(rm.Name, e.GetParameters())
	if err != nil {
		return nil, err
	}
	return rs, nil
}
