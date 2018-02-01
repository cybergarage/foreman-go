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

func (v *BaseFQLVisitor) VisitSet_stmt(ctx *Set_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFQLVisitor) VisitQos_stmt(ctx *Qos_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}
