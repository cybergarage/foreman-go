// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/kb"
)

// InitFunctions is called when QoS manager is generated
func (qos *QoS) InitFunctions() {
}

// CreateFunctionOperand  is an interface method of kb.Factory
func (qos *QoS) CreateFunctionOperand(name string) (kb.Function, error) {
	fn, ok := qos.GetFunction(name)
	if !ok {
		return nil, fmt.Errorf(errorUnknownFunction, name)
	}
	return fn, nil
}
