// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// Rule represents a interface.
type Rule interface {
	AddClause(clause Clause) error
	GetClauses() []Clause
	ParseString(factory Factory, ruleString string) error
	String() string
}
