// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package action

import "testing"

const (
	echoMethod     = "echo"
	echoParamName  = "message"
	echoParamValue = "hello"
	echoPythonCode = "def echo(params,results):\n" +
		"  for key, value in params.iteritems():\n" +
		"    results[key] = value\n" +
		"  return True\n"
)

const (
	errorMethodNotFound  = "Method (%s) is not found"
	errorEngineNotFound  = "Script Engine (%s) is not found"
	errorObjectNotEquals = "Object does not equal (%v != %v)"
)

func echoExecutionTest(t *testing.T, mgr *Manager, method *Method, params Parameters, results Parameters) error {
	err := mgr.AddMethod(method)
	if err != nil {
		t.Error(err)
	}

	if !mgr.HasMethod(method.Name) {
		t.Errorf(errorMethodNotFound, method.Name)
		return nil
	}

	err = mgr.ExecMethod(method.Name, params, results)
	if err != nil {
		t.Error(err)
	}

	if !params.Equals(results) {
		t.Errorf(errorObjectNotEquals, params, results)
	}

	return nil
}

func TestPythonEngine(t *testing.T) {
	mgr := NewManager()

	if !mgr.HasEngine(PythonEngine) {
		t.Errorf(errorEngineNotFound, PythonEngine)
		return
	}

	method := NewPythonMethod()
	method.Name = echoMethod
	method.Code = []byte(echoPythonCode)

	params := NewParameters()
	param := NewParameterFromString(echoParamName, echoParamValue)
	params.AddParameter(param)

	results := NewParameters()

	err := echoExecutionTest(t, mgr, method, params, results)
	if err != nil {
		t.Error(err)
	}
}
