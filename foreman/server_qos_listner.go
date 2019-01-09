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

// RuleSatisfied is a listener for kb.Rule
func newQosRuleSourceWithRule(rule kb.Rule) action.RouteSource {
	src := &qosRuleSource{
		Rule: rule,
	}
	return src
}

// RuleSatisfied is a listener for kb.Rule
func (server *Server) RuleSatisfied(rule kb.Rule) {
	logging.Trace("Satisfied : %s", rule.String())
}

// RuleUnsatisfied is a listener for kb.Rule
func (server *Server) RuleUnsatisfied(rule kb.Rule) {
	logging.Info("Unsatisfied : %s", rule.String())

	e := action.NewEventWithSource(newQosRuleSourceWithRule(rule))

	// Set only unsatisfied variables to the event parameter

	for _, clause := range rule.GetClauses() {
		for _, formula := range clause.GetFormulas() {
			isSatisfied, err := formula.IsSatisfied()
			if isSatisfied || (err != nil) {
				continue
			}

			ver := formula.GetVariable()
			val, err := ver.GetValue()
			if err != nil {
				logging.Warn(err.Error())
				continue
			}

			param, err := action.NewParameterFromInterface(ver.GetName(), val)
			if err != nil {
				logging.Warn(err.Error())
				continue
			}

			err = e.AddParameter(param)
			if err != nil {
				logging.Warn(err.Error())
				continue
			}
		}
	}

	server.actionMgr.PostEvent(e)
}
