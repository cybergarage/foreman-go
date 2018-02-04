// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"strings"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
)

const (
	errorServerQueryTargetNotFound = "Not found the query target (%s) "
)

// ExecuteQuery executes the specified query
func (server *Server) ExecuteQuery(q fql.Query) (interface{}, *errors.Error) {

	executors := map[string]fql.QueryExecutor{
		QueryTargetQos:      server.qosMgr,
		QueryTargetConfig:   server.config,
		QueryTargetMetric:   server.metricMgr,
		QueryTargetRegister: server.registerMgr,
		QueryTargetRegistry: server.registryMgr,
	}

	target, ok := q.GetTarget()
	if !ok {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryTargetNotFound)
	}

	target = strings.ToUpper(target)
	executor, ok := executors[target]
	if !ok {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryTargetNotFound)
	}

	return executor.ExecuteQuery(q)
}
