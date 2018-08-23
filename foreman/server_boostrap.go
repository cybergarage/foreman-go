// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/fql"
)

// impportMonitoringConfigurations gets all monitoring configuration.
func (server *Server) impportMonitoringConfigurations(configMap map[string]interface{}) error {
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

// updateMonitoringConfigurationWithRemoteNode updates the configuration from the specified node
func (server *Server) updateMonitoringConfigurationWithRemoteNode(node *RemoteNode) error {
	configObj, err := node.exportMonitoringConfigurations()
	if err != nil {
		return err
	}

	err = server.impportMonitoringConfigurations(configObj)
	if err != nil {
		return err
	}

	return nil
}
