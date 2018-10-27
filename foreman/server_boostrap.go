// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/action"
	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/qos"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

// importBoostrapConfig gets all monitoring configuration.
func (server *Server) importBoostrapConfig(config *boostracConfig) error {
	configMap := config.Objects
	for target, configObj := range configMap {
		var err error
		switch target {
		case fql.QueryTargetQos:
			err = server.qosMgr.ImportQoSJSONObject(configObj)
		case fql.QueryTargetAction:
			err = server.actionMgr.ImportMethodJSONObject(configObj)
		case fql.QueryTargetRoute:
			err = server.actionMgr.ImportRouteJSONObject(configObj)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// exportBoostrapConfig gets all monitoring configuration.
func (node *baseNode) exportBoostrapConfig() (*boostracConfig, error) {
	targets := []string{
		fql.QueryTargetQos,
		fql.QueryTargetAction,
		fql.QueryTargetRoute,
	}

	configMap := map[string]interface{}{}

	for _, target := range targets {
		q := fql.NewSelectAllQuery()
		q.SetTarget(fql.NewTargetWithString(target))
		qStr := q.String()
		jsonObj, err := node.Node.PostQuery(qStr)
		if err != nil {
			return nil, err
		}

		jsonPath := json.NewPathWithObject(jsonObj)

		var configObj interface{}

		switch target {
		case fql.QueryTargetQos:
			configObj, err = jsonPath.GetPathObject(qos.GetJSONExportPath())
		case fql.QueryTargetAction:
			configObj, err = jsonPath.GetPathObject(action.GetJSONExportMethodPath())
		case fql.QueryTargetRoute:
			configObj, err = jsonPath.GetPathObject(action.GetJSONExportRoutePath())
		}

		if err != nil {
			return nil, err
		}

		configMap[target] = configObj
	}

	return NewBoostrapConfigWithObject(configMap), nil
}

// executeBoostrapWithRemoteNode executes boostrap with the specified node.
func (server *Server) executeBoostrapWithRemoteNode(srcNode *RemoteNode) error {
	configObj, err := srcNode.exportBoostrapConfig()
	if err != nil {
		return err
	}

	err = server.importBoostrapConfig(configObj)
	if err != nil {
		return err
	}

	return nil
}

// executeBoostrap executes boostrap.
func (server *Server) executeBoostrap() error {
	srcNode, err := server.Controller.GetNeighborhoodRemoteNode(server)
	if err != nil {
		return err
	}
	return server.executeBoostrapWithRemoteNode(srcNode)
}
