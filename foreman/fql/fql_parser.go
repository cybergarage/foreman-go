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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 38, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 3, 20, 10, 3, 12, 3, 14, 3, 23, 11, 3, 3,
	4, 3, 4, 5, 4, 27, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10, 12, 2, 2, 2, 33, 2, 14, 3, 2, 2,
	2, 4, 16, 3, 2, 2, 2, 6, 26, 3, 2, 2, 2, 8, 28, 3, 2, 2, 2, 10, 30, 3,
	2, 2, 2, 12, 35, 3, 2, 2, 2, 14, 15, 5, 4, 3, 2, 15, 3, 3, 2, 2, 2, 16,
	21, 5, 6, 4, 2, 17, 18, 7, 20, 2, 2, 18, 20, 5, 6, 4, 2, 19, 17, 3, 2,
	2, 2, 20, 23, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 5,
	3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 24, 27, 5, 8, 5, 2, 25, 27, 5, 10, 6, 2,
	26, 24, 3, 2, 2, 2, 26, 25, 3, 2, 2, 2, 27, 7, 3, 2, 2, 2, 28, 29, 7, 6,
	2, 2, 29, 9, 3, 2, 2, 2, 30, 31, 7, 3, 2, 2, 31, 32, 7, 9, 2, 2, 32, 33,
	7, 8, 2, 2, 33, 34, 5, 12, 7, 2, 34, 11, 3, 2, 2, 2, 35, 36, 7, 25, 2,
	2, 36, 13, 3, 2, 2, 2, 4, 21, 26,
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
	"COMMA", "SEMICOLON", "WS", "ID", "NUMBER", "FLOAT", "STRING", "CHAR",
}

var ruleNames = []string{
	"fql", "statement_list", "statement", "set_stmt", "select_stmt", "table",
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
	FQLParserEOF       = antlr.TokenEOF
	FQLParserSELECT    = 1
	FQLParserINSERT    = 2
	FQLParserDELETE    = 3
	FQLParserSET       = 4
	FQLParserEXPORT    = 5
	FQLParserFROM      = 6
	FQLParserASTERISK  = 7
	FQLParserSINGLE_EQ = 8
	FQLParserDOUBLE_EQ = 9
	FQLParserOP_LT     = 10
	FQLParserLE        = 11
	FQLParserGT        = 12
	FQLParserGE        = 13
	FQLParserNOTEQ     = 14
	FQLParserAND       = 15
	FQLParserOR        = 16
	FQLParserCOMMA     = 17
	FQLParserSEMICOLON = 18
	FQLParserWS        = 19
	FQLParserID        = 20
	FQLParserNUMBER    = 21
	FQLParserFLOAT     = 22
	FQLParserSTRING    = 23
	FQLParserCHAR      = 24
)

// FQLParser rules.
const (
	FQLParserRULE_fql            = 0
	FQLParserRULE_statement_list = 1
	FQLParserRULE_statement      = 2
	FQLParserRULE_set_stmt       = 3
	FQLParserRULE_select_stmt    = 4
	FQLParserRULE_table          = 5
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

func (s *FqlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FQLVisitor:
		return t.VisitFql(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(12)
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

func (s *Statement_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FQLVisitor:
		return t.VisitStatement_list(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(14)
		p.Statement()
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FQLParserSEMICOLON {
		{
			p.SetState(15)
			p.Match(FQLParserSEMICOLON)
		}
		{
			p.SetState(16)
			p.Statement()
		}

		p.SetState(21)
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

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FQLVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
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

	p.SetState(24)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FQLParserSET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Set_stmt()
		}

	case FQLParserSELECT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Select_stmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *SetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FQLVisitor:
		return t.VisitSet(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(26)
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

func (s *SelectContext) Table() ITableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableContext)
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

func (s *SelectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FQLVisitor:
		return t.VisitSelect(s)

	default:
		return t.VisitChildren(s)
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
		p.SetState(28)
		p.Match(FQLParserSELECT)
	}
	{
		p.SetState(29)
		p.Match(FQLParserASTERISK)
	}
	{
		p.SetState(30)
		p.Match(FQLParserFROM)
	}
	{
		p.SetState(31)
		p.Table()
	}

	return localctx
}

// ITableContext is an interface to support dynamic dispatch.
type ITableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableContext differentiates from other interfaces.
	IsTableContext()
}

type TableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableContext() *TableContext {
	var p = new(TableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_table
	return p
}

func (*TableContext) IsTableContext() {}

func NewTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableContext {
	var p = new(TableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_table

	return p
}

func (s *TableContext) GetParser() antlr.Parser { return s.parser }

func (s *TableContext) STRING() antlr.TerminalNode {
	return s.GetToken(FQLParserSTRING, 0)
}

func (s *TableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterTable(s)
	}
}

func (s *TableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitTable(s)
	}
}

func (s *TableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FQLVisitor:
		return t.VisitTable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FQLParser) Table() (localctx ITableContext) {
	localctx = NewTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FQLParserRULE_table)

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
		p.SetState(33)
		p.Match(FQLParserSTRING)
	}

	return localctx
}
