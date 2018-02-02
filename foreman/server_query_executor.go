// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"fmt"
	"strings"

	"github.com/cybergarage/foreman-go/foreman/fql"
)

const (
	errorServerQueryTargetNotFound = "Not found the query target (%s) "
)

// ExecuteQuery executes the specified query
func (server *Server) ExecuteQuery(q fql.Query) (interface{}, error) {

	executors := map[string]fql.QueryExecutor{
		"QOS": server.qosMgr,
	}

	target, ok := q.GetTarget()
	if !ok {
		return nil, fmt.Errorf(errorServerQueryTargetNotFound, "")
	}

	target = strings.ToUpper(target)
	executor, ok := executors[target]
	if !ok {
		return nil, fmt.Errorf(errorServerQueryTargetNotFound, target)
	}

	return executor.ExecuteQuery(q)
}
