// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// #include <foreman/foreman-c.h>
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/cybergarage/foreman-go/foreman/register"
	"github.com/cybergarage/foreman-go/foreman/registry"
)

const (
	cgoScriptManagerExecErrorFormat   = "EXECUTE %s(%s) : %s"
	cgoScriptManagerExecMessageFormat = "EXECUTE %s(%s) -> (%t, %s)"
)

// cgoScriptManager represents an action manager using foreman-cc.
type cgoScriptManager struct {
	cManager unsafe.Pointer
}

// SetRegistryStore sets the registry store to use in the scripts.
func (mgr *cgoScriptManager) SetRegistryStore(store registry.Store) error {
	cStore, ok := store.(*registry.CgoStore)
	if !ok {
		return fmt.Errorf(errorInvalidClangObject)
	}
	cStoreObj := cStore.GetCObject()
	if cStoreObj == nil {
		return fmt.Errorf(errorInvalidClangObject)
	}

	if !C.foreman_action_manager_setregistrystore(mgr.cManager, cStoreObj) {
		return fmt.Errorf(errorInvalidClangObject)
	}
	return nil
}

// SetRegisterStore sets the register store to use in the scripts.
func (mgr *cgoScriptManager) SetRegisterStore(store register.Store) error {
	cStore, ok := store.(*register.CgoStore)
	if !ok {
		return fmt.Errorf(errorInvalidClangObject)
	}
	cStoreObj := cStore.GetCObject()
	if cStoreObj == nil {
		return fmt.Errorf(errorInvalidClangObject)
	}

	if !C.foreman_action_manager_setregisterstore(mgr.cManager, cStoreObj) {
		return fmt.Errorf(errorInvalidClangObject)
	}
	return nil
}

// HasEngine checks whether a script engine of the specified programming language is added.
func (mgr *cgoScriptManager) HasEngine(lang string) bool {
	return bool(C.foreman_action_manager_hasengine(mgr.cManager, C.CString(lang)))
}

// HasMethod checks whether the specified method is added.
func (mgr *cgoScriptManager) HasMethod(method string) bool {
	return bool(C.foreman_action_manager_hasmethod(mgr.cManager, C.CString(method)))
}

// AddMethod adds a new method to the action manager.
func (mgr *cgoScriptManager) AddMethod(method *Method) error {
	if mgr.cManager == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cmethod, err := method.CObject()
	if err != nil {
		return err
	}
	defer C.foreman_action_method_delete(cmethod)

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_action_manager_addmethod(mgr.cManager, cmethod, cerr) {
		return errors.NewWithCObject(cerr).Error()
	}

	return nil
}

// RemoveMethod removes the specified method.
func (mgr *cgoScriptManager) RemoveMethod(name string) error {
	if mgr.cManager == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_action_manager_removemethod(mgr.cManager, C.CString(name), cerr) {
		return errors.NewWithCObject(cerr).Error()
	}

	return nil
}

// RemoveAllMethods removes the all methods.
func (mgr *cgoScriptManager) RemoveAllMethods() error {
	if mgr.cManager == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_action_manager_removeallmethods(mgr.cManager, cerr) {
		return errors.NewWithCObject(cerr).Error()
	}

	return nil
}

// ExecMethod executes a specified method with the parameters.
func (mgr *cgoScriptManager) ExecMethod(name string, params Parameters) (Parameters, error) {
	if mgr.cManager == nil {
		return nil, fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cParams, err := params.CObject()
	if err != nil {
		return nil, err
	}

	results := NewParameters()
	cResults, err := results.CObject()
	if err != nil {
		return nil, err
	}

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	var executeErr error

	executedResult := bool(C.foreman_action_manager_execmethod(mgr.cManager, C.CString(name), cParams, cResults, cerr))
	if !executedResult {
		err := errors.NewWithCObject(cerr)
		if C.foreman_error_isinternalerror(cerr) {
			executeErr = err.Error()
		} else {
			executeErr = fmt.Errorf(cgoScriptManagerExecErrorFormat, name, params.String(), err.Error().Error())
		}
	}

	err = results.SetCObject(cResults)
	if err != nil {
		logging.Error(err.Error())
	}

	logging.Info(cgoScriptManagerExecMessageFormat, name, params.String(), executedResult, results.String())

	return results, executeErr
}

// GetMethod retrun a method with the specified name.
func (mgr *cgoScriptManager) GetMethod(name string) *Method {
	if mgr.cManager == nil {
		return nil
	}

	cMethod := C.foreman_action_manager_getmethod(mgr.cManager, C.CString(name))
	if cMethod == nil {
		return nil
	}

	return NewMethodWithCObject(cMethod)
}

// GetFirstMethod retrun a first method in the manager.
func (mgr *cgoScriptManager) GetFirstMethod() *Method {
	if mgr.cManager == nil {
		return nil
	}

	cMethod := C.foreman_action_manager_getfirstmethod(mgr.cManager)
	if cMethod == nil {
		return nil
	}

	return NewMethodWithCObject(cMethod)
}

// GetNextMethod retrun a first method in the manager.
func (mgr *cgoScriptManager) GetNextMethod(method *Method) *Method {
	if mgr.cManager == nil {
		return nil
	}

	cobj, err := method.CObject()
	if err != nil {
		return nil
	}

	cMethod := C.foreman_action_manager_nextmethod(mgr.cManager, cobj)
	if cMethod == nil {
		return nil
	}

	return NewMethodWithCObject(cMethod)
}
