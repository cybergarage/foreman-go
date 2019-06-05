// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package function

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/kb"
)

// Function represents a function instance for Knowledge base.
type Function struct {
	*kb.BaseFunction
}

// NewFunctionWithName returns a new function instance with the specified name.
func NewFunctionWithName(name string) *Function {
	fn := &Function{
		BaseFunction: kb.NewFunctionWithName(name),
	}
	return fn
}

// GetParameterError returns an error for invalid parameters
func (fn *Function) GetParameterError(params []interface{}) error {
	return fmt.Errorf(errorInvalidParameters, fn.ExpressionWithParameters(params))
}
