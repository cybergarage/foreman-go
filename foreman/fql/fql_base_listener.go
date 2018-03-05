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

// EnterQueryList is called when production queryList is entered.
func (s *BaseFQLListener) EnterQueryList(ctx *QueryListContext) {}

// ExitQueryList is called when production queryList is exited.
func (s *BaseFQLListener) ExitQueryList(ctx *QueryListContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseFQLListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseFQLListener) ExitQuery(ctx *QueryContext) {}

// EnterInsertQuery is called when production insertQuery is entered.
func (s *BaseFQLListener) EnterInsertQuery(ctx *InsertQueryContext) {}

// ExitInsertQuery is called when production insertQuery is exited.
func (s *BaseFQLListener) ExitInsertQuery(ctx *InsertQueryContext) {}

// EnterSetQuery is called when production setQuery is entered.
func (s *BaseFQLListener) EnterSetQuery(ctx *SetQueryContext) {}

// ExitSetQuery is called when production setQuery is exited.
func (s *BaseFQLListener) ExitSetQuery(ctx *SetQueryContext) {}

// EnterSelectQuery is called when production selectQuery is entered.
func (s *BaseFQLListener) EnterSelectQuery(ctx *SelectQueryContext) {}

// ExitSelectQuery is called when production selectQuery is exited.
func (s *BaseFQLListener) ExitSelectQuery(ctx *SelectQueryContext) {}

// EnterExportQuery is called when production exportQuery is entered.
func (s *BaseFQLListener) EnterExportQuery(ctx *ExportQueryContext) {}

// ExitExportQuery is called when production exportQuery is exited.
func (s *BaseFQLListener) ExitExportQuery(ctx *ExportQueryContext) {}

// EnterDeleteQuery is called when production deleteQuery is entered.
func (s *BaseFQLListener) EnterDeleteQuery(ctx *DeleteQueryContext) {}

// ExitDeleteQuery is called when production deleteQuery is exited.
func (s *BaseFQLListener) ExitDeleteQuery(ctx *DeleteQueryContext) {}

// EnterAnalyzeQuery is called when production analyzeQuery is entered.
func (s *BaseFQLListener) EnterAnalyzeQuery(ctx *AnalyzeQueryContext) {}

// ExitAnalyzeQuery is called when production analyzeQuery is exited.
func (s *BaseFQLListener) ExitAnalyzeQuery(ctx *AnalyzeQueryContext) {}

// EnterExecuteQuery is called when production executeQuery is entered.
func (s *BaseFQLListener) EnterExecuteQuery(ctx *ExecuteQueryContext) {}

// ExitExecuteQuery is called when production executeQuery is exited.
func (s *BaseFQLListener) ExitExecuteQuery(ctx *ExecuteQueryContext) {}

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

// EnterColumn is called when production column is entered.
func (s *BaseFQLListener) EnterColumn(ctx *ColumnContext) {}

// ExitColumn is called when production column is exited.
func (s *BaseFQLListener) ExitColumn(ctx *ColumnContext) {}

// EnterColumns is called when production columns is entered.
func (s *BaseFQLListener) EnterColumns(ctx *ColumnsContext) {}

// ExitColumns is called when production columns is exited.
func (s *BaseFQLListener) ExitColumns(ctx *ColumnsContext) {}

// EnterConditions is called when production conditions is entered.
func (s *BaseFQLListener) EnterConditions(ctx *ConditionsContext) {}

// ExitConditions is called when production conditions is exited.
func (s *BaseFQLListener) ExitConditions(ctx *ConditionsContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseFQLListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseFQLListener) ExitCondition(ctx *ConditionContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseFQLListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseFQLListener) ExitOperator(ctx *OperatorContext) {}

// EnterLeftOperand is called when production leftOperand is entered.
func (s *BaseFQLListener) EnterLeftOperand(ctx *LeftOperandContext) {}

// ExitLeftOperand is called when production leftOperand is exited.
func (s *BaseFQLListener) ExitLeftOperand(ctx *LeftOperandContext) {}

// EnterRightOperand is called when production rightOperand is entered.
func (s *BaseFQLListener) EnterRightOperand(ctx *RightOperandContext) {}

// ExitRightOperand is called when production rightOperand is exited.
func (s *BaseFQLListener) ExitRightOperand(ctx *RightOperandContext) {}
