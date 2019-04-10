// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// Factory represents an abstract interface for the parser.
type Factory interface {
	// CreateRule must return a rule instance.
	CreateRule(obj interface{}) (Rule, error)
	// CreateClause must return a clause instance.
	CreateClause(obj interface{}) (Clause, error)
	// CreateFormula must return a formula instance.
	CreateFormula(obj interface{}) (Formula, error)
	// CreateVariable must create a singleton variable, add it into KnowledgeBase.Variables, and return it.
	CreateVariable(obj interface{}) (Variable, error)
	// CreateOperator must return an operater instance.
	CreateOperator(obj interface{}) (Operator, error)
	// CreateObjective must return an objective instance.
	CreateObjective(obj interface{}) (Operand, error)
}
