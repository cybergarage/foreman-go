// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

// retrieveMonitoringConfigurationWithRemoteNode gets all monitoring configuration from the specified remote node.
func (server *Server) retrieveMonitoringConfigurationWithRemoteNode(node *RemoteNode) error {
	err := node.exportMonitoringConfigurations()
	if err != nil {
		return err
	}
	return nil
}

// updateMonitoringConfigurationWithRemoteNode updates the configuration from the specified node
func (server *Server) updateMonitoringConfigurationWithRemoteNode(node *RemoteNode) error {
	err := server.retrieveMonitoringConfigurationWithRemoteNode(node)
	if err != nil {
		return err
	}
	return nil
}
