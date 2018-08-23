// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"strings"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
)

func (mgr *Manager) executeInsertRoute(q fql.Query) (interface{}, *errors.Error) {
	values, ok := q.GetValues()
	if !ok || len(values) < 3 {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidValues)
	}

	name := values[0].String()
	src := values[1].String()
	dst := values[2].String()

	err := mgr.CreateRoute(name, src, dst)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return nil, nil
}

func (mgr *Manager) executeSelectRoute(q fql.Query) (interface{}, *errors.Error) {
	ope, name, hasName := q.GetConditionByColumn(fql.QueryColumnName)
	if hasName {
		if ope.GetType() != fql.OperatorTypeEQ {
			return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
		}
	}

	var routes interface{}
	var err error

	if hasName {
		routes, err = mgr.exportRouteJSONObjectWithName(name)
	} else {
		routes, err = mgr.exportRouteJSONObject()
	}

	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	routeMap := map[string]interface{}{}
	routeMap[RouteColumnRoutes] = routes

	actionContainer := map[string]interface{}{}
	actionContainer[strings.ToLower(fql.QueryTargetAction)] = routeMap

	return actionContainer, nil
}

func (mgr *Manager) executeDeleteRoute(q fql.Query) (interface{}, *errors.Error) {
	ope, name, hasName := q.GetConditionByColumn(fql.QueryColumnName)
	if hasName {
		if ope.GetType() != fql.OperatorTypeEQ {
			return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
		}
	}

	// Remove only a specified route

	if hasName {
		ok := mgr.RemoveRoute(name)
		if !ok {
			return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
		}
		return nil, nil
	}

	// Remove all routes

	ok := mgr.RemoveAllRoutes()
	if !ok {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
	}

	return nil, nil
}
