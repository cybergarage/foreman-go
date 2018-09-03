// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"encoding/base64"

	"github.com/cybergarage/foreman-go/foreman/register"
	"github.com/cybergarage/foreman-go/foreman/registry"
)

// Scripting represents an abstract interface of the script manager
type Scripting interface {
	SetRegistryStore(store registry.Store) error
	SetRegisterStore(store register.Store) error

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

// GetMethodCount returns a count of all methods.
func (mgr *ScriptManager) GetMethodCount() int {
	methodCnt := 0
	method := mgr.GetFirstMethod()
	for method != nil {
		methodCnt++
		method = mgr.GetNextMethod(method)
	}
	return methodCnt
}

// CreateMethod creates a new method with the specified parameters
func (mgr *ScriptManager) CreateMethod(name string, lang string, code string, enc string) error {
	m := NewMethodWithLanguage(lang)
	m.Name = name
	if 0 < len(enc) && enc == ActionEncodingBase64 {
		code, err := base64.StdEncoding.DecodeString(code)
		if err != nil {
			return err
		}
		m.Code = code
	} else {
		m.Code = []byte(code)
	}

	err := mgr.AddMethod(m)
	if err != nil {
		return err
	}

	return nil
}

// FindRouteDestination returns a destination object with the specified name.
func (mgr *ScriptManager) FindRouteDestination(name string) RouteDestination {
	method := mgr.GetMethod(name)
	if method == nil {
		return nil
	}
	return newRouteDestinationMethod(mgr, method.GetName())
}
