// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// QueryExecutor represents an interface to execute a query.
type QueryExecutor interface {
	// ExecuteQuery must return the result as a standard array or map.
	ExecuteQuery(Query) (interface{}, error)
}
