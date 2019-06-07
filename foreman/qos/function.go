// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"github.com/cybergarage/foreman-go/foreman/kb"
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/foreman-go/foreman/register"
	"github.com/cybergarage/foreman-go/foreman/registry"
)

// Function represents a function operand object.
type Function interface {
	kb.Function
	SetMetricManager(mgr *metric.Manager)
	SetRegisterManager(mgr *register.Manager)
	SetRegistryManager(mgr *registry.Manager)
}
