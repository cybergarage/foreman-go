// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// RuleListener represents a listener interface.
type RuleListener interface {
	RuleSatisfied(Rule)
	RuleUnsatisfied(Rule)
}

// Rule represents a interface.
type Rule interface {
	SetName(string) error
	GetName() string
	AddClause(Clause) error
	GetClauses() []Clause
	GetVariables() []Variable
	IsSatisfied() (bool, error)
	ParseString(Factory, string) error
	String() string
}
