// Generated from FQL.g4 by ANTLR 4.7.

package fql // FQL

type antlrParserListener struct {
	*BaseFQLListener
	rank int
}

func newANTLRParserListener() *antlrParserListener {
	l := &antlrParserListener{
		BaseFQLListener: &BaseFQLListener{},
		rank:            0,
	}
	return l
}

var _ FQLListener = &BaseFQLListener{}

// ExitSelect is called when production Select is exited.
func (l *antlrParserListener) ExitSelect(ctx *SelectContext) {
	q := NewSelectQuery()
	p := NewParameterWithObject(parameterTable, ctx.Table().GetText())
	q.AddParameter(p)
}
