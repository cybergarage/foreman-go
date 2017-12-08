// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// Scripting represents an abstract interface of the script manager
type Scripting interface {
	HasEngine(method string) bool

	AddMethod(method *Method) error
	HasMethod(method string) bool
	ExecMethod(name string, params Parameters) (Parameters, error)
}

// ScriptManager represents an interface of the script manager
type ScriptManager struct {
	Scripting
}
