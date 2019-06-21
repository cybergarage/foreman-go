// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package function

import (
	"fmt"
	"math"

	"github.com/cybergarage/foreman-go/foreman/kb"
)

// Abs represents a abs function .
type Abs struct {
	*Function
}

// NewAbs returns a abs function instance.
func NewAbs() *Abs {
	fn := &Abs{
		Function: NewFunctionWithName("ABS"),
	}
	fn.SetExecutor(fn)
	return fn
}

// Execute returns the operand value with the specified parameters.
func (fn *Abs) Execute(params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, fn.GetParameterError(params)
	}

	var param interface{}
	switch params[0].(type) {
	case kb.Variable:
		v, _ := params[0].(kb.Variable)
		val, err := v.GetValue()
		if err != nil {
			return nil, err
		}
		param = val
	default:
		param = params[0]
	}

	switch param.(type) {
	case float64:
		val, _ := param.(float64)
		return math.Abs(val), nil
	case float32:
		val, _ := param.(float32)
		return math.Abs(float64(val)), nil
	case int:
		val, _ := param.(int)
		return math.Abs(float64(val)), nil
	case int8:
		val, _ := param.(int8)
		return math.Abs(float64(val)), nil
	case int16:
		val, _ := param.(int16)
		return math.Abs(float64(val)), nil
	case int64:
		val, _ := param.(int64)
		return math.Abs(float64(val)), nil
	case int32:
		val, _ := param.(int32)
		return math.Abs(float64(val)), nil
	}

	return nil, fmt.Errorf("Invalid parameter type : %T(%v)", params[0], params[0])
}
