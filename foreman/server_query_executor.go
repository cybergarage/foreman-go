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

// executeRetransmissionQuery executes the specified query
func (server *Server) executeRetransmissionQuery(node Node, q fql.Query) {
	remoteNode, ok := node.(*RemoteNode)
	if !ok {
		return
	}
	remoteNode.PostRetransmissionQuery(q.String())
}

// ExecuteQuery executes the specified query
func (server *Server) ExecuteQuery(q fql.Query) (interface{}, *errors.Error) {

	// Query type

	queryExecutors := map[fql.QueryType]fql.QueryExecutor{
		fql.QueryTypeExecute: server.actionMgr,
	}

	executor, ok := queryExecutors[q.GetType()]
	if ok {
		return executor.ExecuteQuery(q)
	}

	// Target type

	targetExecutors := map[string]fql.QueryExecutor{
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

	target := strings.ToUpper(targetObj.String())
	executor, ok = targetExecutors[target]
	if !ok {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
	}

	// Execute query to local node

	resObj, err := executor.ExecuteQuery(q)
	if err != nil {
		return nil, err
	}

	// Execute query to remote nodes

	for _, node := range server.GetAllClusterNodes() {
		if NodeEqual(node, server) {
			continue
		}
		go server.executeRetransmissionQuery(node, q)
	}

	return resObj, err
}
