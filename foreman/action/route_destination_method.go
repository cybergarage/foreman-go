// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

type routeDestinationMethod struct {
	RouteDestination
	*ScriptManager
	Name string
}

func newRouteDestinationMethod(scriptMgr *ScriptManager, name string) *routeDestinationMethod {
	rm := &routeDestinationMethod{
		ScriptManager: scriptMgr,
		Name:          name,
	}
	return rm
}

func (m *routeDestinationMethod) GetName() string {
	return m.Name
}

func (m *routeDestinationMethod) String() string {
	return m.Name
}

func (m *routeDestinationMethod) ProcessEvent(e *Event) (ResultSet, error) {
	rs, err := m.ScriptManager.ExecMethod(m.Name, e.GetParameters())
	if err != nil {
		return nil, err
	}
	return rs, nil
}
