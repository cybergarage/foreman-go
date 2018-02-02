// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

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

// ExitSetQuery is called when production SetQuery is exited.
func (l *antlrParserListener) ExitSetQuery(ctx *SetQueryContext) {
	q := NewSetQuery()
	l.Queries = append(l.Queries, q)

	param := NewParameterWithObject(parameterTable, ctx.Target().GetText())
	q.SetParameter(param)

	value := ctx.Value().GetText()
	values := NewValues()
	values = append(values, value)
	param = NewParameterWithObject(parameterValues, values)
	q.SetParameter(param)
}

// ExitSelectQuery is called when production SelectQuery is exited.
func (l *antlrParserListener) ExitSelectQuery(ctx *SelectQueryContext) {
	q := NewSelectQuery()
	l.Queries = append(l.Queries, q)

	param := NewParameterWithObject(parameterTable, ctx.Target().GetText())
	q.SetParameter(param)
}

// ExitExportQuery is called when production ExportQuery is exited.
func (l *antlrParserListener) ExitExportQuery(ctx *ExportQueryContext) {
	q := NewExportQuery()
	l.Queries = append(l.Queries, q)

	param := NewParameterWithObject(parameterTable, ctx.Target().GetText())
	q.SetParameter(param)
}
