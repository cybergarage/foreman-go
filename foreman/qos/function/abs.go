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
func NewAbs() kb.Function {
	fn := &Abs{
		Function: NewFunctionWithName("ABS"),
	}
	return fn
}

// Execute returns the operand value with the specified parameters.
func (fn *Abs) Execute(params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, fn.GetParameterError(params)
	}
	switch params[0].(type) {
	case float64:
		val, _ := params[0].(float64)
		return math.Abs(val), nil
	case float32:
		val, _ := params[0].(float32)
		return math.Abs(float64(val)), nil
	case int:
		val, _ := params[0].(int)
		return math.Abs(float64(val)), nil
	case int8:
		val, _ := params[0].(int8)
		return math.Abs(float64(val)), nil
	case int16:
		val, _ := params[0].(int16)
		return math.Abs(float64(val)), nil
	case int64:
		val, _ := params[0].(int64)
		return math.Abs(float64(val)), nil
	case int32:
		val, _ := params[0].(int32)
		return math.Abs(float64(val)), nil
	}
	return nil, fmt.Errorf("Invalid parameter type : %T(%v)", params[0], params[0])
}
