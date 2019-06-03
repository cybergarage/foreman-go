// Code generated from Knowledge.g4 by ANTLR 4.7.2. DO NOT EDIT.

package kb // Knowledge
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseKnowledgeListener is a complete listener for a parse tree produced by KnowledgeParser.
type BaseKnowledgeListener struct{}

var _ KnowledgeListener = &BaseKnowledgeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseKnowledgeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseKnowledgeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseKnowledgeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseKnowledgeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterKnowledge is called when production knowledge is entered.
func (s *BaseKnowledgeListener) EnterKnowledge(ctx *KnowledgeContext) {}

// ExitKnowledge is called when production knowledge is exited.
func (s *BaseKnowledgeListener) ExitKnowledge(ctx *KnowledgeContext) {}

// EnterClauses is called when production clauses is entered.
func (s *BaseKnowledgeListener) EnterClauses(ctx *ClausesContext) {}

// ExitClauses is called when production clauses is exited.
func (s *BaseKnowledgeListener) ExitClauses(ctx *ClausesContext) {}

// EnterClause is called when production clause is entered.
func (s *BaseKnowledgeListener) EnterClause(ctx *ClauseContext) {}

// ExitClause is called when production clause is exited.
func (s *BaseKnowledgeListener) ExitClause(ctx *ClauseContext) {}

// EnterFormulas is called when production formulas is entered.
func (s *BaseKnowledgeListener) EnterFormulas(ctx *FormulasContext) {}

// ExitFormulas is called when production formulas is exited.
func (s *BaseKnowledgeListener) ExitFormulas(ctx *FormulasContext) {}

// EnterFormula is called when production formula is entered.
func (s *BaseKnowledgeListener) EnterFormula(ctx *FormulaContext) {}

// ExitFormula is called when production formula is exited.
func (s *BaseKnowledgeListener) ExitFormula(ctx *FormulaContext) {}

// EnterLeftOperand is called when production leftOperand is entered.
func (s *BaseKnowledgeListener) EnterLeftOperand(ctx *LeftOperandContext) {}

// ExitLeftOperand is called when production leftOperand is exited.
func (s *BaseKnowledgeListener) ExitLeftOperand(ctx *LeftOperandContext) {}

// EnterRightOperand is called when production rightOperand is entered.
func (s *BaseKnowledgeListener) EnterRightOperand(ctx *RightOperandContext) {}

// ExitRightOperand is called when production rightOperand is exited.
func (s *BaseKnowledgeListener) ExitRightOperand(ctx *RightOperandContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseKnowledgeListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseKnowledgeListener) ExitOperand(ctx *OperandContext) {}

// EnterLiteralOperand is called when production literalOperand is entered.
func (s *BaseKnowledgeListener) EnterLiteralOperand(ctx *LiteralOperandContext) {}

// ExitLiteralOperand is called when production literalOperand is exited.
func (s *BaseKnowledgeListener) ExitLiteralOperand(ctx *LiteralOperandContext) {}

// EnterVariableOperand is called when production variableOperand is entered.
func (s *BaseKnowledgeListener) EnterVariableOperand(ctx *VariableOperandContext) {}

// ExitVariableOperand is called when production variableOperand is exited.
func (s *BaseKnowledgeListener) ExitVariableOperand(ctx *VariableOperandContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseKnowledgeListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseKnowledgeListener) ExitOperator(ctx *OperatorContext) {}
