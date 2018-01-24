// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"github.com/cybergarage/foreman-go/foreman/kb"
	"github.com/cybergarage/foreman-go/foreman/metric"
)

// Formula represents a formula.
type Formula struct {
	metric.RegisterListener
	*kb.BaseFormula
}

// NewFormula returns a new base formula.
func NewFormula() *Formula {
	formula := &Formula{
		BaseFormula: kb.NewFormula(),
	}
	return formula
}
