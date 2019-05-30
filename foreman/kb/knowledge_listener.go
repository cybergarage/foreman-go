// Code generated from Knowledge.g4 by ANTLR 4.7.2. DO NOT EDIT.

package kb // Knowledge
import "github.com/antlr/antlr4/runtime/Go/antlr"

// KnowledgeListener is a complete listener for a parse tree produced by KnowledgeParser.
type KnowledgeListener interface {
	antlr.ParseTreeListener

	// EnterKnowledge is called when entering the knowledge production.
	EnterKnowledge(c *KnowledgeContext)

	// EnterClauses is called when entering the clauses production.
	EnterClauses(c *ClausesContext)

	// EnterClause is called when entering the clause production.
	EnterClause(c *ClauseContext)

	// EnterFormulas is called when entering the formulas production.
	EnterFormulas(c *FormulasContext)

	// EnterFormula is called when entering the formula production.
	EnterFormula(c *FormulaContext)

	// EnterLeftOperand is called when entering the leftOperand production.
	EnterLeftOperand(c *LeftOperandContext)

	// EnterRightOperand is called when entering the rightOperand production.
	EnterRightOperand(c *RightOperandContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// ExitKnowledge is called when exiting the knowledge production.
	ExitKnowledge(c *KnowledgeContext)

	// ExitClauses is called when exiting the clauses production.
	ExitClauses(c *ClausesContext)

	// ExitClause is called when exiting the clause production.
	ExitClause(c *ClauseContext)

	// ExitFormulas is called when exiting the formulas production.
	ExitFormulas(c *FormulasContext)

	// ExitFormula is called when exiting the formula production.
	ExitFormula(c *FormulaContext)

	// ExitLeftOperand is called when exiting the leftOperand production.
	ExitLeftOperand(c *LeftOperandContext)

	// ExitRightOperand is called when exiting the rightOperand production.
	ExitRightOperand(c *RightOperandContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)
}
