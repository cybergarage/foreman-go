// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"testing"
)

func TestNewIntegerParameter(t *testing.T) {
	param := NewParameterFromInteger("", 0)
	if !param.IsInteger() {
		t.Errorf(errorInvalidType, param.Type, IntegerType)
	}
	_, err := param.GetInteger()
	if err != nil {
		t.Error(err)
	}
}

func TestNewRealParameter(t *testing.T) {
	param := NewParameterFromReal("", 0)
	if !param.IsReal() {
		t.Errorf(errorInvalidType, param.Type, RealType)
	}
	_, err := param.GetReal()
	if err != nil {
		t.Error(err)
	}
}

func TestNewBoolParameter(t *testing.T) {
	param := NewParameterFromBool("", true)
	if !param.IsBool() {
		t.Errorf(errorInvalidType, param.Type, BoolType)
	}
	_, err := param.GetBool()
	if err != nil {
		t.Error(err)
	}
}

func TestNewStringParameter(t *testing.T) {
	param := NewParameterFromString("", "")
	if !param.IsString() {
		t.Errorf(errorInvalidType, param.Type, StringType)
	}
	_, err := param.GetString()
	if err != nil {
		t.Error(err)
	}
}
