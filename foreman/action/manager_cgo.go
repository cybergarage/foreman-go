// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -llua -lpython
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

// cgoManager represents an action manager using foreman-cc.
type cgoManager struct {
	cManager unsafe.Pointer
}

// HasEngine checks whether a script engine of the specified programming language is added.
func (mgr *cgoManager) HasEngine(lang string) bool {
	return bool(C.foreman_action_script_manager_hasengine(mgr.cManager, C.CString(lang)))
}

// HasMethod checks whether the specified method is added.
func (mgr *cgoManager) HasMethod(method string) bool {
	return bool(C.foreman_action_script_manager_hasmethod(mgr.cManager, C.CString(method)))
}

// AddMethod adds a new method to the action manager.
func (mgr *cgoManager) AddMethod(method *Method) error {
	if mgr.cManager == nil {
		return fmt.Errorf(errors.ErrorClangObjectNotInitialized)
	}

	cmethod, err := method.CObject()
	if err != nil {
		return err
	}

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_action_script_manager_addmethod(mgr.cManager, cmethod, cerr) {
		err = errors.NewWithCObject(cerr)
		return err
	}

	return nil
}

// ExecMethod executes a specified method with the parameters.
func (mgr *cgoManager) ExecMethod(name string, params Parameters, results Parameters) error {
	cparams, err := params.CObject()
	if err != nil {
		return err
	}

	cresults, err := results.CObject()
	if err != nil {
		return err
	}

	cerr := C.foreman_error_new()
	defer C.foreman_error_delete(cerr)

	if !C.foreman_action_script_manager_execmethod(mgr.cManager, C.CString(name), cparams, cresults, cerr) {
		err = errors.NewWithCObject(cerr)
		return err
	}

	err = results.SetCObject(cresults)
	if err != nil {
		return err
	}

	return nil
}
