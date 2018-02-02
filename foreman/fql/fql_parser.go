// Generated from FQL.g4 by ANTLR 4.7.

package fql // FQL
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 24, 44, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 3, 22, 10, 3, 12, 3, 14, 3, 25,
	11, 3, 3, 4, 3, 4, 3, 4, 5, 4, 30, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10,
	12, 14, 2, 2, 2, 39, 2, 16, 3, 2, 2, 2, 4, 18, 3, 2, 2, 2, 6, 29, 3, 2,
	2, 2, 8, 31, 3, 2, 2, 2, 10, 33, 3, 2, 2, 2, 12, 38, 3, 2, 2, 2, 14, 41,
	3, 2, 2, 2, 16, 17, 5, 4, 3, 2, 17, 3, 3, 2, 2, 2, 18, 23, 5, 6, 4, 2,
	19, 20, 7, 20, 2, 2, 20, 22, 5, 6, 4, 2, 21, 19, 3, 2, 2, 2, 22, 25, 3,
	2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 5, 3, 2, 2, 2, 25,
	23, 3, 2, 2, 2, 26, 30, 5, 8, 5, 2, 27, 30, 5, 10, 6, 2, 28, 30, 5, 10,
	6, 2, 29, 26, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 28, 3, 2, 2, 2, 30, 7,
	3, 2, 2, 2, 31, 32, 7, 6, 2, 2, 32, 9, 3, 2, 2, 2, 33, 34, 7, 3, 2, 2,
	34, 35, 7, 9, 2, 2, 35, 36, 7, 8, 2, 2, 36, 37, 5, 14, 8, 2, 37, 11, 3,
	2, 2, 2, 38, 39, 7, 7, 2, 2, 39, 40, 5, 14, 8, 2, 40, 13, 3, 2, 2, 2, 41,
	42, 7, 22, 2, 2, 42, 15, 3, 2, 2, 2, 4, 23, 29,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "'*'", "'='", "'=='", "'<'", "'<='", "'>'",
	"'>='", "'!='", "", "", "','", "';'",
}
var symbolicNames = []string{
	"", "SELECT", "INSERT", "DELETE", "SET", "EXPORT", "FROM", "ASTERISK",
	"SINGLE_EQ", "DOUBLE_EQ", "OP_LT", "LE", "GT", "GE", "NOTEQ", "AND", "OR",
	"COMMA", "SEMICOLON", "WS", "IDENTIFIER", "NUMBER", "REAL",
}

var ruleNames = []string{
	"fql", "statement_list", "statement", "set_stmt", "select_stmt", "export_stmt",
	"target",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FQLParser struct {
	*antlr.BaseParser
}

func NewFQLParser(input antlr.TokenStream) *FQLParser {
	this := new(FQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "FQL.g4"

	return this
}

// FQLParser tokens.
const (
	FQLParserEOF        = antlr.TokenEOF
	FQLParserSELECT     = 1
	FQLParserINSERT     = 2
	FQLParserDELETE     = 3
	FQLParserSET        = 4
	FQLParserEXPORT     = 5
	FQLParserFROM       = 6
	FQLParserASTERISK   = 7
	FQLParserSINGLE_EQ  = 8
	FQLParserDOUBLE_EQ  = 9
	FQLParserOP_LT      = 10
	FQLParserLE         = 11
	FQLParserGT         = 12
	FQLParserGE         = 13
	FQLParserNOTEQ      = 14
	FQLParserAND        = 15
	FQLParserOR         = 16
	FQLParserCOMMA      = 17
	FQLParserSEMICOLON  = 18
	FQLParserWS         = 19
	FQLParserIDENTIFIER = 20
	FQLParserNUMBER     = 21
	FQLParserREAL       = 22
)

// FQLParser rules.
const (
	FQLParserRULE_fql            = 0
	FQLParserRULE_statement_list = 1
	FQLParserRULE_statement      = 2
	FQLParserRULE_set_stmt       = 3
	FQLParserRULE_select_stmt    = 4
	FQLParserRULE_export_stmt    = 5
	FQLParserRULE_target         = 6
)

// IFqlContext is an interface to support dynamic dispatch.
type IFqlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFqlContext differentiates from other interfaces.
	IsFqlContext()
}

type FqlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFqlContext() *FqlContext {
	var p = new(FqlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_fql
	return p
}

func (*FqlContext) IsFqlContext() {}

func NewFqlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FqlContext {
	var p = new(FqlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_fql

	return p
}

func (s *FqlContext) GetParser() antlr.Parser { return s.parser }

func (s *FqlContext) Statement_list() IStatement_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatement_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *FqlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FqlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FqlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterFql(s)
	}
}

func (s *FqlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitFql(s)
	}
}

func (p *FQLParser) Fql() (localctx IFqlContext) {
	localctx = NewFqlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FQLParserRULE_fql)

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
		p.SetState(14)
		p.Statement_list()
	}

	return localctx
}

// IStatement_listContext is an interface to support dynamic dispatch.
type IStatement_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatement_listContext differentiates from other interfaces.
	IsStatement_listContext()
}

type Statement_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatement_listContext() *Statement_listContext {
	var p = new(Statement_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_statement_list
	return p
}

func (*Statement_listContext) IsStatement_listContext() {}

func NewStatement_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statement_listContext {
	var p = new(Statement_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_statement_list

	return p
}

func (s *Statement_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Statement_listContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *Statement_listContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Statement_listContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(FQLParserSEMICOLON)
}

func (s *Statement_listContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(FQLParserSEMICOLON, i)
}

func (s *Statement_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Statement_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterStatement_list(s)
	}
}

func (s *Statement_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitStatement_list(s)
	}
}

func (p *FQLParser) Statement_list() (localctx IStatement_listContext) {
	localctx = NewStatement_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FQLParserRULE_statement_list)
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
		p.SetState(16)
		p.Statement()
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FQLParserSEMICOLON {
		{
			p.SetState(17)
			p.Match(FQLParserSEMICOLON)
		}
		{
			p.SetState(18)
			p.Statement()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Set_stmt() ISet_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_stmtContext)
}

func (s *StatementContext) Select_stmt() ISelect_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelect_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelect_stmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *FQLParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FQLParserRULE_statement)

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

	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Set_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Select_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(26)
			p.Select_stmt()
		}

	}

	return localctx
}

// ISet_stmtContext is an interface to support dynamic dispatch.
type ISet_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_stmtContext differentiates from other interfaces.
	IsSet_stmtContext()
}

type Set_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_stmtContext() *Set_stmtContext {
	var p = new(Set_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_set_stmt
	return p
}

func (*Set_stmtContext) IsSet_stmtContext() {}

func NewSet_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_stmtContext {
	var p = new(Set_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_set_stmt

	return p
}

func (s *Set_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_stmtContext) CopyFrom(ctx *Set_stmtContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Set_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SetContext struct {
	*Set_stmtContext
}

func NewSetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetContext {
	var p = new(SetContext)

	p.Set_stmtContext = NewEmptySet_stmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Set_stmtContext))

	return p
}

func (s *SetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetContext) SET() antlr.TerminalNode {
	return s.GetToken(FQLParserSET, 0)
}

func (s *SetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterSet(s)
	}
}

func (s *SetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitSet(s)
	}
}

func (p *FQLParser) Set_stmt() (localctx ISet_stmtContext) {
	localctx = NewSet_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FQLParserRULE_set_stmt)

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

	localctx = NewSetContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(FQLParserSET)
	}

	return localctx
}

// ISelect_stmtContext is an interface to support dynamic dispatch.
type ISelect_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelect_stmtContext differentiates from other interfaces.
	IsSelect_stmtContext()
}

type Select_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_stmtContext() *Select_stmtContext {
	var p = new(Select_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_select_stmt
	return p
}

func (*Select_stmtContext) IsSelect_stmtContext() {}

func NewSelect_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_stmtContext {
	var p = new(Select_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_select_stmt

	return p
}

func (s *Select_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_stmtContext) CopyFrom(ctx *Select_stmtContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Select_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SelectContext struct {
	*Select_stmtContext
}

func NewSelectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectContext {
	var p = new(SelectContext)

	p.Select_stmtContext = NewEmptySelect_stmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Select_stmtContext))

	return p
}

func (s *SelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectContext) SELECT() antlr.TerminalNode {
	return s.GetToken(FQLParserSELECT, 0)
}

func (s *SelectContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(FQLParserASTERISK, 0)
}

func (s *SelectContext) FROM() antlr.TerminalNode {
	return s.GetToken(FQLParserFROM, 0)
}

func (s *SelectContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *SelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterSelect(s)
	}
}

func (s *SelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitSelect(s)
	}
}

func (p *FQLParser) Select_stmt() (localctx ISelect_stmtContext) {
	localctx = NewSelect_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FQLParserRULE_select_stmt)

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

	localctx = NewSelectContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(FQLParserSELECT)
	}
	{
		p.SetState(32)
		p.Match(FQLParserASTERISK)
	}
	{
		p.SetState(33)
		p.Match(FQLParserFROM)
	}
	{
		p.SetState(34)
		p.Target()
	}

	return localctx
}

// IExport_stmtContext is an interface to support dynamic dispatch.
type IExport_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExport_stmtContext differentiates from other interfaces.
	IsExport_stmtContext()
}

type Export_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExport_stmtContext() *Export_stmtContext {
	var p = new(Export_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_export_stmt
	return p
}

func (*Export_stmtContext) IsExport_stmtContext() {}

func NewExport_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Export_stmtContext {
	var p = new(Export_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_export_stmt

	return p
}

func (s *Export_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Export_stmtContext) CopyFrom(ctx *Export_stmtContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Export_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Export_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExportContext struct {
	*Export_stmtContext
}

func NewExportContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExportContext {
	var p = new(ExportContext)

	p.Export_stmtContext = NewEmptyExport_stmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Export_stmtContext))

	return p
}

func (s *ExportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportContext) EXPORT() antlr.TerminalNode {
	return s.GetToken(FQLParserEXPORT, 0)
}

func (s *ExportContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *ExportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterExport(s)
	}
}

func (s *ExportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitExport(s)
	}
}

func (p *FQLParser) Export_stmt() (localctx IExport_stmtContext) {
	localctx = NewExport_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FQLParserRULE_export_stmt)

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

	localctx = NewExportContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Match(FQLParserEXPORT)
	}
	{
		p.SetState(37)
		p.Target()
	}

	return localctx
}

// ITargetContext is an interface to support dynamic dispatch.
type ITargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetContext differentiates from other interfaces.
	IsTargetContext()
}

type TargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetContext() *TargetContext {
	var p = new(TargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_target
	return p
}

func (*TargetContext) IsTargetContext() {}

func NewTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetContext {
	var p = new(TargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_target

	return p
}

func (s *TargetContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FQLParserIDENTIFIER, 0)
}

func (s *TargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterTarget(s)
	}
}

func (s *TargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitTarget(s)
	}
}

func (p *FQLParser) Target() (localctx ITargetContext) {
	localctx = NewTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FQLParserRULE_target)

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
		p.SetState(39)
		p.Match(FQLParserIDENTIFIER)
	}

	return localctx
}
