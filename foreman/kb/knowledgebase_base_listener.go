// Code generated from knowledgebase.g4 by ANTLR 4.7.2. DO NOT EDIT.

package kb // knowledgebase
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseknowledgebaseListener is a complete listener for a parse tree produced by knowledgebaseParser.
type BaseknowledgebaseListener struct{}

var _ knowledgebaseListener = &BaseknowledgebaseListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseknowledgebaseListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseknowledgebaseListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseknowledgebaseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseknowledgebaseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterKnowledge is called when production knowledge is entered.
func (s *BaseknowledgebaseListener) EnterKnowledge(ctx *KnowledgeContext) {}

// ExitKnowledge is called when production knowledge is exited.
func (s *BaseknowledgebaseListener) ExitKnowledge(ctx *KnowledgeContext) {}

// EnterClauses is called when production clauses is entered.
func (s *BaseknowledgebaseListener) EnterClauses(ctx *ClausesContext) {}

// ExitClauses is called when production clauses is exited.
func (s *BaseknowledgebaseListener) ExitClauses(ctx *ClausesContext) {}

// EnterClause is called when production clause is entered.
func (s *BaseknowledgebaseListener) EnterClause(ctx *ClauseContext) {}

// ExitClause is called when production clause is exited.
func (s *BaseknowledgebaseListener) ExitClause(ctx *ClauseContext) {}

// EnterFormulas is called when production formulas is entered.
func (s *BaseknowledgebaseListener) EnterFormulas(ctx *FormulasContext) {}

// ExitFormulas is called when production formulas is exited.
func (s *BaseknowledgebaseListener) ExitFormulas(ctx *FormulasContext) {}

// EnterFormula is called when production formula is entered.
func (s *BaseknowledgebaseListener) EnterFormula(ctx *FormulaContext) {}

// ExitFormula is called when production formula is exited.
func (s *BaseknowledgebaseListener) ExitFormula(ctx *FormulaContext) {}

// EnterLeftOperand is called when production leftOperand is entered.
func (s *BaseknowledgebaseListener) EnterLeftOperand(ctx *LeftOperandContext) {}

// ExitLeftOperand is called when production leftOperand is exited.
func (s *BaseknowledgebaseListener) ExitLeftOperand(ctx *LeftOperandContext) {}

// EnterRightOperand is called when production rightOperand is entered.
func (s *BaseknowledgebaseListener) EnterRightOperand(ctx *RightOperandContext) {}

// ExitRightOperand is called when production rightOperand is exited.
func (s *BaseknowledgebaseListener) ExitRightOperand(ctx *RightOperandContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseknowledgebaseListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseknowledgebaseListener) ExitOperand(ctx *OperandContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseknowledgebaseListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseknowledgebaseListener) ExitOperator(ctx *OperatorContext) {}
