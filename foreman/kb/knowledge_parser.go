// Code generated from Knowledge.g4 by ANTLR 4.7.2. DO NOT EDIT.

package kb // Knowledge
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 101,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 5, 2, 32, 10, 2, 3, 2, 3, 2, 5,
	2, 36, 10, 2, 3, 3, 3, 3, 3, 3, 7, 3, 41, 10, 3, 12, 3, 14, 3, 44, 11,
	3, 3, 4, 5, 4, 47, 10, 4, 3, 4, 3, 4, 5, 4, 51, 10, 4, 3, 5, 3, 5, 3, 5,
	7, 5, 56, 10, 5, 12, 5, 14, 5, 59, 11, 5, 3, 6, 5, 6, 62, 10, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 5, 6, 68, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 9, 5, 9, 77, 10, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 7, 12, 88, 10, 12, 12, 12, 14, 12, 91, 11, 12, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 2, 2, 16, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 2, 5, 3, 2, 19, 20, 4, 2, 17,
	17, 19, 20, 3, 2, 6, 11, 2, 97, 2, 31, 3, 2, 2, 2, 4, 37, 3, 2, 2, 2, 6,
	46, 3, 2, 2, 2, 8, 52, 3, 2, 2, 2, 10, 61, 3, 2, 2, 2, 12, 69, 3, 2, 2,
	2, 14, 71, 3, 2, 2, 2, 16, 76, 3, 2, 2, 2, 18, 78, 3, 2, 2, 2, 20, 80,
	3, 2, 2, 2, 22, 82, 3, 2, 2, 2, 24, 94, 3, 2, 2, 2, 26, 96, 3, 2, 2, 2,
	28, 98, 3, 2, 2, 2, 30, 32, 7, 4, 2, 2, 31, 30, 3, 2, 2, 2, 31, 32, 3,
	2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 35, 5, 4, 3, 2, 34, 36, 7, 5, 2, 2, 35,
	34, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 3, 3, 2, 2, 2, 37, 42, 5, 6, 4,
	2, 38, 39, 7, 13, 2, 2, 39, 41, 5, 6, 4, 2, 40, 38, 3, 2, 2, 2, 41, 44,
	3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 5, 3, 2, 2, 2,
	44, 42, 3, 2, 2, 2, 45, 47, 7, 4, 2, 2, 46, 45, 3, 2, 2, 2, 46, 47, 3,
	2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 50, 5, 8, 5, 2, 49, 51, 7, 5, 2, 2, 50,
	49, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 7, 3, 2, 2, 2, 52, 57, 5, 10, 6,
	2, 53, 54, 7, 12, 2, 2, 54, 56, 5, 10, 6, 2, 55, 53, 3, 2, 2, 2, 56, 59,
	3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 9, 3, 2, 2, 2,
	59, 57, 3, 2, 2, 2, 60, 62, 7, 4, 2, 2, 61, 60, 3, 2, 2, 2, 61, 62, 3,
	2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 5, 12, 7, 2, 64, 65, 5, 28, 15, 2,
	65, 67, 5, 14, 8, 2, 66, 68, 7, 5, 2, 2, 67, 66, 3, 2, 2, 2, 67, 68, 3,
	2, 2, 2, 68, 11, 3, 2, 2, 2, 69, 70, 5, 16, 9, 2, 70, 13, 3, 2, 2, 2, 71,
	72, 5, 16, 9, 2, 72, 15, 3, 2, 2, 2, 73, 77, 5, 18, 10, 2, 74, 77, 5, 20,
	11, 2, 75, 77, 5, 22, 12, 2, 76, 73, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76,
	75, 3, 2, 2, 2, 77, 17, 3, 2, 2, 2, 78, 79, 9, 2, 2, 2, 79, 19, 3, 2, 2,
	2, 80, 81, 7, 17, 2, 2, 81, 21, 3, 2, 2, 2, 82, 83, 5, 24, 13, 2, 83, 84,
	7, 4, 2, 2, 84, 89, 5, 26, 14, 2, 85, 86, 7, 14, 2, 2, 86, 88, 5, 26, 14,
	2, 87, 85, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90,
	3, 2, 2, 2, 90, 92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 93, 7, 5, 2, 2,
	93, 23, 3, 2, 2, 2, 94, 95, 7, 17, 2, 2, 95, 25, 3, 2, 2, 2, 96, 97, 9,
	3, 2, 2, 97, 27, 3, 2, 2, 2, 98, 99, 9, 4, 2, 2, 99, 29, 3, 2, 2, 2, 12,
	31, 35, 42, 46, 50, 57, 61, 67, 76, 89,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'*'", "'('", "')'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='",
	"'&'", "'|'", "','", "';'",
}
var symbolicNames = []string{
	"", "ASTERISK", "BS", "BE", "DEQ", "NEQ", "LT", "LE", "GT", "GE", "AND",
	"OR", "COMMA", "SEMICOLON", "WS", "IDENTIFIER", "STRING", "NUMBER", "REAL",
}

var ruleNames = []string{
	"knowledge", "clauses", "clause", "formulas", "formula", "leftOperand",
	"rightOperand", "operand", "literalOperand", "variableOperand", "functionOperand",
	"functionName", "parameter", "operator",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type KnowledgeParser struct {
	*antlr.BaseParser
}

func NewKnowledgeParser(input antlr.TokenStream) *KnowledgeParser {
	this := new(KnowledgeParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Knowledge.g4"

	return this
}

// KnowledgeParser tokens.
const (
	KnowledgeParserEOF        = antlr.TokenEOF
	KnowledgeParserASTERISK   = 1
	KnowledgeParserBS         = 2
	KnowledgeParserBE         = 3
	KnowledgeParserDEQ        = 4
	KnowledgeParserNEQ        = 5
	KnowledgeParserLT         = 6
	KnowledgeParserLE         = 7
	KnowledgeParserGT         = 8
	KnowledgeParserGE         = 9
	KnowledgeParserAND        = 10
	KnowledgeParserOR         = 11
	KnowledgeParserCOMMA      = 12
	KnowledgeParserSEMICOLON  = 13
	KnowledgeParserWS         = 14
	KnowledgeParserIDENTIFIER = 15
	KnowledgeParserSTRING     = 16
	KnowledgeParserNUMBER     = 17
	KnowledgeParserREAL       = 18
)

// KnowledgeParser rules.
const (
	KnowledgeParserRULE_knowledge       = 0
	KnowledgeParserRULE_clauses         = 1
	KnowledgeParserRULE_clause          = 2
	KnowledgeParserRULE_formulas        = 3
	KnowledgeParserRULE_formula         = 4
	KnowledgeParserRULE_leftOperand     = 5
	KnowledgeParserRULE_rightOperand    = 6
	KnowledgeParserRULE_operand         = 7
	KnowledgeParserRULE_literalOperand  = 8
	KnowledgeParserRULE_variableOperand = 9
	KnowledgeParserRULE_functionOperand = 10
	KnowledgeParserRULE_functionName    = 11
	KnowledgeParserRULE_parameter       = 12
	KnowledgeParserRULE_operator        = 13
)

// IKnowledgeContext is an interface to support dynamic dispatch.
type IKnowledgeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKnowledgeContext differentiates from other interfaces.
	IsKnowledgeContext()
}

type KnowledgeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKnowledgeContext() *KnowledgeContext {
	var p = new(KnowledgeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnowledgeParserRULE_knowledge
	return p
}

func (*KnowledgeContext) IsKnowledgeContext() {}

func NewKnowledgeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KnowledgeContext {
	var p = new(KnowledgeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnowledgeParserRULE_knowledge

	return p
}

func (s *KnowledgeContext) GetParser() antlr.Parser { return s.parser }

func (s *KnowledgeContext) Clauses() IClausesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClausesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClausesContext)
}

func (s *KnowledgeContext) BS() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserBS, 0)
}

func (s *KnowledgeContext) BE() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserBE, 0)
}

func (s *KnowledgeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KnowledgeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KnowledgeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.EnterKnowledge(s)
	}
}

func (s *KnowledgeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.ExitKnowledge(s)
	}
}

func (p *KnowledgeParser) Knowledge() (localctx IKnowledgeContext) {
	localctx = NewKnowledgeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, KnowledgeParserRULE_knowledge)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(28)
			p.Match(KnowledgeParserBS)
		}

	}
	{
		p.SetState(31)
		p.Clauses()
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KnowledgeParserBE {
		{
			p.SetState(32)
			p.Match(KnowledgeParserBE)
		}

	}

	return localctx
}

// IClausesContext is an interface to support dynamic dispatch.
type IClausesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClausesContext differentiates from other interfaces.
	IsClausesContext()
}

type ClausesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClausesContext() *ClausesContext {
	var p = new(ClausesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnowledgeParserRULE_clauses
	return p
}

func (*ClausesContext) IsClausesContext() {}

func NewClausesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClausesContext {
	var p = new(ClausesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnowledgeParserRULE_clauses

	return p
}

func (s *ClausesContext) GetParser() antlr.Parser { return s.parser }

func (s *ClausesContext) AllClause() []IClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClauseContext)(nil)).Elem())
	var tst = make([]IClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClauseContext)
		}
	}

	return tst
}

func (s *ClausesContext) Clause(i int) IClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClauseContext)
}

func (s *ClausesContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(KnowledgeParserOR)
}

func (s *ClausesContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(KnowledgeParserOR, i)
}

func (s *ClausesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClausesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClausesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.EnterClauses(s)
	}
}

func (s *ClausesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.ExitClauses(s)
	}
}

func (p *KnowledgeParser) Clauses() (localctx IClausesContext) {
	localctx = NewClausesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, KnowledgeParserRULE_clauses)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Clause()
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KnowledgeParserOR {
		{
			p.SetState(36)
			p.Match(KnowledgeParserOR)
		}
		{
			p.SetState(37)
			p.Clause()
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IClauseContext is an interface to support dynamic dispatch.
type IClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClauseContext differentiates from other interfaces.
	IsClauseContext()
}

type ClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseContext() *ClauseContext {
	var p = new(ClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnowledgeParserRULE_clause
	return p
}

func (*ClauseContext) IsClauseContext() {}

func NewClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseContext {
	var p = new(ClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnowledgeParserRULE_clause

	return p
}

func (s *ClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseContext) Formulas() IFormulasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormulasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormulasContext)
}

func (s *ClauseContext) BS() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserBS, 0)
}

func (s *ClauseContext) BE() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserBE, 0)
}

func (s *ClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.EnterClause(s)
	}
}

func (s *ClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.ExitClause(s)
	}
}

func (p *KnowledgeParser) Clause() (localctx IClauseContext) {
	localctx = NewClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, KnowledgeParserRULE_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(44)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(43)
			p.Match(KnowledgeParserBS)
		}

	}
	{
		p.SetState(46)
		p.Formulas()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(47)
			p.Match(KnowledgeParserBE)
		}

	}

	return localctx
}

// IFormulasContext is an interface to support dynamic dispatch.
type IFormulasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormulasContext differentiates from other interfaces.
	IsFormulasContext()
}

type FormulasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormulasContext() *FormulasContext {
	var p = new(FormulasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnowledgeParserRULE_formulas
	return p
}

func (*FormulasContext) IsFormulasContext() {}

func NewFormulasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormulasContext {
	var p = new(FormulasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnowledgeParserRULE_formulas

	return p
}

func (s *FormulasContext) GetParser() antlr.Parser { return s.parser }

func (s *FormulasContext) AllFormula() []IFormulaContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormulaContext)(nil)).Elem())
	var tst = make([]IFormulaContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormulaContext)
		}
	}

	return tst
}

func (s *FormulasContext) Formula(i int) IFormulaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormulaContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormulaContext)
}

func (s *FormulasContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(KnowledgeParserAND)
}

func (s *FormulasContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(KnowledgeParserAND, i)
}

func (s *FormulasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormulasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormulasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.EnterFormulas(s)
	}
}

func (s *FormulasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.ExitFormulas(s)
	}
}

func (p *KnowledgeParser) Formulas() (localctx IFormulasContext) {
	localctx = NewFormulasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, KnowledgeParserRULE_formulas)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Formula()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KnowledgeParserAND {
		{
			p.SetState(51)
			p.Match(KnowledgeParserAND)
		}
		{
			p.SetState(52)
			p.Formula()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFormulaContext is an interface to support dynamic dispatch.
type IFormulaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormulaContext differentiates from other interfaces.
	IsFormulaContext()
}

type FormulaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormulaContext() *FormulaContext {
	var p = new(FormulaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnowledgeParserRULE_formula
	return p
}

func (*FormulaContext) IsFormulaContext() {}

func NewFormulaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormulaContext {
	var p = new(FormulaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnowledgeParserRULE_formula

	return p
}

func (s *FormulaContext) GetParser() antlr.Parser { return s.parser }

func (s *FormulaContext) LeftOperand() ILeftOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILeftOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILeftOperandContext)
}

func (s *FormulaContext) Operator() IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *FormulaContext) RightOperand() IRightOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRightOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRightOperandContext)
}

func (s *FormulaContext) BS() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserBS, 0)
}

func (s *FormulaContext) BE() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserBE, 0)
}

func (s *FormulaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormulaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormulaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.EnterFormula(s)
	}
}

func (s *FormulaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.ExitFormula(s)
	}
}

func (p *KnowledgeParser) Formula() (localctx IFormulaContext) {
	localctx = NewFormulaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, KnowledgeParserRULE_formula)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KnowledgeParserBS {
		{
			p.SetState(58)
			p.Match(KnowledgeParserBS)
		}

	}
	{
		p.SetState(61)
		p.LeftOperand()
	}
	{
		p.SetState(62)
		p.Operator()
	}
	{
		p.SetState(63)
		p.RightOperand()
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(64)
			p.Match(KnowledgeParserBE)
		}

	}

	return localctx
}

// ILeftOperandContext is an interface to support dynamic dispatch.
type ILeftOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeftOperandContext differentiates from other interfaces.
	IsLeftOperandContext()
}

type LeftOperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeftOperandContext() *LeftOperandContext {
	var p = new(LeftOperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnowledgeParserRULE_leftOperand
	return p
}

func (*LeftOperandContext) IsLeftOperandContext() {}

func NewLeftOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftOperandContext {
	var p = new(LeftOperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnowledgeParserRULE_leftOperand

	return p
}

func (s *LeftOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *LeftOperandContext) Operand() IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *LeftOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeftOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeftOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.EnterLeftOperand(s)
	}
}

func (s *LeftOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.ExitLeftOperand(s)
	}
}

func (p *KnowledgeParser) LeftOperand() (localctx ILeftOperandContext) {
	localctx = NewLeftOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, KnowledgeParserRULE_leftOperand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Operand()
	}

	return localctx
}

// IRightOperandContext is an interface to support dynamic dispatch.
type IRightOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRightOperandContext differentiates from other interfaces.
	IsRightOperandContext()
}

type RightOperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRightOperandContext() *RightOperandContext {
	var p = new(RightOperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnowledgeParserRULE_rightOperand
	return p
}

func (*RightOperandContext) IsRightOperandContext() {}

func NewRightOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RightOperandContext {
	var p = new(RightOperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnowledgeParserRULE_rightOperand

	return p
}

func (s *RightOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *RightOperandContext) Operand() IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *RightOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RightOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RightOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.EnterRightOperand(s)
	}
}

func (s *RightOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.ExitRightOperand(s)
	}
}

func (p *KnowledgeParser) RightOperand() (localctx IRightOperandContext) {
	localctx = NewRightOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, KnowledgeParserRULE_rightOperand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Operand()
	}

	return localctx
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnowledgeParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnowledgeParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) LiteralOperand() ILiteralOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralOperandContext)
}

func (s *OperandContext) VariableOperand() IVariableOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableOperandContext)
}

func (s *OperandContext) FunctionOperand() IFunctionOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionOperandContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (p *KnowledgeParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, KnowledgeParserRULE_operand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.LiteralOperand()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.VariableOperand()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.FunctionOperand()
		}

	}

	return localctx
}

// ILiteralOperandContext is an interface to support dynamic dispatch.
type ILiteralOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralOperandContext differentiates from other interfaces.
	IsLiteralOperandContext()
}

type LiteralOperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralOperandContext() *LiteralOperandContext {
	var p = new(LiteralOperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnowledgeParserRULE_literalOperand
	return p
}

func (*LiteralOperandContext) IsLiteralOperandContext() {}

func NewLiteralOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralOperandContext {
	var p = new(LiteralOperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnowledgeParserRULE_literalOperand

	return p
}

func (s *LiteralOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralOperandContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserNUMBER, 0)
}

func (s *LiteralOperandContext) REAL() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserREAL, 0)
}

func (s *LiteralOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.EnterLiteralOperand(s)
	}
}

func (s *LiteralOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.ExitLiteralOperand(s)
	}
}

func (p *KnowledgeParser) LiteralOperand() (localctx ILiteralOperandContext) {
	localctx = NewLiteralOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, KnowledgeParserRULE_literalOperand)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KnowledgeParserNUMBER || _la == KnowledgeParserREAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IVariableOperandContext is an interface to support dynamic dispatch.
type IVariableOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableOperandContext differentiates from other interfaces.
	IsVariableOperandContext()
}

type VariableOperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableOperandContext() *VariableOperandContext {
	var p = new(VariableOperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnowledgeParserRULE_variableOperand
	return p
}

func (*VariableOperandContext) IsVariableOperandContext() {}

func NewVariableOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableOperandContext {
	var p = new(VariableOperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnowledgeParserRULE_variableOperand

	return p
}

func (s *VariableOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableOperandContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserIDENTIFIER, 0)
}

func (s *VariableOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.EnterVariableOperand(s)
	}
}

func (s *VariableOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.ExitVariableOperand(s)
	}
}

func (p *KnowledgeParser) VariableOperand() (localctx IVariableOperandContext) {
	localctx = NewVariableOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, KnowledgeParserRULE_variableOperand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(KnowledgeParserIDENTIFIER)
	}

	return localctx
}

// IFunctionOperandContext is an interface to support dynamic dispatch.
type IFunctionOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionOperandContext differentiates from other interfaces.
	IsFunctionOperandContext()
}

type FunctionOperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionOperandContext() *FunctionOperandContext {
	var p = new(FunctionOperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnowledgeParserRULE_functionOperand
	return p
}

func (*FunctionOperandContext) IsFunctionOperandContext() {}

func NewFunctionOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionOperandContext {
	var p = new(FunctionOperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnowledgeParserRULE_functionOperand

	return p
}

func (s *FunctionOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionOperandContext) FunctionName() IFunctionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionNameContext)
}

func (s *FunctionOperandContext) BS() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserBS, 0)
}

func (s *FunctionOperandContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *FunctionOperandContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *FunctionOperandContext) BE() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserBE, 0)
}

func (s *FunctionOperandContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KnowledgeParserCOMMA)
}

func (s *FunctionOperandContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KnowledgeParserCOMMA, i)
}

func (s *FunctionOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.EnterFunctionOperand(s)
	}
}

func (s *FunctionOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.ExitFunctionOperand(s)
	}
}

func (p *KnowledgeParser) FunctionOperand() (localctx IFunctionOperandContext) {
	localctx = NewFunctionOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, KnowledgeParserRULE_functionOperand)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.FunctionName()
	}
	{
		p.SetState(81)
		p.Match(KnowledgeParserBS)
	}
	{
		p.SetState(82)
		p.Parameter()
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KnowledgeParserCOMMA {
		{
			p.SetState(83)
			p.Match(KnowledgeParserCOMMA)
		}
		{
			p.SetState(84)
			p.Parameter()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(KnowledgeParserBE)
	}

	return localctx
}

// IFunctionNameContext is an interface to support dynamic dispatch.
type IFunctionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionNameContext differentiates from other interfaces.
	IsFunctionNameContext()
}

type FunctionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionNameContext() *FunctionNameContext {
	var p = new(FunctionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnowledgeParserRULE_functionName
	return p
}

func (*FunctionNameContext) IsFunctionNameContext() {}

func NewFunctionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionNameContext {
	var p = new(FunctionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnowledgeParserRULE_functionName

	return p
}

func (s *FunctionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserIDENTIFIER, 0)
}

func (s *FunctionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.EnterFunctionName(s)
	}
}

func (s *FunctionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.ExitFunctionName(s)
	}
}

func (p *KnowledgeParser) FunctionName() (localctx IFunctionNameContext) {
	localctx = NewFunctionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, KnowledgeParserRULE_functionName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(KnowledgeParserIDENTIFIER)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnowledgeParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnowledgeParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserIDENTIFIER, 0)
}

func (s *ParameterContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserNUMBER, 0)
}

func (s *ParameterContext) REAL() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserREAL, 0)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *KnowledgeParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, KnowledgeParserRULE_parameter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KnowledgeParserIDENTIFIER)|(1<<KnowledgeParserNUMBER)|(1<<KnowledgeParserREAL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnowledgeParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnowledgeParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) DEQ() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserDEQ, 0)
}

func (s *OperatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserNEQ, 0)
}

func (s *OperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserLT, 0)
}

func (s *OperatorContext) LE() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserLE, 0)
}

func (s *OperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserGT, 0)
}

func (s *OperatorContext) GE() antlr.TerminalNode {
	return s.GetToken(KnowledgeParserGE, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnowledgeListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *KnowledgeParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, KnowledgeParserRULE_operator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KnowledgeParserDEQ)|(1<<KnowledgeParserNEQ)|(1<<KnowledgeParserLT)|(1<<KnowledgeParserLE)|(1<<KnowledgeParserGT)|(1<<KnowledgeParserGE))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
