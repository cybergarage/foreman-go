// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"encoding/base64"
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
	m.Language = lang
	if enc == ActionEncodingBase64 {
		code, err := base64.StdEncoding.DecodeString(code)
		if err != nil {
			return nil, errors.NewErrorWithError(err)
		}
		m.Code = code
	} else {
		m.Code = []byte(code)
	}

	err := mgr.AddMethod(m)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return nil, nil
}

func (mgr *Manager) executeSelectAction(q fql.Query) (interface{}, *errors.Error) {
	ope, whereName, ok := q.GetConditionByColumn(fql.QueryColumnName)
	if ok {
		if ope.GetType() != fql.OperatorTypeEQ {
			ok = false
		}
	}

	var methods []interface{}

	method := mgr.GetFirstMethod()
	for method != nil {
		method = mgr.GetNextMethod(method)
		if ok {
			if method.Name != whereName {
				continue
			}
		}
		var methodMap map[string]interface{}
		methodMap[ActionColumnName] = method.Name
		methodMap[ActionColumnLanguage] = method.Language
		methodMap[ActionColumnCode] = base64.StdEncoding.EncodeToString(method.Code)
		methodMap[ActionColumnEncoding] = ActionEncodingBase64

		methods = append(methods, methodMap)
	}

	var actionMap map[string]interface{}
	actionMap[ActionColumnMethods] = methods
	actionContainer := map[string]interface{}{}
	actionContainer[strings.ToLower(fql.QueryTargetAction)] = actionMap

	return actionContainer, nil
}

func (mgr *Manager) executeDeleteAction(q fql.Query) (interface{}, *errors.Error) {
	ope, whereName, ok := q.GetConditionByColumn(fql.QueryColumnName)
	if ok {
		if ope.GetType() != fql.OperatorTypeEQ {
			ok = false
		}
	}

	// Delete only a specified rule

	if ok {
		err := mgr.RemoveMethod(whereName)
		if err != nil {
			return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
		}
		return nil, nil
	}

	// Delete all rules

	err := mgr.RemoveAllMethods()
	if err != nil {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
	}

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
