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

}

// ExitSelectQuery is called when production SelectQuery is exited.
func (l *antlrParserListener) ExitSelectQuery(ctx *SelectQueryContext) {
	q := NewSelectQuery()
	p := NewParameterWithObject(parameterTable, ctx.Target().GetText())
	q.AddParameter(p)

	l.Queries = append(l.Queries, q)
}

// ExitExportQuery is called when production ExportQuery is exited.
func (l *antlrParserListener) ExitExportQuery(ctx *ExportQueryContext) {
	q := NewExportQuery()
	p := NewParameterWithObject(parameterTable, ctx.Target().GetText())
	q.AddParameter(p)

	l.Queries = append(l.Queries, q)
}
