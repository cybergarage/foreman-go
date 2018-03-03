// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package action

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	testActionScriptDirectory = "../../test/actions/"
)

const (
	errorMethodNotFound  = "Method (%s) is not found"
	errorEngineNotFound  = "Script Engine (%s) is not found"
	errorObjectNotEquals = "Object does not equal (%v != %v)"
)

const (
	testEchoMethod = "test_echo"
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
	testSetRegisterMethod = "test_setregister"
	testGetRegisterMethod = "test_getregister"
)

func executeTestRegisterMethods(mgr *ScriptManager) (Parameters, error) {
	rand := rand.New(rand.NewSource(99))

	regName := fmt.Sprintf("%d", rand.Int31())
	regValue := fmt.Sprintf("%f", rand.Float64())

	// test_setregister

	inParams := NewParameters()
	inParam := NewParameterFromString(regName, regValue)
	inParams.AddParameter(inParam)

	_, err := mgr.ExecMethod(testSetRegisterMethod, inParams)
	if err != nil {
		return nil, err
	}

	// test_getregister

	outParams := NewParameters()

	results, err := mgr.ExecMethod(testGetRegisterMethod, outParams)
	if err != nil {
		return nil, err
	}

	// Check result

	if !inParams.Equals(results) {
		err = fmt.Errorf(errorObjectNotEquals, inParams, results)
		return nil, err
	}

	return results, nil
}
