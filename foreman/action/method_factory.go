// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// NewMethodWithName returns a new method with the specified name.
func NewMethodWithName(name string) *Method {
	m := NewMethod()
	m.Name = name
	return m
}

// NewMethodWithLanguage returns a new method with the specified language.
func NewMethodWithLanguage(lang string) *Method {
	m := NewMethod()
	m.Language = lang
	return m
}

// NewPythonMethod returns a new Python method.
func NewPythonMethod() *Method {
	return NewMethodWithLanguage(ActionLanguagePython)
}
