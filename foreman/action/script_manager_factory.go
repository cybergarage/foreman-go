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

func newScriptManagerWithInterface(mgrImpl Scripting) *ScriptManager {
	mgr := &ScriptManager{
		Scripting: mgrImpl,
	}
	return mgr
}

func newScriptManagerWithCObject(cObject unsafe.Pointer) *ScriptManager {
	mgrImp := &cgoScriptManager{}
	mgrImp.cManager = cObject
	runtime.SetFinalizer(mgrImp, scriptManagerFinalizer)
	return newScriptManagerWithInterface(mgrImp)
}

func scriptManagerFinalizer(self *cgoScriptManager) {
	if self.cManager != nil {
		if C.foreman_action_manager_delete(self.cManager) {
			self.cManager = nil
		}
	}
}

// NewScriptManager returns a new script manager.
func NewScriptManager() *ScriptManager {
	mgr := newScriptManagerWithCObject(C.foreman_action_manager_new())
	return mgr
}
