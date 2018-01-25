// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"github.com/cybergarage/foreman-go/foreman/kb"
	"github.com/cybergarage/foreman-go/foreman/metric"
)

// Rule represents a formula.
type Rule struct {
	metric.RegisterListener
	*kb.BaseRule
}

// NewRule returns a new rule.
func NewRule() *Rule {
	rule := &Rule{
		BaseRule: kb.NewRule(),
	}
	return rule
}

// ObjectUpdated is a listener for the metric register store.
func (rule *Rule) ObjectUpdated() {
}
