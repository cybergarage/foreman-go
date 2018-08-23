// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/action"
	"github.com/cybergarage/foreman-go/foreman/qos"

	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

const (
	errorNodeExportObjectNotFound = "Export Object Not Found (%s) : %v"
)

// baseNode is a super class to define methods using Node interface.
type baseNode struct {
	Node Node
}

// newBaseNodeWithNode returns a remote new node.
func newBaseNodeWithNode(node Node) *baseNode {
	baseNode := &baseNode{
		Node: node,
	}
	return baseNode
}

// exportMonitoringConfigurations gets all monitoring configuration.
func (node *baseNode) exportMonitoringConfigurations() (map[string]interface{}, error) {
	targets := []string{
		fql.QueryTargetQos,
		fql.QueryTargetAction,
		fql.QueryTargetRoute,
	}

	configMap := map[string]interface{}{}

	for _, target := range targets {
		q := fql.NewSelectAllQuery()
		q.SetTarget(fql.NewTargetWithString(target))
		qStr := q.String()
		jsonObj, err := node.Node.PostQuery(qStr)
		if err != nil {
			return nil, err
		}

		jsonPath := json.NewPathWithObject(jsonObj)

		var configObj interface{}

		switch target {
		case fql.QueryTargetQos:
			configObj, err = jsonPath.GetPathObject(qos.GetJSONExportPath())
		case fql.QueryTargetAction:
			configObj, err = jsonPath.GetPathObject(action.GetJSONExportMethodPath())
		case fql.QueryTargetRoute:
			configObj, err = jsonPath.GetPathObject(action.GetJSONExportRoutePath())
		}

		if err != nil {
			return nil, err
		}

		configMap[target] = configObj
	}

	return configMap, nil
}
