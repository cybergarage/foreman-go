// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import "strings"

type antlrParserListener struct {
	*BaseFQLListener
	Queries
	*ParserObjectStack
}

func newANTLRParserListener() *antlrParserListener {
	l := &antlrParserListener{
		BaseFQLListener:   &BaseFQLListener{},
		Queries:           NewQueries(),
		ParserObjectStack: NewParserObjectStack(),
	}
	return l
}

////////////////////////////////////////
// Insert
////////////////////////////////////////

// EnterInsertQuery is called when production InsertQuery is entered.
func (l *antlrParserListener) EnterInsertQuery(ctx *InsertQueryContext) {
	q := NewInsertQuery()
	l.PushObject(q)
}

// ExitInsertQuery is called when production InsertQuery is exited.
func (l *antlrParserListener) ExitInsertQuery(ctx *InsertQueryContext) {
	q, ok := l.PopObject().(*InsertQuery)
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
	l.PushObject(q)
}

// ExitSetQuery is called when production SetQuery is exited.
func (l *antlrParserListener) ExitSetQuery(ctx *SetQueryContext) {
	q, ok := l.PopObject().(*SetQuery)
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
	l.PushObject(q)
}

// ExitSelectQuery is called when production SelectQuery is exited.
func (l *antlrParserListener) ExitSelectQuery(ctx *SelectQueryContext) {
	q, ok := l.PopObject().(*SelectQuery)
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
	l.PushObject(q)
}

// ExitExportQuery is called when production ExportQuery is exited.
func (l *antlrParserListener) ExitExportQuery(ctx *ExportQueryContext) {
	q, ok := l.PopObject().(*ExportQuery)
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
	l.PushObject(q)
}

// ExitDeleteQuery is called when production DeleteQuery is exited.
func (l *antlrParserListener) ExitDeleteQuery(ctx *DeleteQueryContext) {
	q, ok := l.PopObject().(*DeleteQuery)
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
	q, ok := l.PeekObject().(Query)
	if !ok {
		return
	}
	q.SetTarget(NewTargetWithString(ctx.GetText()))
}

////////////////////////////////////////
// Value (Set/Select)
////////////////////////////////////////

// ExitValue is called when production value is exited.
func (l *antlrParserListener) ExitValue(ctx *ValueContext) {
	q, ok := l.PeekObject().(Query)
	if !ok {
		return
	}
	value := strings.Trim(ctx.GetText(), "\"")
	q.AddValue(NewValueWithString(value))
}

////////////////////////////////////////
// Condition (Set/Select)
////////////////////////////////////////

// ExitCondition is called when production condition is exited.
func (l *antlrParserListener) ExitCondition(ctx *ConditionContext) {
	q, ok := l.PeekObject().(Query)
	if !ok {
		return
	}
	condString := []string{
		ctx.LeftOperand().GetText(),
		ctx.Operator().GetText(),
		ctx.RightOperand().GetText(),
	}
	cond := NewConditionWithObjects(condString)
	q.AddCondition(cond)
}
