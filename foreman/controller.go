// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/discovery"
	"github.com/cybergarage/foreman-go/foreman/metric"
)

// Controller represents a controller.
type Controller struct {
	Finder discovery.Finder
}

// NewController returns a new controller.
func NewController() *Controller {
	ctrl := &Controller{
		Finder: nil,
	}
	return ctrl
}

// SetFinder sets a specified finder.
func (ctrl *Controller) SetFinder(finder discovery.Finder) error {
	ctrl.Finder = finder
	return nil
}

// HasFinders returns whether the controller has no finder.
func (ctrl *Controller) HasFinders() bool {
	if ctrl.Finder == nil {
		return false
	}
	return true
}

// Search searches all nodes.
func (ctrl *Controller) Search() error {
	return ctrl.Finder.Search()
}

// SetSearchListener sets a specified listener.
func (ctrl *Controller) SetSearchListener(l discovery.FinderSearchListener) error {
	return ctrl.Finder.SetSearchListener(l)
}

// SetNotifyListener sets a specified listener.
func (ctrl *Controller) SetNotifyListener(l discovery.FinderNotifyListener) error {
	return ctrl.Finder.SetNotifyListener(l)
}

// GetAllNodes returns all found nodes.
func (ctrl *Controller) GetAllNodes() ([]Node, error) {
	finderNodes, err := ctrl.Finder.GetAllNodes()
	if err != nil {
		return nil, err
	}

	allNodes := make([]Node, 0)
	for _, finderNode := range finderNodes {
		allNodes = append(allNodes, NewRemoteNodeWithNode(finderNode))
	}

	return allNodes, nil
}

// GetResponsibleNodesForMetric returns a responsible node for a specified metric.
func (ctrl *Controller) GetResponsibleNodesForMetric(m *metric.Metric) ([]Node, error) {
	name := m.GetName()
	finderNodes, err := ctrl.Finder.GetPrefixNodes(name)
	if err != nil {
		return nil, err
	}

	respNodes := make([]Node, 0)
	for _, finderNode := range finderNodes {
		respNodes = append(respNodes, NewRemoteNodeWithNode(finderNode))
	}

	return respNodes, nil
}

// GetNeighborhoodRemoteNode returns a neighborhood remote node of this node.
func (ctrl *Controller) GetNeighborhoodRemoteNode(node Node) (*RemoteNode, error) {
	finderNode, err := ctrl.Finder.GetNeighborhoodNode(node)
	if err != nil {
		return nil, err
	}
	return NewRemoteNodeWithNode(finderNode), nil
}
