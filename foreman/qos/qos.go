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

// QoS includes all QoS rules.
type QoS struct {
	kb.KnowledgeBase
	MetricManager   *metric.Manager
	RegisterManager *register.Manager
	RegistryManager *registry.Manager
}

// NewQoS returns a new null object.
func NewQoS() *QoS {
	qos := &QoS{
		KnowledgeBase: *kb.NewKnowledgeBase(),
	}
	return qos
}

// SetMetricManager sets a metric manager.
func (qos *QoS) SetMetricManager(mgr *metric.Manager) {
	qos.MetricManager = mgr
}

// GetMetricManager gets a metric manager.
func (qos *QoS) GetMetricManager() *metric.Manager {
	return qos.MetricManager
}

// SetRegisterManager sets a register manager.
func (qos *QoS) SetRegisterManager(mgr *register.Manager) {
	qos.RegisterManager = mgr
}

// GetRegisterManager gets a register manager.
func (qos *QoS) GetRegisterManager() *register.Manager {
	return qos.RegisterManager
}

// SetRegistryManager sets a metric manager.
func (qos *QoS) SetRegistryManager(mgr *registry.Manager) {
	qos.RegistryManager = mgr
}

// GetRegistryManager gets a metric manager.
func (qos *QoS) GetRegistryManager() *registry.Manager {
	return qos.RegistryManager
}

// ParseString parses a specified QoS string.
func (qos *QoS) ParseString(qosString string) (kb.Rule, error) {
	return qos.ParseRuleString(qos, qosString)
}
