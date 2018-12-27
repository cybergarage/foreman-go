// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/logging"
)

// Restart restats the server.
func (server *Server) Restart() error {
	logging.Info("Stopping server for restart...\n")
	err := server.Stop()
	if err != nil {
		logging.Error("Could not stop server: %s\n", err)
		return err
	}

	if 0 < len(server.configFile) {
		logging.Info("Reloading config...")
		err := server.LoadConfig(server.configFile)
		if err != nil {
			logging.Error(err.Error())
			return err
		}
	}

	if 0 < len(server.queryFile) {
		logging.Info("Reloading query...")
		err := server.clearBootstrapConfig()
		if err != nil {
			logging.Error(err.Error())
			return err
		}

		err = server.LoadQuery(server.queryFile)
		if err != nil {
			logging.Error(err.Error())
			return err
		}
	}

	logging.Info("Restarting server...")
	err = server.Start()
	if err != nil {
		logging.Error("Could not start server: %s\n", err)
		return err
	}
	logging.Info("Restarted.")

	return nil
}
