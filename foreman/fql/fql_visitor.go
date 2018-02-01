// Generated from FQL.g4 by ANTLR 4.7.

package fql // FQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by FQLParser.
type FQLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FQLParser#fql.
	VisitFql(ctx *FqlContext) interface{}

	// Visit a parse tree produced by FQLParser#statement_list.
	VisitStatement_list(ctx *Statement_listContext) interface{}

	// Visit a parse tree produced by FQLParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by FQLParser#set_stmt.
	VisitSet_stmt(ctx *Set_stmtContext) interface{}

	// Visit a parse tree produced by FQLParser#qos_stmt.
	VisitQos_stmt(ctx *Qos_stmtContext) interface{}
}
