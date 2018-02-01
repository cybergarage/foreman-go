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

// EnterFql is called when production fql is entered.
func (s *BaseFQLListener) EnterFql(ctx *FqlContext) {}

// ExitFql is called when production fql is exited.
func (s *BaseFQLListener) ExitFql(ctx *FqlContext) {}

// EnterStatement_list is called when production statement_list is entered.
func (s *BaseFQLListener) EnterStatement_list(ctx *Statement_listContext) {}

// ExitStatement_list is called when production statement_list is exited.
func (s *BaseFQLListener) ExitStatement_list(ctx *Statement_listContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseFQLListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseFQLListener) ExitStatement(ctx *StatementContext) {}

// EnterSet is called when production Set is entered.
func (s *BaseFQLListener) EnterSet(ctx *SetContext) {}

// ExitSet is called when production Set is exited.
func (s *BaseFQLListener) ExitSet(ctx *SetContext) {}

// EnterSelect is called when production Select is entered.
func (s *BaseFQLListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production Select is exited.
func (s *BaseFQLListener) ExitSelect(ctx *SelectContext) {}

// EnterTable is called when production table is entered.
func (s *BaseFQLListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseFQLListener) ExitTable(ctx *TableContext) {}
