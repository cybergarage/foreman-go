// Generated from FQL.g4 by ANTLR 4.7.

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

	// EnterSet is called when entering the Set production.
	EnterSet(c *SetContext)

	// EnterSelect is called when entering the Select production.
	EnterSelect(c *SelectContext)

	// EnterExport is called when entering the Export production.
	EnterExport(c *ExportContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// ExitFql is called when exiting the fql production.
	ExitFql(c *FqlContext)

	// ExitStatement_list is called when exiting the statement_list production.
	ExitStatement_list(c *Statement_listContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSet is called when exiting the Set production.
	ExitSet(c *SetContext)

	// ExitSelect is called when exiting the Select production.
	ExitSelect(c *SelectContext)

	// ExitExport is called when exiting the Export production.
	ExitExport(c *ExportContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)
}
