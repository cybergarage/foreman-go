// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
)

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
		case fql.QueryTypeExecute:
			return mgr.executeExecuteAction(q)
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

	return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
}
