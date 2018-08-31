// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package node

// Condition represents node condition types
type Condition uint

// Clock represents a node clock type
type Clock uint

// Version represents a node version type
type Version uint

const (
	ConditionUnknown   = 0x00
	ConditionInitial   = 0x10
	ConditionBoostrap  = 0x20
	ConditionReady     = 0x30
	ConditionStop      = 0x31
	ConditionOutOfDate = 0x32
)

// Status represents an abstract node interface for the status
type Status interface {
	// GetCondition returns the current status
	GetCondition() Condition
	// GetClock returns the current logical clock
	GetClock() Clock
	// GetVersion returns the current repository version
	GetVersion() Version
}
