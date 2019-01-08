// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"strings"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/logging"
)

const (
	errorServerQueryTargetNotFound = "Not found the query target (%s) "
)

// executeRetransmissionQuery executes the specified query
func (server *Server) executeRetransmissionQuery(node Node, query fql.Query) error {
	remoteNode := node.(*RemoteNode)
	queryString := query.String()
	_, err := remoteNode.PostRetransmissionQuery(queryString)
	if err != nil {
		logging.Error("%s", err)
	}
	return err
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

	// Target type (Config)

	targetObj, ok := q.GetTarget()
	if !ok {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryEmptyTarget)
	}

	target := strings.ToUpper(targetObj.String())

	// Target type (Config)

	switch target {
	case fql.QueryTargetConfig:
		return server.ExecuteConfigQuery(q)
	}

	// Target type

	targetExecutors := map[string]fql.QueryExecutor{
		fql.QueryTargetQos:      server.qosMgr,
		fql.QueryTargetMetrics:  server.metricMgr,
		fql.QueryTargetRegister: server.registerMgr,
		fql.QueryTargetRegistry: server.registryMgr,
		fql.QueryTargetAction:   server.actionMgr,
		fql.QueryTargetRoute:    server.actionMgr,
		fql.QueryTargetFinder:   server.Controller,
	}

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

	if q.IsStateChangeQuery() && !q.IsRetransmissionQuery() {
		for _, node := range server.GetAllClusterNodes() {
			if NodeEqual(node, server) {
				continue
			}
			go server.executeRetransmissionQuery(node, q)
		}
	}

	return resObj, err
}
