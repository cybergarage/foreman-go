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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 60, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 5, 2, 18, 10, 2, 3, 2, 3, 2, 5, 2, 22, 10, 2, 3, 3, 3, 3,
	3, 3, 7, 3, 27, 10, 3, 12, 3, 14, 3, 30, 11, 3, 3, 4, 5, 4, 33, 10, 4,
	3, 4, 3, 4, 5, 4, 37, 10, 4, 3, 5, 3, 5, 3, 5, 7, 5, 42, 10, 5, 12, 5,
	14, 5, 45, 11, 5, 3, 6, 5, 6, 48, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6,
	54, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10, 12, 14,
	2, 4, 4, 2, 17, 17, 19, 20, 3, 2, 6, 11, 2, 60, 2, 17, 3, 2, 2, 2, 4, 23,
	3, 2, 2, 2, 6, 32, 3, 2, 2, 2, 8, 38, 3, 2, 2, 2, 10, 47, 3, 2, 2, 2, 12,
	55, 3, 2, 2, 2, 14, 57, 3, 2, 2, 2, 16, 18, 7, 4, 2, 2, 17, 16, 3, 2, 2,
	2, 17, 18, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 21, 5, 4, 3, 2, 20, 22,
	7, 5, 2, 2, 21, 20, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 3, 3, 2, 2, 2,
	23, 28, 5, 6, 4, 2, 24, 25, 7, 12, 2, 2, 25, 27, 5, 6, 4, 2, 26, 24, 3,
	2, 2, 2, 27, 30, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29,
	5, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 31, 33, 7, 4, 2, 2, 32, 31, 3, 2, 2,
	2, 32, 33, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 36, 5, 8, 5, 2, 35, 37,
	7, 5, 2, 2, 36, 35, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 7, 3, 2, 2, 2,
	38, 43, 5, 10, 6, 2, 39, 40, 7, 13, 2, 2, 40, 42, 5, 10, 6, 2, 41, 39,
	3, 2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2,
	44, 9, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 48, 7, 4, 2, 2, 47, 46, 3, 2,
	2, 2, 47, 48, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 50, 5, 12, 7, 2, 50,
	51, 5, 14, 8, 2, 51, 53, 5, 12, 7, 2, 52, 54, 7, 5, 2, 2, 53, 52, 3, 2,
	2, 2, 53, 54, 3, 2, 2, 2, 54, 11, 3, 2, 2, 2, 55, 56, 9, 2, 2, 2, 56, 13,
	3, 2, 2, 2, 57, 58, 9, 3, 2, 2, 58, 15, 3, 2, 2, 2, 10, 17, 21, 28, 32,
	36, 43, 47, 53,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'*'", "'('", "')'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='",
	"'&'", "'|'", "','", "';'",
}
var symbolicNames = []string{
	"", "ASTERISK", "SB", "CB", "DEQ", "NEQ", "LT", "LE", "GT", "GE", "AND",
	"OR", "COMMA", "SEMICOLON", "WS", "IDENTIFIER", "STRING", "NUMBER", "REAL",
}

var ruleNames = []string{
	"knowledgebase", "clauses", "clause", "formulas", "formula", "operand",
	"operator",
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
	knowledgebaseParserSB         = 2
	knowledgebaseParserCB         = 3
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
	knowledgebaseParserRULE_knowledgebase = 0
	knowledgebaseParserRULE_clauses       = 1
	knowledgebaseParserRULE_clause        = 2
	knowledgebaseParserRULE_formulas      = 3
	knowledgebaseParserRULE_formula       = 4
	knowledgebaseParserRULE_operand       = 5
	knowledgebaseParserRULE_operator      = 6
)

// IKnowledgebaseContext is an interface to support dynamic dispatch.
type IKnowledgebaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKnowledgebaseContext differentiates from other interfaces.
	IsKnowledgebaseContext()
}

type KnowledgebaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKnowledgebaseContext() *KnowledgebaseContext {
	var p = new(KnowledgebaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = knowledgebaseParserRULE_knowledgebase
	return p
}

func (*KnowledgebaseContext) IsKnowledgebaseContext() {}

func NewKnowledgebaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KnowledgebaseContext {
	var p = new(KnowledgebaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = knowledgebaseParserRULE_knowledgebase

	return p
}

func (s *KnowledgebaseContext) GetParser() antlr.Parser { return s.parser }

func (s *KnowledgebaseContext) Clauses() IClausesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClausesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClausesContext)
}

func (s *KnowledgebaseContext) SB() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserSB, 0)
}

func (s *KnowledgebaseContext) CB() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserCB, 0)
}

func (s *KnowledgebaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KnowledgebaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KnowledgebaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.EnterKnowledgebase(s)
	}
}

func (s *KnowledgebaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(knowledgebaseListener); ok {
		listenerT.ExitKnowledgebase(s)
	}
}

func (p *knowledgebaseParser) Knowledgebase() (localctx IKnowledgebaseContext) {
	localctx = NewKnowledgebaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, knowledgebaseParserRULE_knowledgebase)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(14)
			p.Match(knowledgebaseParserSB)
		}

	}
	{
		p.SetState(17)
		p.Clauses()
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == knowledgebaseParserCB {
		{
			p.SetState(18)
			p.Match(knowledgebaseParserCB)
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

func (s *ClausesContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(knowledgebaseParserAND)
}

func (s *ClausesContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserAND, i)
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
		p.SetState(21)
		p.Clause()
	}
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == knowledgebaseParserAND {
		{
			p.SetState(22)
			p.Match(knowledgebaseParserAND)
		}
		{
			p.SetState(23)
			p.Clause()
		}

		p.SetState(28)
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

func (s *ClauseContext) SB() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserSB, 0)
}

func (s *ClauseContext) CB() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserCB, 0)
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
	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(29)
			p.Match(knowledgebaseParserSB)
		}

	}
	{
		p.SetState(32)
		p.Formulas()
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(33)
			p.Match(knowledgebaseParserCB)
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

func (s *FormulasContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(knowledgebaseParserOR)
}

func (s *FormulasContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserOR, i)
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
		p.SetState(36)
		p.Formula()
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == knowledgebaseParserOR {
		{
			p.SetState(37)
			p.Match(knowledgebaseParserOR)
		}
		{
			p.SetState(38)
			p.Formula()
		}

		p.SetState(43)
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

func (s *FormulaContext) AllOperand() []IOperandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperandContext)(nil)).Elem())
	var tst = make([]IOperandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperandContext)
		}
	}

	return tst
}

func (s *FormulaContext) Operand(i int) IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *FormulaContext) Operator() IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *FormulaContext) SB() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserSB, 0)
}

func (s *FormulaContext) CB() antlr.TerminalNode {
	return s.GetToken(knowledgebaseParserCB, 0)
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == knowledgebaseParserSB {
		{
			p.SetState(44)
			p.Match(knowledgebaseParserSB)
		}

	}
	{
		p.SetState(47)
		p.Operand()
	}
	{
		p.SetState(48)
		p.Operator()
	}
	{
		p.SetState(49)
		p.Operand()
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(50)
			p.Match(knowledgebaseParserCB)
		}

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
	p.EnterRule(localctx, 10, knowledgebaseParserRULE_operand)
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
		p.SetState(53)
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
	p.EnterRule(localctx, 12, knowledgebaseParserRULE_operator)
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
		p.SetState(55)
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
