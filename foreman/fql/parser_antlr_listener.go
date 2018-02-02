// Generated from FQL.g4 by ANTLR 4.7.

package fql // FQL

type antlrParserListener struct {
	*BaseFQLListener
	Queries
}

func newANTLRParserListener() *antlrParserListener {
	l := &antlrParserListener{
		BaseFQLListener: &BaseFQLListener{},
		Queries:         NewQueries(),
	}
	return l
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
