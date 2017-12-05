// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package action provides scripting interfaces.
package action

// Manager represents an abstract interface of action manager
type Manager interface {
	AddMethod(method *Method) error
	ExecMethod(name string, params Parameters, results Parameters) error
}
