// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"strings"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
)

func (mgr *Manager) executeInsertAction(q fql.Query) (interface{}, *errors.Error) {
	values, ok := q.GetValues()
	if !ok || len(values) < 4 {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidValues)
	}

	name := values[0].String()
	lang := values[1].String()
	code := values[2].String()
	enc := values[3].String()

	m := NewMethodWithLanguage(lang)
	m.Name = name
	m.Code = []byte(code)
	m.Encoding = EncodingNone
	if enc == ActionEncodingBase64 {
		m.Encoding = EncodingBase64
	}

	err := mgr.AddMethod(m)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return nil, nil
}

func (mgr *Manager) executeSelectAction(q fql.Query) (interface{}, *errors.Error) {
	ope /*whereName*/, _, ok := q.GetConditionByColumn(fql.QueryColumnName)
	if ok {
		if ope.GetType() != fql.OperatorTypeEQ {
			ok = false
		}
	}

	var actionMap map[string]interface{}
	/*
		FIXME : Update foraman-cc
		for _, rule := range mgr.GetRules() {
			ruleName := rule.GetName()
			if ok {
				if ruleName != whereName {
					continue
				}
			}
			actioMap[ruleName] = rule.String()
		}
	*/

	actionContainer := map[string]interface{}{}
	actionContainer[strings.ToLower(fql.QueryTargetAction)] = actionMap

	return actionContainer, nil
}

func (mgr *Manager) executeDeleteAction(q fql.Query) (interface{}, *errors.Error) {
	ope /*whereName*/, _, ok := q.GetConditionByColumn(fql.QueryColumnName)
	if ok {
		if ope.GetType() != fql.OperatorTypeEQ {
			ok = false
		}
	}

	// Delete only a specified rule

	/*
		FIXME : Update foraman-cc
		if ok {
			if !mgr.De. RemoveRule(whereName) {
				return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
			}
			return nil, nil
		}
	*/

	// Delete all rules

	/*
		FIXME : Update foraman-cc
		err := mgr.Clear()
		if err != nil {
			return nil, errors.NewErrorWithError(err)
		}
	*/

	return nil, nil
}

// ExecuteQuery must return the result as a standard array or map.
func (mgr *Manager) ExecuteQuery(q fql.Query) (interface{}, *errors.Error) {
	targetObj, ok := q.GetTarget()
	if !ok {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryEmptyTarget)
	}
	switch targetObj.String() {
	case fql.QueryTargetAction:
		switch q.GetType() {
		case fql.QueryTypeInsert:
			return mgr.executeInsertAction(q)
		case fql.QueryTypeSelect:
			return mgr.executeSelectAction(q)
		case fql.QueryTypeDelete:
			return mgr.executeDeleteAction(q)
		}
	case fql.QueryTargetRoute:
		// Not yet implemented
	}

	return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
}
