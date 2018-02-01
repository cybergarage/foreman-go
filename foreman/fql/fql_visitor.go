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

	// Visit a parse tree produced by FQLParser#Set.
	VisitSet(ctx *SetContext) interface{}

	// Visit a parse tree produced by FQLParser#Select.
	VisitSelect(ctx *SelectContext) interface{}

	// Visit a parse tree produced by FQLParser#table.
	VisitTable(ctx *TableContext) interface{}
}
