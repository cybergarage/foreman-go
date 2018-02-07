// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"strings"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
)

func (mgr *Manager) executeInsertQuery(q fql.Query) (interface{}, *errors.Error) {
	values, ok := q.GetValues()
	if !ok || len(values) < 2 {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidValues)
	}

	name := values[0].String()
	qos := values[1].String()

	rule, err := mgr.ParseQoSString(qos)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	err = rule.SetName(name)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	err = mgr.SetRule(rule)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return nil, nil
}

func (mgr *Manager) getNameCondition(q fql.Query) (string, bool) {
	wheres, ok := q.GetConditions()
	if !ok || len(wheres) < 1 {
		return "", false
	}
	where := wheres[0]
	if where.GetLeftOperand() != fql.QueryColumnName {
		return "", false
	}
	name := where.GetRightOperand()
	if !mgr.RemoveRule(name) {
		return "", false
	}

	return name, true
}

func (mgr *Manager) executeSelectQuery(q fql.Query) (interface{}, *errors.Error) {
	whereName, ok := mgr.getNameCondition(q)

	var ruleMap map[string]string
	for _, rule := range mgr.GetRules() {
		ruleName := rule.GetName()
		if ok {
			if ruleName != whereName {
				continue
			}
		}
		ruleMap[ruleName] = rule.String()
	}

	qosContainer := map[string]interface{}{}
	qosContainer[strings.ToLower(fql.QueryTargetQos)] = ruleMap

	return qosContainer, nil
}

func (mgr *Manager) executeDeleteQuery(q fql.Query) (interface{}, *errors.Error) {
	// Delete only a specified rule

	whereName, ok := mgr.getNameCondition(q)
	if ok {
		if !mgr.RemoveRule(whereName) {
			return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
		}
		return nil, nil
	}

	// Delete all rules

	err := mgr.Clear()
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return nil, nil
}

// ExecuteQuery must return the result as a standard array or map.
func (mgr *Manager) ExecuteQuery(q fql.Query) (interface{}, *errors.Error) {
	switch q.GetType() {
	case fql.QueryTypeInsert:
		return mgr.executeInsertQuery(q)
	case fql.QueryTypeSelect:
		return mgr.executeSelectQuery(q)
	case fql.QueryTypeDelete:
		return mgr.executeDeleteQuery(q)
	}

	return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
}
