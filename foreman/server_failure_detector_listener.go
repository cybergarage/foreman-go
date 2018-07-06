// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/fd"
	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/node"
)

const (
	exportAllQoS     = "EXPORT FROM QOS"
	exportAllActions = "EXPORT FROM ACTION"
	exportAllRoutes  = "EXPORT FROM ROUTE"
)

// FailureDetectorNodeAdded is called when a new node added in the cluster.
func (server *Server) FailureDetectorNodeAdded(fd.Node) {

}

// FailureDetectorNodeRemoved is called when the node is removed in the cluster.
func (server *Server) FailureDetectorNodeRemoved(fd.Node) {

}

// FailureDetectorNodeStatusChanged is called when a status of the node is changed  in the cluster.
func (server *Server) FailureDetectorNodeStatusChanged(fd.Node) {

}

// exportAllKnowledgebase export all knowledgebase
func (server *Server) exportAllKnowledgebase(node node.Node) (fql.Queries, error) {
	return nil, nil
}

// FailureDetectorNodeOutOfDate is called when the node is out of date in the cluster.
func (server *Server) FailureDetectorNodeOutOfDate(fd.Node) {
	finder := server.finder
	nodes, err := finder.GetAllNodes()
	if err != nil {
		return
	}

	// FIXME : Choose a latest node

	node := nodes[0]

	queries, err := server.exportAllKnowledgebase(node)
	if err != nil {
		return
	}

	for _, query := range queries {
		_, err := server.ExecuteQuery(query)
		if err != nil {
			// TODO : Recovery errors
		}
	}

	// TODO : Update version
}
