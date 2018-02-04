// Code generated from FQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

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

// EnterInsertQuery is called when production InsertQuery is entered.
func (s *BaseFQLListener) EnterInsertQuery(ctx *InsertQueryContext) {}

// ExitInsertQuery is called when production InsertQuery is exited.
func (s *BaseFQLListener) ExitInsertQuery(ctx *InsertQueryContext) {}

// EnterSetQuery is called when production SetQuery is entered.
func (s *BaseFQLListener) EnterSetQuery(ctx *SetQueryContext) {}

// ExitSetQuery is called when production SetQuery is exited.
func (s *BaseFQLListener) ExitSetQuery(ctx *SetQueryContext) {}

// EnterSelectQuery is called when production SelectQuery is entered.
func (s *BaseFQLListener) EnterSelectQuery(ctx *SelectQueryContext) {}

// ExitSelectQuery is called when production SelectQuery is exited.
func (s *BaseFQLListener) ExitSelectQuery(ctx *SelectQueryContext) {}

// EnterExportQuery is called when production ExportQuery is entered.
func (s *BaseFQLListener) EnterExportQuery(ctx *ExportQueryContext) {}

// ExitExportQuery is called when production ExportQuery is exited.
func (s *BaseFQLListener) ExitExportQuery(ctx *ExportQueryContext) {}

// EnterDeleteQuery is called when production DeleteQuery is entered.
func (s *BaseFQLListener) EnterDeleteQuery(ctx *DeleteQueryContext) {}

// ExitDeleteQuery is called when production DeleteQuery is exited.
func (s *BaseFQLListener) ExitDeleteQuery(ctx *DeleteQueryContext) {}

// EnterTarget is called when production target is entered.
func (s *BaseFQLListener) EnterTarget(ctx *TargetContext) {}

// ExitTarget is called when production target is exited.
func (s *BaseFQLListener) ExitTarget(ctx *TargetContext) {}

// EnterValue is called when production value is entered.
func (s *BaseFQLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseFQLListener) ExitValue(ctx *ValueContext) {}

// EnterValues is called when production values is entered.
func (s *BaseFQLListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production values is exited.
func (s *BaseFQLListener) ExitValues(ctx *ValuesContext) {}
