// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// Factory represents an abstract interface for the parser.
type Factory interface {
	// CreateRule return a rule instance.
	CreateRule(obj interface{}) (Rule, error)
	// CreateClause return a clause instance.
	CreateClause(obj interface{}) (Clause, error)
	// CreateFormula return a formula instance.
	CreateFormula(obj interface{}) (Formula, error)
	// CreateVariable return a singleton value.
	CreateVariable(obj interface{}) (Variable, error)
	// CreateOperator return an operater instance.
	CreateOperator(obj interface{}) (Operator, error)
	// CreateObjective return an objective instance.
	CreateObjective(obj interface{}) (Objective, error)
}
