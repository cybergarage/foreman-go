// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/discovery"
	"github.com/cybergarage/foreman-go/foreman/metric"
)

const (
	controllerErrorNoFinder = "Finder not found"
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

// HasFinder returns whether the controller has no finder.
func (ctrl *Controller) HasFinder() bool {
	if ctrl.Finder == nil {
		return false
	}
	return true
}

// Start starts the controller.
func (ctrl *Controller) Start() error {
	if ctrl.Finder == nil {
		return fmt.Errorf(controllerErrorNoFinder)
	}

	err := ctrl.Finder.Start()
	if err != nil {
		return err
	}

	return nil
}

// Stop stops the controller.
func (ctrl *Controller) Stop() error {
	if ctrl.Finder == nil {
		return fmt.Errorf(controllerErrorNoFinder)
	}

	err := ctrl.Finder.Stop()
	if err != nil {
		return err
	}

	return nil
}

// Search searches all nodes.
func (ctrl *Controller) Search() error {
	if ctrl.Finder == nil {
		return fmt.Errorf(controllerErrorNoFinder)
	}
	return ctrl.Finder.Search()
}

// SetSearchListener sets a specified listener.
func (ctrl *Controller) SetSearchListener(l discovery.FinderSearchListener) error {
	if ctrl.Finder == nil {
		return fmt.Errorf(controllerErrorNoFinder)
	}
	return ctrl.Finder.SetSearchListener(l)
}

// SetNotifyListener sets a specified listener.
func (ctrl *Controller) SetNotifyListener(l discovery.FinderNotifyListener) error {
	if ctrl.Finder == nil {
		return fmt.Errorf(controllerErrorNoFinder)
	}
	return ctrl.Finder.SetNotifyListener(l)
}

// GetAllNodes returns all found nodes.
func (ctrl *Controller) GetAllNodes() ([]Node, error) {
	if ctrl.Finder == nil {
		return nil, fmt.Errorf(controllerErrorNoFinder)
	}

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
	if ctrl.Finder == nil {
		return nil, fmt.Errorf(controllerErrorNoFinder)
	}

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
	if ctrl.Finder == nil {
		return nil, fmt.Errorf(controllerErrorNoFinder)
	}

	finderNode, err := ctrl.Finder.GetNeighborhoodNode(node)
	if err != nil {
		return nil, err
	}
	return NewRemoteNodeWithNode(finderNode), nil
}
