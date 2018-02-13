// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// Scripting represents an abstract interface of the script manager
type Scripting interface {
	HasEngine(method string) bool

	AddMethod(method *Method) error
	RemoveMethod(method string) error
	HasMethod(method string) bool
	RemoveAllMethods() error

	GetMethod(name string) *Method
	GetFirstMethod() *Method
	GetNextMethod(method *Method) *Method

	ExecMethod(name string, params Parameters) (Parameters, error)
}

// ScriptManager represents an interface of the script manager
type ScriptManager struct {
	Scripting
	RouteContainer
}

// FindRouteSource searches a source object with the specified regex name.
func (mgr *Manager) FindRouteSource(name string) RouteSource {
	return nil
}

// FindRouteDestination searches a destination object with the specified regex name.
func (mgr *Manager) FindRouteDestination(name string) RouteSource {
	return nil
}
