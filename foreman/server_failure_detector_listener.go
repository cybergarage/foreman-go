// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/fd"
	"github.com/cybergarage/foreman-go/foreman/fql"
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
func (server *Server) exportAllKnowledgebase(node Node) (fql.Queries, error) {
	queries := fql.NewQueries()

	exportQueries := []string{
		"EXPORT FROM QOS",
		"EXPORT FROM ACTION",
		"EXPORT FROM ROUTE",
	}

	for _, exportQuery := range exportQueries {
		exportRes, err := node.PostQuery(exportQuery)
		if err != nil {
			return nil, err
		}

		exportedQueryStrings, ok := exportRes.([]string)
		if !ok {
			return nil, fmt.Errorf("%s", exportQuery)
		}

		for _, exportedQueryString := range exportedQueryStrings {
			fqlParser := fql.NewParser()
			exportedQueries, err := fqlParser.ParseString(exportedQueryString)
			if err != nil {
				return nil, err
			}
			queries = append(queries, exportedQueries...)
		}
	}

	return queries, nil
}

// FailureDetectorNodeOutOfDate is called when the node is out of date in the cluster.
func (server *Server) FailureDetectorNodeOutOfDate(fd.Node) {
	nodes, err := server.Controller.GetAllNodes()
	if err != nil {
		return
	}

	// FIXME : Choose a latest node

	node := nodes[0]

	latestNode := NewRemoteNodeWithNode(node)
	queries, err := server.exportAllKnowledgebase(latestNode)
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
