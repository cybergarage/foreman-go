// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

// executeBootstrapConfig executes to get configuration from other node in the cluster.
func (server *Server) executeBootstrap() error {
	var lastErr error

	if server.HasBootstrapQuery() {
		err := server.executeBootstrapQuery()
		if err != nil {
			lastErr = err
		}
	}

	if server.IsBootstrapEnabled() {
		err := server.executeBootstrapConfig()
		if err != nil {
			lastErr = err
		}
	}

	return lastErr
}
