// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// Factory represents an abstract interface for the parser.
type Factory interface {
	// CreateRule must return a rule instance of the specified object.
	CreateRule() Rule
	// CreateClause must return a clause instance of the specified object.
	CreateClause() Clause
	// CreateFormula must return a formula instance  of the specified object.
	CreateFormula(lope Operand, op Operator, rope Operand) Formula
	// CreateOperator must return an operater of the specified object.
	CreateOperator(obj interface{}) (Operator, error)
	// CreateOperand must return an operand of the specified object.
	CreateOperand(obj interface{}) (Operand, error)
}
