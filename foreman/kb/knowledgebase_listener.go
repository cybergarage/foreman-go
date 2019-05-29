// Code generated from knowledgebase.g4 by ANTLR 4.7.2. DO NOT EDIT.

package kb // knowledgebase
import "github.com/antlr/antlr4/runtime/Go/antlr"

// knowledgebaseListener is a complete listener for a parse tree produced by knowledgebaseParser.
type knowledgebaseListener interface {
	antlr.ParseTreeListener

	// EnterKnowledgebase is called when entering the knowledgebase production.
	EnterKnowledgebase(c *KnowledgebaseContext)

	// EnterClauses is called when entering the clauses production.
	EnterClauses(c *ClausesContext)

	// EnterClause is called when entering the clause production.
	EnterClause(c *ClauseContext)

	// EnterFormulas is called when entering the formulas production.
	EnterFormulas(c *FormulasContext)

	// EnterFormula is called when entering the formula production.
	EnterFormula(c *FormulaContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// ExitKnowledgebase is called when exiting the knowledgebase production.
	ExitKnowledgebase(c *KnowledgebaseContext)

	// ExitClauses is called when exiting the clauses production.
	ExitClauses(c *ClausesContext)

	// ExitClause is called when exiting the clause production.
	ExitClause(c *ClauseContext)

	// ExitFormulas is called when exiting the formulas production.
	ExitFormulas(c *FormulasContext)

	// ExitFormula is called when exiting the formula production.
	ExitFormula(c *FormulaContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)
}
