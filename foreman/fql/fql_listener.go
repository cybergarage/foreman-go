// Generated from FQL.g4 by ANTLR 4.6.

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

	// EnterSet_stmt is called when entering the set_stmt production.
	EnterSet_stmt(c *Set_stmtContext)

	// EnterQos_stmt is called when entering the qos_stmt production.
	EnterQos_stmt(c *Qos_stmtContext)

	// ExitFql is called when exiting the fql production.
	ExitFql(c *FqlContext)

	// ExitStatement_list is called when exiting the statement_list production.
	ExitStatement_list(c *Statement_listContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSet_stmt is called when exiting the set_stmt production.
	ExitSet_stmt(c *Set_stmtContext)

	// ExitQos_stmt is called when exiting the qos_stmt production.
	ExitQos_stmt(c *Qos_stmtContext)
}
