// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"strings"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
)

// ExecuteQuery must return the result as a standard array or map.
func (ctrl *Controller) ExecuteQuery(q fql.Query) (interface{}, *errors.Error) {
	if q.GetType() != fql.QueryTypeSelect {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
	}

	finderNodes := []map[string]interface{}{}

	if ctrl.HasFinder() {
		foundNodes, err := ctrl.Finder.GetAllNodes()
		if err != nil {
			return nil, errors.NewErrorWithError(err)
		}

		for _, foundNode := range foundNodes {
			nodeMap := map[string]interface{}{}
			nodeMap[FinderNodeCluster] = foundNode.GetCluster()
			nodeMap[FinderNodeName] = foundNode.GetName()
			nodeMap[FinderNodeAddress] = foundNode.GetAddress()
			nodeMap[FinderNodeRpcPort] = foundNode.GetRPCPort()
			nodeMap[FinderNodeCarbonPort] = foundNode.GetCarbonPort()
			nodeMap[FinderNodeRenderPort] = foundNode.GetRenderPort()
			finderNodes = append(finderNodes, nodeMap)
		}
	}

	finderContainer := map[string]interface{}{}
	finderContainer[strings.ToLower(fql.QueryTargetFinder)] = finderNodes

	return finderContainer, nil
}
