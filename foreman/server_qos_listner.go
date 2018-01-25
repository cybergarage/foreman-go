// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import "github.com/cybergarage/foreman-go/foreman/kb"

// RuleSatisfied is a listener for kb.Rule
func (server *Server) RuleSatisfied(kb.Rule) {
	// TODO : Post the event to the action manager
}

// RuleUnsatisfied is a listener for kb.Rule
func (server *Server) RuleUnsatisfied(kb.Rule) {
}
