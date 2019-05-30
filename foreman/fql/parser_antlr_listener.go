// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"strings"

	"github.com/cybergarage/foreman-go/foreman/util"
)

const (
	parserValueTrimStrings = "\""
)

type antlrParserListener struct {
	*BaseFQLListener
	Queries
	*util.Stack
}

func newANTLRParserListener() *antlrParserListener {
	l := &antlrParserListener{
		BaseFQLListener: &BaseFQLListener{},
		Queries:         NewQueries(),
		Stack:           util.NewStack(),
	}
	return l
}

////////////////////////////////////////
// Insert
////////////////////////////////////////

// EnterInsertQuery is called when production InsertQuery is entered.
func (l *antlrParserListener) EnterInsertQuery(ctx *InsertQueryContext) {
	q := NewInsertQuery()
	l.Push(q)
}

// ExitInsertQuery is called when production InsertQuery is exited.
func (l *antlrParserListener) ExitInsertQuery(ctx *InsertQueryContext) {
	q, ok := l.Pop().(*InsertQuery)
	if !ok {
		return
	}
	l.Queries = append(l.Queries, q)
}

////////////////////////////////////////
// Set
////////////////////////////////////////

// EnterSetQuery is called when production SetQuery is entered.
func (l *antlrParserListener) EnterSetQuery(ctx *SetQueryContext) {
	q := NewSetQuery()
	l.Push(q)
}

// ExitSetQuery is called when production SetQuery is exited.
func (l *antlrParserListener) ExitSetQuery(ctx *SetQueryContext) {
	q, ok := l.Pop().(*SetQuery)
	if !ok {
		return
	}
	l.Queries = append(l.Queries, q)
}

////////////////////////////////////////
// Select
////////////////////////////////////////

// EnterSelectQuery is called when production SelectQuery is entered.
func (l *antlrParserListener) EnterSelectQuery(ctx *SelectQueryContext) {
	q := NewSelectQuery()
	l.Push(q)
}

// ExitSelectQuery is called when production SelectQuery is exited.
func (l *antlrParserListener) ExitSelectQuery(ctx *SelectQueryContext) {
	q, ok := l.Pop().(*SelectQuery)
	if !ok {
		return
	}
	l.Queries = append(l.Queries, q)
}

////////////////////////////////////////
// Export
////////////////////////////////////////

// EnterExportQuery is called when production ExportQuery is entered.
func (l *antlrParserListener) EnterExportQuery(ctx *ExportQueryContext) {
	q := NewExportQuery()
	l.Push(q)
}

// ExitExportQuery is called when production ExportQuery is exited.
func (l *antlrParserListener) ExitExportQuery(ctx *ExportQueryContext) {
	q, ok := l.Pop().(*ExportQuery)
	if !ok {
		return
	}
	l.Queries = append(l.Queries, q)
}

////////////////////////////////////////
// Update
////////////////////////////////////////

// EnterUpdateQuery is called when production UpdateQuery is entered.
func (l *antlrParserListener) EnterUpdateQuery(ctx *UpdateQueryContext) {
	q := NewUpdateQuery()
	l.Push(q)
}

// ExitUpdateQuery is called when production UpdateQuery is exited.
func (l *antlrParserListener) ExitUpdateQuery(ctx *UpdateQueryContext) {
	q, ok := l.Pop().(*UpdateQuery)
	if !ok {
		return
	}
	l.Queries = append(l.Queries, q)
}

////////////////////////////////////////
// Delete
////////////////////////////////////////

// EnterDeleteQuery is called when production DeleteQuery is entered.
func (l *antlrParserListener) EnterDeleteQuery(ctx *DeleteQueryContext) {
	q := NewDeleteQuery()
	l.Push(q)
}

// ExitDeleteQuery is called when production DeleteQuery is exited.
func (l *antlrParserListener) ExitDeleteQuery(ctx *DeleteQueryContext) {
	q, ok := l.Pop().(*DeleteQuery)
	if !ok {
		return
	}
	l.Queries = append(l.Queries, q)
}

////////////////////////////////////////
// Analyze
////////////////////////////////////////

// EnterAnalyzeQuery is called when production AnalyzeQuery is entered.
func (l *antlrParserListener) EnterAnalyzeQuery(ctx *AnalyzeQueryContext) {
	q := NewAnalyzeQuery()
	l.Push(q)
}

// ExitAnalyzeQuery is called when production AnalyzeQuery is exited.
func (l *antlrParserListener) ExitAnalyzeQuery(ctx *AnalyzeQueryContext) {
	q, ok := l.Pop().(*AnalyzeQuery)
	if !ok {
		return
	}
	l.Queries = append(l.Queries, q)
}

////////////////////////////////////////
// Execute
////////////////////////////////////////

// EnterExecuteQuery is called when production ExecuteQuery is entered.
func (l *antlrParserListener) EnterExecuteQuery(ctx *ExecuteQueryContext) {
	q := NewExecuteQuery()
	l.Push(q)
}

// ExitExecuteQuery is called when production ExecuteQuery is exited.
func (l *antlrParserListener) ExitExecuteQuery(ctx *ExecuteQueryContext) {
	q, ok := l.Pop().(*ExecuteQuery)
	if !ok {
		return
	}
	l.Queries = append(l.Queries, q)
}

////////////////////////////////////////
// Target
////////////////////////////////////////

// ExitTarget is called when production target is exited.
func (l *antlrParserListener) ExitTarget(ctx *TargetContext) {
	q, ok := l.Peek().(Query)
	if !ok {
		return
	}
	q.SetTarget(NewTargetWithString(ctx.GetText()))
}

////////////////////////////////////////
// Column (Set/Select)
////////////////////////////////////////

// ExitColumn is called when production column is exited.
func (l *antlrParserListener) ExitColumn(ctx *ColumnContext) {
	q, ok := l.Peek().(Query)
	if !ok {
		return
	}
	q.AddColumn(NewColumnWithString(strings.ToLower(ctx.GetText())))
}

////////////////////////////////////////
// Value (Set/Select)
////////////////////////////////////////

// ExitValue is called when production value is exited.
func (l *antlrParserListener) ExitValue(ctx *ValueContext) {
	q, ok := l.Peek().(Query)
	if !ok {
		return
	}
	value := strings.Trim(ctx.GetText(), parserValueTrimStrings)
	q.AddValue(NewValueWithString(value))
}

////////////////////////////////////////
// Condition (Set/Select)
////////////////////////////////////////

// ExitCondition is called when production condition is exited.
func (l *antlrParserListener) ExitCondition(ctx *ConditionContext) {
	q, ok := l.Peek().(Query)
	if !ok {
		return
	}
	condString := []string{
		ctx.LeftOperand().GetText(),
		ctx.Operator().GetText(),
		strings.Trim(ctx.RightOperand().GetText(), parserValueTrimStrings),
	}
	cond := NewConditionWithObjects(condString)
	q.AddCondition(cond)
}
