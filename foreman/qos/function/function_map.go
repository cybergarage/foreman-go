// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package function

import (
	"strings"

	"github.com/cybergarage/foreman-go/foreman/kb"
)

// FunctionMap represents a function map.
type FunctionMap map[string]kb.Function

// NewFunctionMap returns a new function map.
func NewFunctionMap() FunctionMap {
	return map[string]kb.Function{}
}

// AddFunction adds the specified function.
func (m FunctionMap) AddFunction(fn kb.Function) {
	name := fn.GetName()
	upperName := strings.ToUpper(name)
	m[upperName] = fn
}

// GetFunction returns the specified function.
func (m FunctionMap) GetFunction(name string) (kb.Function, bool) {
	upperName := strings.ToUpper(name)
	fn, ok := m[upperName]
	return fn, ok
}
