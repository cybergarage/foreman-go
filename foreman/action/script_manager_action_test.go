// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package action

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cybergarage/foreman-go/foreman/register"
	"github.com/cybergarage/foreman-go/foreman/registry"
)

const (
	testActionScriptDirectory = "../../test/actions/"
)

const (
	errorEngineNotFound  = "Script Engine (%s) is not found"
	errorObjectNotEquals = "Object does not equal (%v != %v)"
)

func newTestScriptEngine() *ScriptManager {
	mgr := NewScriptManager()
	mgr.SetRegisterStore(register.NewStore())
	mgr.SetRegistryStore(registry.NewStore())
	return mgr
}

const (
	testEchoMethod = "echo"
)

func executeTestEchoMethod(mgr *ScriptManager) (Parameters, error) {
	echoParamName := fmt.Sprintf("%d", time.Now().Unix())
	echoParamValue := fmt.Sprintf("%d", time.Now().Unix())

	params := NewParameters()
	param := NewParameterFromString(echoParamName, echoParamValue)
	params.AddParameter(param)

	results, err := mgr.ExecMethod(testEchoMethod, params)
	if err != nil {
		return nil, err
	}

	if !params.Equals(results) {
		err = fmt.Errorf(errorObjectNotEquals, params, results)
		return nil, err
	}

	return results, nil
}

const (
	testSetRegisterMethod = "setregister"
	testGetRegisterMethod = "getregister"
)

func executeTestRegisterMethods(mgr *ScriptManager) (Parameters, error) {
	rand := rand.New(rand.NewSource(99))

	regName := fmt.Sprintf("%d", rand.Int31())
	regValue := fmt.Sprintf("%f", rand.Float64())

	// test_setregister

	setParams := NewParameters()
	setParams.AddParameter(NewParameterFromString(regName, regValue))

	_, err := mgr.ExecMethod(testSetRegisterMethod, setParams)
	if err != nil {
		return nil, err
	}

	// test_getregister

	getParams := NewParameters()
	getParams.AddParameter(NewParameterFromString(regName, ""))

	getResults, err := mgr.ExecMethod(testGetRegisterMethod, getParams)
	if err != nil {
		return nil, err
	}

	// Check result

	if !setParams.Equals(getResults) {
		err = fmt.Errorf(errorObjectNotEquals, setParams, getResults)
		return nil, err
	}

	return getResults, nil
}
