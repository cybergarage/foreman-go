// Generated from FQL.g4 by ANTLR 4.7.

package fql // FQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFQLListener is a complete listener for a parse tree produced by FQLParser.
type BaseFQLListener struct{}

var _ FQLListener = &BaseFQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatement_list is called when production statement_list is entered.
func (s *BaseFQLListener) EnterStatement_list(ctx *Statement_listContext) {}

// ExitStatement_list is called when production statement_list is exited.
func (s *BaseFQLListener) ExitStatement_list(ctx *Statement_listContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseFQLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseFQLListener) ExitStatement(ctx *StatementContext) {}

// EnterSet_stmt is called when production set_stmt is entered.
func (s *BaseFQLListener) EnterSet_stmt(ctx *Set_stmtContext) {}

// ExitSet_stmt is called when production set_stmt is exited.
func (s *BaseFQLListener) ExitSet_stmt(ctx *Set_stmtContext) {}
