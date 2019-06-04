// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import (
	"github.com/cybergarage/foreman-go/foreman/util"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type antlrParserListener struct {
	*BaseKnowledgeListener
	*util.Stack
	Factory
	Rule
}

func newANTLRParserListener(rule Rule, factory Factory) *antlrParserListener {
	l := &antlrParserListener{
		BaseKnowledgeListener: &BaseKnowledgeListener{},
		Stack:                 util.NewStack(),
		Factory:               factory,
		Rule:                  rule,
	}
	return l
}

func (l *antlrParserListener) SetInternalError(parser antlr.Parser, err error) {
	el, ok := (parser.GetErrorListenerDispatch()).(*antlrParserErrorListener)
	if ok {
		el.SetInternalError(err)
	}
}

// EnterClause is called when production clause is entered.
func (l *antlrParserListener) EnterClause(ctx *ClauseContext) {
	clause := l.Factory.CreateClause()
	l.Push(clause)
	l.Rule.AddClause(clause)
}

// ExitClause is called when production clause is exited.
func (l *antlrParserListener) ExitClause(ctx *ClauseContext) {
	l.Pop()
}

// ExitFormula is called when production formula is exited.
func (l *antlrParserListener) ExitFormula(ctx *FormulaContext) {
	rop, ok := l.Pop().(Operand)
	if !ok {
		return
	}
	op, ok := l.Pop().(Operator)
	if !ok {
		return
	}
	lop, ok := l.Pop().(Operand)
	if !ok {
		return
	}

	clause, ok := l.Peek().(Clause)
	if !ok {
		return
	}

	clause.AddFormula(l.Factory.CreateFormula(lop, op, rop))
}

// EnterOperator is called when production operator is entered.
func (l *antlrParserListener) EnterOperator(ctx *OperatorContext) {
	op, err := l.Factory.CreateOperator(ctx.GetText())
	if err != nil {
		l.SetInternalError(ctx.GetParser(), err)
		return
	}
	l.Push(op)
}

// EnterLiteralOperand is called when production literalOperand is entered.
func (l *antlrParserListener) EnterLiteralOperand(ctx *LiteralOperandContext) {
	op, err := l.Factory.CreateLiteralOperand(ctx.GetText())
	if err != nil {
		l.SetInternalError(ctx.GetParser(), err)
		return
	}
	l.Push(op)
}

// EnterVariableOperand is called when production variableOperand is entered.
func (l *antlrParserListener) EnterVariableOperand(ctx *VariableOperandContext) {
	op, err := l.Factory.CreateVariableOperand(ctx.GetText())
	if err != nil {
		l.SetInternalError(ctx.GetParser(), err)
		return
	}
	l.Push(op)
}