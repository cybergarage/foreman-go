// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

import (
	"fmt"
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
	value := values[1].String()

	obj := NewObject()
	obj.SetName(name)
	obj.SetData(value)

	err := mgr.SetObject(obj)
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

	if !hasName {
		return nil, errors.NewErrorWithError(fmt.Errorf(errorQueryEmptyID))
	}

	obj, ok := mgr.GetObject(name)
	if !ok {
		return nil, errors.NewErrorWithError(fmt.Errorf(errorQueryNotFoundID, name))
	}

	data, err := obj.GetData()
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	ver, err := obj.GetVersion()
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	ts, err := obj.GetTimestamp()
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	regData := map[string]interface{}{}
	regData[fql.QueryColumnValue] = data
	regData[fql.QueryColumnVersion] = ver
	regData[fql.QueryColumnTimestamp] = ts.Unix()

	regMap := map[string]interface{}{}
	regMap[name] = regData

	regContainer := map[string]interface{}{}
	regContainer[strings.ToLower(fql.QueryTargetRegister)] = regMap

	return regContainer, nil
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
		if !mgr.RemoveObject(name) {
			return nil, errors.NewErrorWithError(fmt.Errorf(errorQueryNotFoundID, name))
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
