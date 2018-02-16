// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package kb

import (
	"testing"
)

type TestVariable struct {
	Variable
	Name  string
	Value float64
}

func (v *TestVariable) GetName() string {
	return v.Name
}

func (v *TestVariable) SetValue(value interface{}) error {
	floatValue, ok := value.(float64)
	if !ok {
		return nil
	}
	v.Value = floatValue
	return nil
}

func (v *TestVariable) GetValue() (interface{}, error) {
	return v.Value, nil
}

func newTestVariableWithName(name string) *TestVariable {
	return &TestVariable{
		Name:  name,
		Value: 0.0,
	}
}

func newTestVariable() *TestVariable {
	return newTestVariableWithName("")
}

func TestNewVariable(t *testing.T) {
	var v Variable
	v = newTestVariable()
	v.GetName()
}
