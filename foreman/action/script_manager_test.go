// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package action

import "testing"
import "fmt"

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

func echoExecutionTest(t *testing.T, mgr *ScriptManager, method *Method, params Parameters) (Parameters, error) {
	err := mgr.AddMethod(method)
	if err != nil {
		t.Error(err)
		return nil, err
	}

	if !mgr.HasMethod(method.Name) {
		err = fmt.Errorf(errorMethodNotFound, method.Name)
		t.Error(err)
		return nil, err
	}

	results, err := mgr.ExecMethod(method.Name, params)
	if err != nil {
		t.Error(err)
		return nil, err
	}

	if !params.Equals(results) {
		err = fmt.Errorf(errorObjectNotEquals, params, results)
		t.Error(err)
		return nil, err
	}

	return results, nil
}

func TestPythonEngine(t *testing.T) {
	mgr := NewScriptManager()

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

	_, err := echoExecutionTest(t, mgr, method, params)
	if err != nil {
		t.Error(err)
	}
}