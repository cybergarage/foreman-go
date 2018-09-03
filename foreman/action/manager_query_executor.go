// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"strings"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
)

func getJSONExportActionMapName() string {
	return strings.ToLower(fql.QueryTargetAction)
}

// ExecuteQuery must return the result as a standard array or map.
func (mgr *Manager) ExecuteQuery(q fql.Query) (interface{}, *errors.Error) {
	// Target type

	targetObj, ok := q.GetTarget()
	if !ok {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryEmptyTarget)
	}

	switch strings.ToUpper(targetObj.String()) {
	case fql.QueryTargetAction:
		switch q.GetType() {
		case fql.QueryTypeInsert:
			return mgr.executeInsertMethod(q)
		case fql.QueryTypeSelect:
			return mgr.executeSelectMethod(q)
		case fql.QueryTypeDelete:
			return mgr.executeDeleteMethod(q)
		case fql.QueryTypeExecute:
			return mgr.executeExecuteMethod(q)
		}
	case fql.QueryTargetRoute:
		switch q.GetType() {
		case fql.QueryTypeInsert:
			return mgr.executeInsertRoute(q)
		case fql.QueryTypeSelect:
			return mgr.executeSelectRoute(q)
		case fql.QueryTypeDelete:
			return mgr.executeDeleteRoute(q)
		}
	}

	// Query type

	switch q.GetType() {
	case fql.QueryTypeExecute:
		return mgr.executeExecuteMethod(q)
	}

	return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
}
