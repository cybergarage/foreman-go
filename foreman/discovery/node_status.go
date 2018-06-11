// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

// NodeCondition represents node condition types
type NodeCondition uint

// NodeClock represents a node clock type
type NodeClock uint

// NodeClock represents a node clock type

const (
	NodeConditionUnknown = iota
	NodeConditionInitial
	NodeConditionReady
	NodeConditionStop
)

// NodeStatus represents an abstract node interface for the status
type NodeStatus interface {
	// GetCondition returns the current status
	GetCondition() NodeCondition
	// GetClock returns the current logical clock
	GetClock() NodeClock
}
