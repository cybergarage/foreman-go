// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lm -lstdc++ -lsqlite3 -lfolly -lgflags -lglog -llua -lpython
import "C"

import (
	"runtime"
	"unsafe"
)

func newManagerWithInterface(mgrImpl Scripting) *Manager {
	mgr := &Manager{
		Scripting: mgrImpl,
	}
	return mgr
}

func newManagerWithCObject(cObject unsafe.Pointer) *Manager {
	mgrImp := &cgoManager{}
	mgrImp.cManager = cObject
	runtime.SetFinalizer(mgrImp, managerFinalizer)
	return newManagerWithInterface(mgrImp)
}

func managerFinalizer(self *cgoManager) {
	if self.cManager != nil {
		if C.foreman_action_script_manager_delete(self.cManager) {
			self.cManager = nil
		}
	}
}

// NewManager returns a new action manager.
func NewManager() *Manager {
	mgr := newManagerWithCObject(C.foreman_action_script_manager_new())
	return mgr
}
