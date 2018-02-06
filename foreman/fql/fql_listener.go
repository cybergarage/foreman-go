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

	// EnterCondition is called when entering the Condition production.
	EnterCondition(c *ConditionContext)

	// EnterSelect_condition is called when entering the select_condition production.
	EnterSelect_condition(c *Select_conditionContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterLeftOperand is called when entering the LeftOperand production.
	EnterLeftOperand(c *LeftOperandContext)

	// EnterRightOperand is called when entering the RightOperand production.
	EnterRightOperand(c *RightOperandContext)

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

	// ExitCondition is called when exiting the Condition production.
	ExitCondition(c *ConditionContext)

	// ExitSelect_condition is called when exiting the select_condition production.
	ExitSelect_condition(c *Select_conditionContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitLeftOperand is called when exiting the LeftOperand production.
	ExitLeftOperand(c *LeftOperandContext)

	// ExitRightOperand is called when exiting the RightOperand production.
	ExitRightOperand(c *RightOperandContext)

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
