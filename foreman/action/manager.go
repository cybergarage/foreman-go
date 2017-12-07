// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package action provides scripting interfaces.
package action

// Scripting represents an abstract interface of action script manager
type Scripting interface {
	HasEngine(method string) bool

	AddMethod(method *Method) error
	HasMethod(method string) bool
	ExecMethod(name string, params Parameters, results Parameters) error
}

// Manager represents an action manager.
type Manager struct {
	Scripting
}
