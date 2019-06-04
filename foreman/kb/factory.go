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
	// CreateLiteralOperand must return a literal operand of the specified object.
	CreateLiteralOperand(obj interface{}) (Literal, error)
	// CreateVariableOperand must return a variable operand of the specified object.
	CreateVariableOperand(name string) (Variable, error)
	// CreateFunctionOperand must return a function operand of the specified objects.
	CreateFunctionOperand(name string) (Function, error)
}
