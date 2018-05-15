// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"strings"
)

// TargetType represents the object type.
type TargetType int

// Target is a destination or source target of FQL.
type Target struct {
	Value string
}

// NewTargetWithString returns a new target with the specified string.
func NewTargetWithString(targetString string) *Target {
	t := &Target{
		Value: targetString,
	}
	return t
}

// NewTarget returns a new target.
func NewTarget() *Target {
	return NewTargetWithString("")
}

// GetTargetType returns the target type.
func (t *Target) GetTargetType() TargetType {
	targetTypes := map[string]TargetType{
		QueryTargetQos:      QueryTargetTypeShared,
		QueryTargetConfig:   QueryTargetTypeStandalone,
		QueryTargetMetrics:  QueryTargetTypeFederated,
		QueryTargetRegister: QueryTargetTypeStandalone,
		QueryTargetRegistry: QueryTargetTypeShared,
		QueryTargetAction:   QueryTargetTypeShared,
		QueryTargetRoute:    QueryTargetTypeShared,
	}

	target := strings.ToUpper(t.Value)
	targetType, ok := targetTypes[target]
	if ok {
		return targetType
	}

	return QueryTargetTypeNone
}

// IsSharedObject returns whether shared object.
func (t *Target) IsSharedObject() bool {
	if t.GetTargetType() == QueryTargetTypeShared {
		return true
	}
	return false
}

// IsFederatedObject returns whether federated object.
func (t *Target) IsFederatedObject() bool {
	if t.GetTargetType() == QueryTargetTypeFederated {
		return true
	}
	return false
}

// String returns the target value.
func (t *Target) String() string {
	return t.Value
}
