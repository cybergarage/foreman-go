// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

// executeBootstrapQuery executes the specified boostrap queries.
func (server *Server) executeBootstrapQuery() error {
	queryFiles, err := server.GetBootstrapQueryFiles()
	if err != nil {
		return err
	}

	var lastError error

	for _, queryFile := range queryFiles {
		err := server.LoadQuery(queryFile.GetPath())
		if err != nil {
			lastError = err
		}
	}

	return lastError
}
