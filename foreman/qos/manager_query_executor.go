// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"strings"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

func getJSONExportMapName() string {
	return strings.ToLower(fql.QueryTargetQos)
}

// GetJSONExportPath returns a JSON path for QoS.
func GetJSONExportPath() string {
	return json.PathSep + getJSONExportMapName()
}

func (mgr *Manager) executeInsertQuery(q fql.Query) (interface{}, *errors.Error) {
	values, ok := q.GetValues()
	if !ok || len(values) < 2 {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidValues)
	}

	name := values[0].String()
	formula := values[1].String()

	err := mgr.CreateQoS(name, formula)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return nil, nil
}

func (mgr *Manager) executeSelectQuery(q fql.Query) (interface{}, *errors.Error) {
	ope, name, hasName := q.GetConditionByColumn(fql.QueryColumnName)
	if hasName {
		if ope.GetType() != fql.OperatorTypeEQ {
			hasName = false
		}
	}

	var jsonObj interface{}
	var err error

	if hasName {
		jsonObj, err = mgr.exportQoSJSONObjectWithName(name)
	} else {
		jsonObj, err = mgr.exportQoSJSONObject()
	}

	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	qosContainer := map[string]interface{}{}
	qosContainer[getJSONExportMapName()] = jsonObj

	return qosContainer, nil
}

func (mgr *Manager) executeDeleteQuery(q fql.Query) (interface{}, *errors.Error) {
	ope, name, hasName := q.GetConditionByColumn(fql.QueryColumnName)
	if hasName {
		if ope.GetType() != fql.OperatorTypeEQ {
			hasName = false
		}
	}

	// Delete only a specified rule

	if hasName {
		if !mgr.RemoveRule(name) {
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
