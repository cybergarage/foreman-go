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

// FindRelatedRules returns all rules of the the specified name.
func (qos *QoS) FindRelatedRules(q *Query) ([]*Rule, error) {
	name := q.Target
	qosRules := make([]*Rule, 0)

	for _, rule := range qos.Rules {
		for _, clause := range rule.GetClauses() {
			for _, formula := range clause.GetFormulas() {
				for _, operand := range formula.GetOperands() {
					switch operand.(type) {
					case kb.Variable:
						{
							v, _ := operand.(kb.Variable)
							if v.GetName() != name {
								continue
							}
						}
					case kb.Function:
						{
							fn, _ := operand.(kb.Function)
							if !fn.HasParameter(name) {
								continue
							}
						}
					default:
						continue
					}
					qosRule, ok := rule.(*Rule)
					if !ok {
						continue
					}
					qosRules = append(qosRules, qosRule)

				}
			}
		}
	}

	return qosRules, nil
}

// FindRelatedFormulas returns all QoS metrics of the the specified name.
func (qos *QoS) FindRelatedFormulas(q *Query) ([]*Formula, error) {
	name := q.Target
	qoSFormulas := make([]*Formula, 0)

	for _, rule := range qos.Rules {
		for _, clause := range rule.GetClauses() {
			for _, formula := range clause.GetFormulas() {
				for _, operand := range formula.GetOperands() {
					v, ok := operand.(kb.Variable)
					if !ok {
						continue
					}
					if v.GetName() != name {
						continue
					}
					qosFormula, ok := formula.(*Formula)
					if !ok {
						continue
					}
					qoSFormulas = append(qoSFormulas, qosFormula)
				}
			}
		}
	}

	return qoSFormulas, nil
}
