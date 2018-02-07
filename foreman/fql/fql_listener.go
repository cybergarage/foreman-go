// Code generated from FQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package fql // FQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// FQLListener is a complete listener for a parse tree produced by FQLParser.
type FQLListener interface {
	antlr.ParseTreeListener

	// EnterFql is called when entering the fql production.
	EnterFql(c *FqlContext)

	// EnterQueryList is called when entering the queryList production.
	EnterQueryList(c *QueryListContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterInsertQuery is called when entering the insertQuery production.
	EnterInsertQuery(c *InsertQueryContext)

	// EnterSetQuery is called when entering the setQuery production.
	EnterSetQuery(c *SetQueryContext)

	// EnterSelectQuery is called when entering the selectQuery production.
	EnterSelectQuery(c *SelectQueryContext)

	// EnterExportQuery is called when entering the exportQuery production.
	EnterExportQuery(c *ExportQueryContext)

	// EnterDeleteQuery is called when entering the deleteQuery production.
	EnterDeleteQuery(c *DeleteQueryContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterValues is called when entering the values production.
	EnterValues(c *ValuesContext)

	// EnterColumn is called when entering the column production.
	EnterColumn(c *ColumnContext)

	// EnterColumns is called when entering the columns production.
	EnterColumns(c *ColumnsContext)

	// EnterConditions is called when entering the conditions production.
	EnterConditions(c *ConditionsContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterLeftOperand is called when entering the leftOperand production.
	EnterLeftOperand(c *LeftOperandContext)

	// EnterRightOperand is called when entering the rightOperand production.
	EnterRightOperand(c *RightOperandContext)

	// ExitFql is called when exiting the fql production.
	ExitFql(c *FqlContext)

	// ExitQueryList is called when exiting the queryList production.
	ExitQueryList(c *QueryListContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitInsertQuery is called when exiting the insertQuery production.
	ExitInsertQuery(c *InsertQueryContext)

	// ExitSetQuery is called when exiting the setQuery production.
	ExitSetQuery(c *SetQueryContext)

	// ExitSelectQuery is called when exiting the selectQuery production.
	ExitSelectQuery(c *SelectQueryContext)

	// ExitExportQuery is called when exiting the exportQuery production.
	ExitExportQuery(c *ExportQueryContext)

	// ExitDeleteQuery is called when exiting the deleteQuery production.
	ExitDeleteQuery(c *DeleteQueryContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitValues is called when exiting the values production.
	ExitValues(c *ValuesContext)

	// ExitColumn is called when exiting the column production.
	ExitColumn(c *ColumnContext)

	// ExitColumns is called when exiting the columns production.
	ExitColumns(c *ColumnsContext)

	// ExitConditions is called when exiting the conditions production.
	ExitConditions(c *ConditionsContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitLeftOperand is called when exiting the leftOperand production.
	ExitLeftOperand(c *LeftOperandContext)

	// ExitRightOperand is called when exiting the rightOperand production.
	ExitRightOperand(c *RightOperandContext)
}
