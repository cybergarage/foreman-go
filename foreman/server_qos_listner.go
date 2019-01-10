// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/action"
	"github.com/cybergarage/foreman-go/foreman/kb"
	"github.com/cybergarage/foreman-go/foreman/logging"
)

type qosRuleSource struct {
	//action.RouteSource Disabled to conflict Rule::GetName()
	kb.Rule
}

// newQosRuleSourceWithRule create a source object with the specified rule
func newQosRuleSourceWithRule(rule kb.Rule) action.RouteSource {
	src := &qosRuleSource{
		Rule: rule,
	}
	return src
}

// newActionEventWithUnsatisfiedQoSRule create an action event with the specified unsatisfied rule
func newActionEventWithUnsatisfiedQoSRule(rule kb.Rule) (*action.Event, error) {
	e := action.NewEventWithSource(newQosRuleSourceWithRule(rule))

	// Set only unsatisfied variables to the event parameter

	var lastErr error

	for _, clause := range rule.GetClauses() {
		isSatisfied, err := clause.IsSatisfied()
		if err != nil {
			lastErr = err
			continue
		}
		if isSatisfied {
			continue
		}
		for _, formula := range clause.GetFormulas() {
			isSatisfied, err := formula.IsSatisfied()
			if err != nil {
				lastErr = err
				continue
			}
			if isSatisfied {
				continue
			}

			ver := formula.GetVariable()
			val, err := ver.GetValue()
			if err != nil {
				lastErr = err
				continue
			}

			param, err := action.NewParameterFromInterface(ver.GetName(), val)
			if err != nil {
				lastErr = err
				continue
			}

			err = e.AddParameter(param)
			if err != nil {
				lastErr = err
				continue
			}
		}
	}

	return e, lastErr
}

// postQosUnsatisfiedEvent posts a QoS event
func (server *Server) postQosUnsatisfiedEvent(rule kb.Rule) {
	e, err := newActionEventWithUnsatisfiedQoSRule(rule)

	if err != nil {
		logging.Warn(err.Error())
	}

	server.actionMgr.PostEvent(e)
}

// RuleSatisfied is a listener for kb.Rule
func (server *Server) RuleSatisfied(rule kb.Rule) {
	logging.Trace("Satisfied : %s", rule.String())
}

// RuleUnsatisfied is a listener for kb.Rule
func (server *Server) RuleUnsatisfied(rule kb.Rule) {
	logging.Info("Unsatisfied : %s", rule.String())

	server.postQosUnsatisfiedEvent(rule)
}
