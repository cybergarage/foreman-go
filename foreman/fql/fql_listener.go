// Code generated from FQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package fql // FQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// FQLListener is a complete listener for a parse tree produced by FQLParser.
type FQLListener interface {
	antlr.ParseTreeListener

	// EnterFql is called when entering the fql production.
	EnterFql(c *FqlContext)

	// EnterStatement_list is called when entering the statement_list production.
	EnterStatement_list(c *Statement_listContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterInsertQuery is called when entering the InsertQuery production.
	EnterInsertQuery(c *InsertQueryContext)

	// EnterSetQuery is called when entering the SetQuery production.
	EnterSetQuery(c *SetQueryContext)

	// EnterSelectQuery is called when entering the SelectQuery production.
	EnterSelectQuery(c *SelectQueryContext)

	// EnterExportQuery is called when entering the ExportQuery production.
	EnterExportQuery(c *ExportQueryContext)

	// EnterDeleteQuery is called when entering the DeleteQuery production.
	EnterDeleteQuery(c *DeleteQueryContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterValues is called when entering the values production.
	EnterValues(c *ValuesContext)

	// ExitFql is called when exiting the fql production.
	ExitFql(c *FqlContext)

	// ExitStatement_list is called when exiting the statement_list production.
	ExitStatement_list(c *Statement_listContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitInsertQuery is called when exiting the InsertQuery production.
	ExitInsertQuery(c *InsertQueryContext)

	// ExitSetQuery is called when exiting the SetQuery production.
	ExitSetQuery(c *SetQueryContext)

	// ExitSelectQuery is called when exiting the SelectQuery production.
	ExitSelectQuery(c *SelectQueryContext)

	// ExitExportQuery is called when exiting the ExportQuery production.
	ExitExportQuery(c *ExportQueryContext)

	// ExitDeleteQuery is called when exiting the DeleteQuery production.
	ExitDeleteQuery(c *DeleteQueryContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitValues is called when exiting the values production.
	ExitValues(c *ValuesContext)
}
