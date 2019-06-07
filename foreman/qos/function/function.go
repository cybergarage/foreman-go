// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package function

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/kb"
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/foreman-go/foreman/register"
	"github.com/cybergarage/foreman-go/foreman/registry"
)

// Function represents a function instance for Knowledge base.
type Function struct {
	*kb.BaseFunction
	MetricManager   *metric.Manager
	RegisterManager *register.Manager
	RegistryManager *registry.Manager
}

// NewFunctionWithName returns a new function instance with the specified name.
func NewFunctionWithName(name string) *Function {
	fn := &Function{
		BaseFunction: kb.NewFunctionWithName(name),
	}
	return fn
}

// SetMetricManager sets a metric manager.
func (fn *Function) SetMetricManager(mgr *metric.Manager) {
	fn.MetricManager = mgr
}

// GetMetricManager gets a metric manager.
func (fn *Function) GetMetricManager() *metric.Manager {
	return fn.MetricManager
}

// SetRegisterManager sets a register manager.
func (fn *Function) SetRegisterManager(mgr *register.Manager) {
	fn.RegisterManager = mgr
}

// GetRegisterManager gets a register manager.
func (fn *Function) GetRegisterManager() *register.Manager {
	return fn.RegisterManager
}

// SetRegistryManager sets a metric manager.
func (fn *Function) SetRegistryManager(mgr *registry.Manager) {
	fn.RegistryManager = mgr
}

// GetRegistryManager gets a metric manager.
func (fn *Function) GetRegistryManager() *registry.Manager {
	return fn.RegistryManager
}

// GetParameterError returns an error for invalid parameters
func (fn *Function) GetParameterError(params []interface{}) error {
	return fmt.Errorf(errorInvalidParameters, fn.ExpressionWithParameters(params))
}
