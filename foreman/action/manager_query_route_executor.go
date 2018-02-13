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
	source := values[1].String()
	dest := values[2].String()

	err := mgr.CreateRoute(name, source, dest)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return nil, nil
}

func (mgr *Manager) executeSelectRoute(q fql.Query) (interface{}, *errors.Error) {
	ope, whereName, hasWhere := q.GetConditionByColumn(fql.QueryColumnName)
	if hasWhere {
		if ope.GetType() != fql.OperatorTypeEQ {
			return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
		}
	}

	var routes []interface{}
	for _, route := range mgr.GetAllRoutes() {
		if hasWhere {
			if route.Name != whereName {
				continue
			}
		}
		var routeMap map[string]string
		routeMap[RouteColumnName] = route.Name
		routeMap[RouteColumnSource] = route.Source.GetName()
		routeMap[RouteColumnDestination] = route.Destination.GetName()
		routes = append(routes, routeMap)
	}

	var routeMap map[string]interface{}
	routeMap[RouteColumnRoutes] = routes
	actionContainer := map[string]interface{}{}
	actionContainer[strings.ToLower(fql.QueryTargetAction)] = routeMap

	return actionContainer, nil
}

func (mgr *Manager) executeDeleteRoute(q fql.Query) (interface{}, *errors.Error) {
	ope, whereName, hasWhere := q.GetConditionByColumn(fql.QueryColumnName)
	if hasWhere {
		if ope.GetType() != fql.OperatorTypeEQ {
			return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryInvalidConditions)
		}
	}

	// Remove only a specified route

	if hasWhere {
		ok := mgr.RemoveRoute(whereName)
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
