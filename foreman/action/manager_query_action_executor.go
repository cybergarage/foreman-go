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
	ope, whereName, hasWhere := q.GetConditionByColumn(fql.QueryColumnName)
	if hasWhere {
		if ope.GetType() != fql.OperatorTypeEQ {
			return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
		}
	}

	var methods []interface{}

	method := mgr.GetFirstMethod()
	for method != nil {
		if hasWhere {
			if method.Name != whereName {
				method = mgr.GetNextMethod(method)
				continue
			}
		}

		var methodMap map[string]interface{}
		methodMap[ActionColumnName] = method.Name
		methodMap[ActionColumnLanguage] = method.Language
		methodMap[ActionColumnCode] = base64.StdEncoding.EncodeToString(method.Code)
		methodMap[ActionColumnEncoding] = ActionEncodingBase64
		methods = append(methods, methodMap)

		method = mgr.GetNextMethod(method)
	}

	var actionMap map[string]interface{}
	actionMap[ActionColumnMethods] = methods
	actionContainer := map[string]interface{}{}
	actionContainer[strings.ToLower(fql.QueryTargetAction)] = actionMap

	return actionContainer, nil
}

func (mgr *Manager) executeDeleteAction(q fql.Query) (interface{}, *errors.Error) {
	ope, whereName, hasWhere := q.GetConditionByColumn(fql.QueryColumnName)
	if hasWhere {
		if ope.GetType() != fql.OperatorTypeEQ {
			return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
		}
	}

	// Remove only a specified rule

	if hasWhere {
		err := mgr.RemoveMethod(whereName)
		if err != nil {
			return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
		}
		return nil, nil
	}

	// Remove all rules

	err := mgr.RemoveAllMethods()
	if err != nil {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
	}

	return nil, nil
}

func (mgr *Manager) executeExecuteAction(q fql.Query) (interface{}, *errors.Error) {
	ope, methodName, hasName := q.GetConditionByColumn(fql.QueryColumnName)
	if hasName {
		if ope.GetType() != fql.OperatorTypeEQ {
			return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
		}
	}

	columns, ok := q.GetColumns()
	if !ok {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidColumns)
	}

	values, ok := q.GetValues()
	if !ok {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidValues)
	}

	if len(columns) != len(values) {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidValues)
	}

	params := NewParameters()
	for n, column := range columns {
		value := values[n]
		param := NewParameterFromString(column.String(), value.String())
		params.AddParameter(param)
	}

	results, err := mgr.ExecMethod(methodName, params)
	if err != nil {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
	}

	var resultMap map[string]interface{}
	for name, result := range results {
		resultMap[name] = result.GetValue()
	}

	return resultMap, nil
}
