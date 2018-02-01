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

// ExitSelect is called when production Select is exited.
func (l *antlrParserListener) ExitSelect(ctx *SelectContext) {
	q := NewSelectQuery()
	p := NewParameterWithObject(parameterTable, ctx.Table().GetText())
	q.AddParameter(p)

	l.Queries = append(l.Queries, q)
}
