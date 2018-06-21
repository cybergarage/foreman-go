// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/node"
)

// NodeCondition represents node condition types
type NodeCondition = node.Condition

// NodeClock represents a node clock type
type NodeClock = node.Clock

// NodeVersion represents a node version
type NodeVersion = node.Version

const (
	NodeConditionUnknown = node.ConditionUnknown
	NodeConditionInitial = node.ConditionInitial
	NodeConditionReady   = node.ConditionReady
	NodeConditionStop    = node.ConditionStop
)

// Status represents a node status
type Status struct {
	Condition NodeCondition
	Clock     NodeClock
	Version   NodeVersion
}

// NewStatus returns a new status instance.
func NewStatus() *Status {
	status := &Status{
		Condition: NodeConditionInitial,
		Clock:     0,
		Version:   0,
	}
	return status
}

// SetCondition sets a new condition status
func (status *Status) SetCondition(cond NodeCondition) {
	status.Condition = cond
}

// GetCondition returns the current status
func (status *Status) GetCondition() NodeCondition {
	return status.Condition
}

// SetClock sets a new clock
func (status *Status) SetClock(clock NodeClock) {
	status.Clock = clock
}

// GetClock returns the current logical clock
func (status *Status) GetClock() NodeClock {
	return status.Clock
}

// SetVersion sets a new version
func (status *Status) SetVersion(clock NodeVersion) {
	status.Version = clock
}

// GetVersion returns the current repository version
func (status *Status) GetVersion() NodeVersion {
	return status.Version
}
