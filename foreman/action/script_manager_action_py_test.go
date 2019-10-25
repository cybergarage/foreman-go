// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package action

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const (
	testActionPythonFileExt = ".py"
)

func loadPythonEngineMethod(mgr *ScriptManager, methodName string) error {
	methodCode, err := ioutil.ReadFile(testActionScriptDirectory + methodName + testActionPythonFileExt)
	if err != nil {
		return err
	}

	method := NewPythonMethod()
	method.Name = methodName
	method.Code = methodCode

	err = mgr.AddMethod(method)
	if err != nil {
		return err
	}

	if !mgr.HasMethod(method.Name) {
		err = fmt.Errorf(errorMethodNotFound, method.Name)
		return err
	}

	return nil
}

func testPythonEngineEcho(t *testing.T, mgr *ScriptManager) {
	err := loadPythonEngineMethod(mgr, testEchoMethod)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = executeTestEchoMethod(mgr)
	if err != nil {
		t.Error(err)
		return
	}
}

func testPythonEngineRegister(t *testing.T, mgr *ScriptManager) {
	err := loadPythonEngineMethod(mgr, testSetRegisterMethod)
	if err != nil {
		t.Error(err)
		return
	}

	err = loadPythonEngineMethod(mgr, testGetRegisterMethod)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = executeTestRegisterMethods(mgr)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestPythonEngine(t *testing.T) {
	mgr := newTestScriptEngine()

	if !mgr.HasEngine(ActionLanguagePython) {
		t.Logf(errorEngineNotFound, ActionLanguagePython)
		return
	}

	testPythonEngineEcho(t, mgr)
	testPythonEngineRegister(t, mgr)
}
