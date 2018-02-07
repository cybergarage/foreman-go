// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package registry

import (
	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
)

// collectChildRegistryObjects adds the child objects into the parent map.
func (mgr *Manager) collectChildRegistryObjects(parentMap map[string]interface{}, parentObjectID string) error {
	childObjs, err := mgr.GetChildObjects(parentObjectID)
	if err != nil {
		return err
	}

	for _, childObj := range childObjs {
		grandChildObjs, err := mgr.GetChildObjects(childObj.ID)
		if err != nil {
			return nil
		}

		if len(grandChildObjs) <= 0 {
			parentMap[childObj.Name] = childObj.Data
			continue
		}

		childObjMap := map[string]interface{}{}
		parentMap[childObj.Name] = childObjMap
		err = mgr.collectChildRegistryObjects(childObjMap, childObj.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// ExecuteQuery must return the result as a standard array or map.
func (mgr *Manager) ExecuteQuery(q fql.Query) (interface{}, *errors.Error) {
	if q.GetType() != fql.QueryTypeSelect {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
	}

	regMap := map[string]interface{}{}

	err := mgr.collectChildRegistryObjects(regMap, RootObjectID)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return regMap, nil
}
