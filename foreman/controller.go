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
	Finders []discovery.Finder
}

// NewController returns a new controller.
func NewController() *Controller {
	ctrl := &Controller{
		Finders: make([]discovery.Finder, 0),
	}
	return ctrl
}

// AddFinder adds a specified finder.
func (ctrl *Controller) AddFinder(finder discovery.Finder) error {
	ctrl.Finders = append(ctrl.Finders, finder)
	return nil
}

// SearchAll searches all nodes.
func (ctrl *Controller) SearchAll() error {
	for _, finder := range ctrl.Finders {
		err := finder.SearchAll()
		if err != nil {
			return err
		}
	}
	return nil
}

// SetSearchListener sets a specified listener.
func (ctrl *Controller) SetSearchListener(l discovery.FinderSearchListener) error {
	for _, finder := range ctrl.Finders {
		err := finder.SetSearchListener(l)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetNotifyListener sets a specified listener.
func (ctrl *Controller) SetNotifyListener(l discovery.FinderNotifyListener) error {
	for _, finder := range ctrl.Finders {
		err := finder.SetNotifyListener(l)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAllNodes returns all found nodes.
func (ctrl *Controller) GetAllNodes() ([]Node, error) {
	allNodes := make([]Node, 0)
	for _, finder := range ctrl.Finders {
		nodes, err := finder.GetAllNodes()
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			allNodes = append(allNodes, NewRemoteNodeWithDiscoveryNode(node))
		}
	}
	return allNodes, nil
}

// GetResponsibleNodesForMetric returns a responsible node for a specified metric.
func (ctrl *Controller) GetResponsibleNodesForMetric(m *metric.Metric) ([]Node, error) {
	name := m.GetName()
	respNodes := make([]Node, 0)
	for _, finder := range ctrl.Finders {
		nodes, err := finder.GetPrefixNodes(name)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			respNodes = append(respNodes, NewRemoteNodeWithDiscoveryNode(node))
		}
	}
	return respNodes, nil
}
