// Code generated from knowledgebase.g4 by ANTLR 4.7.2. DO NOT EDIT.

package kb // knowledgebase
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 68, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 5, 2, 22, 10, 2, 3, 2, 3, 2, 5,
	2, 26, 10, 2, 3, 3, 3, 3, 3, 3, 7, 3, 31, 10, 3, 12, 3, 14, 3, 34, 11,
	3, 3, 4, 5, 4, 37, 10, 4, 3, 4, 3, 4, 5, 4, 41, 10, 4, 3, 5, 3, 5, 3, 5,
	7, 5, 46, 10, 5, 12, 5, 14, 5, 49, 11, 5, 3, 6, 5, 6, 52, 10, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 5, 6, 58, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 4, 4,
	2, 17, 17, 19, 20, 3, 2, 6, 11, 2, 66, 2, 21, 3, 2, 2, 2, 4, 27, 3, 2,
	2, 2, 6, 36, 3, 2, 2, 2, 8, 42, 3, 2, 2, 2, 10, 51, 3, 2, 2, 2, 12, 59,
	3, 2, 2, 2, 14, 61, 3, 2, 2, 2, 16, 63, 3, 2, 2, 2, 18, 65, 3, 2, 2, 2,
	20, 22, 7, 4, 2, 2, 21, 20, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 23, 3,
	2, 2, 2, 23, 25, 5, 4, 3, 2, 24, 26, 7, 5, 2, 2, 25, 24, 3, 2, 2, 2, 25,
	26, 3, 2, 2, 2, 26, 3, 3, 2, 2, 2, 27, 32, 5, 6, 4, 2, 28, 29, 7, 13, 2,
	2, 29, 31, 5, 6, 4, 2, 30, 28, 3, 2, 2, 2, 31, 34, 3, 2, 2, 2, 32, 30,
	3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 5, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2,
	35, 37, 7, 4, 2, 2, 36, 35, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 38, 3,
	2, 2, 2, 38, 40, 5, 8, 5, 2, 39, 41, 7, 5, 2, 2, 40, 39, 3, 2, 2, 2, 40,
	41, 3, 2, 2, 2, 41, 7, 3, 2, 2, 2, 42, 47, 5, 10, 6, 2, 43, 44, 7, 12,
	2, 2, 44, 46, 5, 10, 6, 2, 45, 43, 3, 2, 2, 2, 46, 49, 3, 2, 2, 2, 47,
	45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 9, 3, 2, 2, 2, 49, 47, 3, 2, 2,
	2, 50, 52, 7, 4, 2, 2, 51, 50, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 53,
	3, 2, 2, 2, 53, 54, 5, 12, 7, 2, 54, 55, 5, 18, 10, 2, 55, 57, 5, 14, 8,
	2, 56, 58, 7, 5, 2, 2, 57, 56, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 11,
	3, 2, 2, 2, 59, 60, 5, 16, 9, 2, 60, 13, 3, 2, 2, 2, 61, 62, 5, 16, 9,
	2, 62, 15, 3, 2, 2, 2, 63, 64, 9, 2, 2, 2, 64, 17, 3, 2, 2, 2, 65, 66,
	9, 3, 2, 2, 66, 19, 3, 2, 2, 2, 10, 21, 25, 32, 36, 40, 47, 51, 57,
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
	"rightOperand", "operand", "operator",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type knowledgebaseParser struct {
	*antlr.BaseParser
}

func NewknowledgebaseParser(input antlr.TokenStream) *knowledgebaseParser {
	this := new(knowledgebaseParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "knowledgebase.g4"

	return this
}

// knowledgebaseParser tokens.
const (
	knowledgebaseParserEOF        = antlr.TokenEOF
	knowledgebaseParserASTERISK   = 1
	knowledgebaseParserBS         = 2
	knowledgebaseParserBE         = 3
	knowledgebaseParserDEQ        = 4
	knowledgebaseParserNEQ        = 5
	knowledgebaseParserLT         = 6
	knowledgebaseParserLE         = 7
	knowledgebaseParserGT         = 8
	knowledgebaseParserGE         = 9
	knowledgebaseParserAND        = 10
	knowledgebaseParserOR         = 11
	knowledgebaseParserCOMMA      = 12
	knowledgebaseParserSEMICOLON  = 13
	knowledgebaseParserWS         = 14
	knowledgebaseParserIDENTIFIER = 15
	knowledgebaseParserSTRING     = 16
	knowledgebaseParserNUMBER     = 17
	knowledgebaseParserREAL       = 18
)

// knowledgebaseParser rules.
const (
	knowledgebaseParserRULE_knowledge    = 0
	knowledgebaseParserRULE_clauses      = 1
	knowledgebaseParserRULE_clause       = 2
	knowledgebaseParserRULE_formulas     = 3
	knowledgebaseParserRULE_formula      = 4
	knowledgebaseParserRULE_leftOperand  = 5
	knowledgebaseParserRULE_rightOperand = 6
	knowledgebaseParserRULE_operand      = 7
	knowledgebaseParserRULE_operator     = 8
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
	p.RuleIndex = knowledgebaseParserRULE_knowledge
	return p
}

func (*KnowledgeContext) IsKnowledgeContext() {}

func NewKnowledgeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KnowledgeContext {
	var p = new(KnowledgeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = knowledgebaseParserRULE_knowledge

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
	return s.GetToken(knowledgebaseParserBS, 0)
}

func (s *KnowledgeContext) BE() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserBE, 0)
}

func (s *KnowledgeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KnowledgeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KnowledgeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.EnterKnowledge(s)
	}
}

func (s *KnowledgeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.ExitKnowledge(s)
	}
}

func (p *knowledgebaseParser) Knowledge() (localctx IKnowledgeContext) {
	localctx = NewKnowledgeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, knowledgebaseParserRULE_knowledge)
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(18)
			p.Match(knowledgebaseParserBS)
		}

	}
	{
		p.SetState(21)
		p.Clauses()
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == knowledgebaseParserBE {
		{
			p.SetState(22)
			p.Match(knowledgebaseParserBE)
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
	p.RuleIndex = knowledgebaseParserRULE_clauses
	return p
}

func (*ClausesContext) IsClausesContext() {}

func NewClausesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClausesContext {
	var p = new(ClausesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = knowledgebaseParserRULE_clauses

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
	return s.GetTokens(knowledgebaseParserOR)
}

func (s *ClausesContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserOR, i)
}

func (s *ClausesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClausesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClausesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.EnterClauses(s)
	}
}

func (s *ClausesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.ExitClauses(s)
	}
}

func (p *knowledgebaseParser) Clauses() (localctx IClausesContext) {
	localctx = NewClausesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, knowledgebaseParserRULE_clauses)
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
		p.SetState(25)
		p.Clause()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == knowledgebaseParserOR {
		{
			p.SetState(26)
			p.Match(knowledgebaseParserOR)
		}
		{
			p.SetState(27)
			p.Clause()
		}

		p.SetState(32)
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
	p.RuleIndex = knowledgebaseParserRULE_clause
	return p
}

func (*ClauseContext) IsClauseContext() {}

func NewClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseContext {
	var p = new(ClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = knowledgebaseParserRULE_clause

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
	return s.GetToken(knowledgebaseParserBS, 0)
}

func (s *ClauseContext) BE() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserBE, 0)
}

func (s *ClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.EnterClause(s)
	}
}

func (s *ClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.ExitClause(s)
	}
}

func (p *knowledgebaseParser) Clause() (localctx IClauseContext) {
	localctx = NewClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, knowledgebaseParserRULE_clause)

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
	p.SetState(34)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(33)
			p.Match(knowledgebaseParserBS)
		}

	}
	{
		p.SetState(36)
		p.Formulas()
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(37)
			p.Match(knowledgebaseParserBE)
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
	p.RuleIndex = knowledgebaseParserRULE_formulas
	return p
}

func (*FormulasContext) IsFormulasContext() {}

func NewFormulasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormulasContext {
	var p = new(FormulasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = knowledgebaseParserRULE_formulas

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
	return s.GetTokens(knowledgebaseParserAND)
}

func (s *FormulasContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserAND, i)
}

func (s *FormulasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormulasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormulasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.EnterFormulas(s)
	}
}

func (s *FormulasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.ExitFormulas(s)
	}
}

func (p *knowledgebaseParser) Formulas() (localctx IFormulasContext) {
	localctx = NewFormulasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, knowledgebaseParserRULE_formulas)
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
		p.SetState(40)
		p.Formula()
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == knowledgebaseParserAND {
		{
			p.SetState(41)
			p.Match(knowledgebaseParserAND)
		}
		{
			p.SetState(42)
			p.Formula()
		}

		p.SetState(47)
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
	p.RuleIndex = knowledgebaseParserRULE_formula
	return p
}

func (*FormulaContext) IsFormulaContext() {}

func NewFormulaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormulaContext {
	var p = new(FormulaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = knowledgebaseParserRULE_formula

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
	return s.GetToken(knowledgebaseParserBS, 0)
}

func (s *FormulaContext) BE() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserBE, 0)
}

func (s *FormulaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormulaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormulaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.EnterFormula(s)
	}
}

func (s *FormulaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.ExitFormula(s)
	}
}

func (p *knowledgebaseParser) Formula() (localctx IFormulaContext) {
	localctx = NewFormulaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, knowledgebaseParserRULE_formula)
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
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == knowledgebaseParserBS {
		{
			p.SetState(48)
			p.Match(knowledgebaseParserBS)
		}

	}
	{
		p.SetState(51)
		p.LeftOperand()
	}
	{
		p.SetState(52)
		p.Operator()
	}
	{
		p.SetState(53)
		p.RightOperand()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(54)
			p.Match(knowledgebaseParserBE)
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
	p.RuleIndex = knowledgebaseParserRULE_leftOperand
	return p
}

func (*LeftOperandContext) IsLeftOperandContext() {}

func NewLeftOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftOperandContext {
	var p = new(LeftOperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = knowledgebaseParserRULE_leftOperand

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
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.EnterLeftOperand(s)
	}
}

func (s *LeftOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.ExitLeftOperand(s)
	}
}

func (p *knowledgebaseParser) LeftOperand() (localctx ILeftOperandContext) {
	localctx = NewLeftOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, knowledgebaseParserRULE_leftOperand)

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
		p.SetState(57)
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
	p.RuleIndex = knowledgebaseParserRULE_rightOperand
	return p
}

func (*RightOperandContext) IsRightOperandContext() {}

func NewRightOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RightOperandContext {
	var p = new(RightOperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = knowledgebaseParserRULE_rightOperand

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
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.EnterRightOperand(s)
	}
}

func (s *RightOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.ExitRightOperand(s)
	}
}

func (p *knowledgebaseParser) RightOperand() (localctx IRightOperandContext) {
	localctx = NewRightOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, knowledgebaseParserRULE_rightOperand)

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
		p.SetState(59)
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
	p.RuleIndex = knowledgebaseParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = knowledgebaseParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserNUMBER, 0)
}

func (s *OperandContext) REAL() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserREAL, 0)
}

func (s *OperandContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserIDENTIFIER, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (p *knowledgebaseParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, knowledgebaseParserRULE_operand)
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
		p.SetState(61)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<knowledgebaseParserIDENTIFIER)|(1<<knowledgebaseParserNUMBER)|(1<<knowledgebaseParserREAL))) != 0) {
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
	p.RuleIndex = knowledgebaseParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = knowledgebaseParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) DEQ() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserDEQ, 0)
}

func (s *OperatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserNEQ, 0)
}

func (s *OperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserLT, 0)
}

func (s *OperatorContext) LE() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserLE, 0)
}

func (s *OperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserGT, 0)
}

func (s *OperatorContext) GE() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserGE, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *knowledgebaseParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, knowledgebaseParserRULE_operator)
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
		p.SetState(63)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<knowledgebaseParserDEQ)|(1<<knowledgebaseParserNEQ)|(1<<knowledgebaseParserLT)|(1<<knowledgebaseParserLE)|(1<<knowledgebaseParserGT)|(1<<knowledgebaseParserGE))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
