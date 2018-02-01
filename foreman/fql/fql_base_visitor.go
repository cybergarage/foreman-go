// Generated from FQL.g4 by ANTLR 4.7.

package fql // FQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseFQLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFQLVisitor) VisitFql(ctx *FqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFQLVisitor) VisitStatement_list(ctx *Statement_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFQLVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFQLVisitor) VisitSet(ctx *SetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFQLVisitor) VisitSelect(ctx *SelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFQLVisitor) VisitTable(ctx *TableContext) interface{} {
	return v.VisitChildren(ctx)
}
