// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
)

const (
	errorServerQueryTargetNotFound = "Not found the query target (%s) "
)

// ExecuteQuery executes the specified query
func (server *Server) ExecuteQuery(q fql.Query) (interface{}, *errors.Error) {

	executors := map[string]fql.QueryExecutor{
		fql.QueryTargetQos:      server.qosMgr,
		fql.QueryTargetConfig:   server.config,
		fql.QueryTargetMetrics:  server.metricMgr,
		fql.QueryTargetRegister: server.registerMgr,
		fql.QueryTargetRegistry: server.registryMgr,
		fql.QueryTargetAction:   server.actionMgr,
		fql.QueryTargetRoute:    server.actionMgr,
	}

	targetObj, ok := q.GetTarget()
	if !ok {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryEmptyTarget)
	}

	target := targetObj.String()
	executor, ok := executors[target]
	if !ok {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
	}

	return executor.ExecuteQuery(q)
}
