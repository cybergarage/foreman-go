// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"strings"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

// GetJSONExportMethodPath returns a JSON path for methods
func GetJSONExportMethodPath() string {
	return json.PathSep + strings.Join([]string{getJSONExportActionMapName(), ActionColumnMethods}, json.PathSep)
}

func (mgr *Manager) executeInsertMethod(q fql.Query) (interface{}, *errors.Error) {
	values, ok := q.GetValues()
	if !ok || len(values) < 4 {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidValues)
	}

	name := values[0].String()
	lang := values[1].String()
	code := values[2].String()
	enc := values[3].String()

	err := mgr.CreateMethod(name, lang, code, enc)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return nil, nil
}

func (mgr *Manager) executeSelectMethod(q fql.Query) (interface{}, *errors.Error) {
	ope, whereName, hasWhere := q.GetConditionByColumn(fql.QueryColumnName)
	if hasWhere {
		if ope.GetType() != fql.OperatorTypeEQ {
			return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
		}
	}

	var methods interface{}
	var err error

	if hasWhere {
		methods, err = mgr.exportMethodJSONObjectWithName(whereName)
	} else {
		methods, err = mgr.exportMethodJSONObject()
	}

	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	methodMap := map[string]interface{}{}
	methodMap[ActionColumnMethods] = methods

	actionContainer := map[string]interface{}{}
	actionContainer[getJSONExportActionMapName()] = methodMap

	return actionContainer, nil
}

func (mgr *Manager) executeDeleteMethod(q fql.Query) (interface{}, *errors.Error) {
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

func (mgr *Manager) executeExecuteMethod(q fql.Query) (interface{}, *errors.Error) {
	targetObj, ok := q.GetTarget()
	if !ok {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalid)

	}

	methodName := targetObj.String()

	columns, _ := q.GetColumns()
	values, _ := q.GetValues()
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

	resultMap := map[string]interface{}{}
	for name, result := range results {
		resultMap[name] = result.GetValue()
	}

	return resultMap, nil
}
