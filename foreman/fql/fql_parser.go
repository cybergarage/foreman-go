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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 29, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 7, 3, 18, 10, 3, 12, 3, 14, 3, 21, 11, 3, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 2, 2, 24, 2, 12,
	3, 2, 2, 2, 4, 14, 3, 2, 2, 2, 6, 22, 3, 2, 2, 2, 8, 24, 3, 2, 2, 2, 10,
	26, 3, 2, 2, 2, 12, 13, 5, 4, 3, 2, 13, 3, 3, 2, 2, 2, 14, 19, 5, 6, 4,
	2, 15, 16, 7, 20, 2, 2, 16, 18, 5, 6, 4, 2, 17, 15, 3, 2, 2, 2, 18, 21,
	3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 5, 3, 2, 2, 2,
	21, 19, 3, 2, 2, 2, 22, 23, 5, 8, 5, 2, 23, 7, 3, 2, 2, 2, 24, 25, 7, 6,
	2, 2, 25, 9, 3, 2, 2, 2, 26, 27, 7, 8, 2, 2, 27, 11, 3, 2, 2, 2, 3, 19,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "'*'", "'='", "'=='", "'<'", "'<='", "'>'",
	"'>='", "'!='", "", "", "','", "';'",
}
var symbolicNames = []string{
	"", "SELECT", "INSERT", "DELETE", "SET", "EXPORT", "QOS", "ASTERISK", "SINGLE_EQ",
	"DOUBLE_EQ", "OP_LT", "LE", "GT", "GE", "NOTEQ", "AND", "OR", "COMMA",
	"SEMICOLON", "WS", "ID", "NUMBER", "FLOAT", "STRING", "CHAR",
}

var ruleNames = []string{
	"fql", "statement_list", "statement", "set_stmt", "qos_stmt",
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
	FQLParserQOS       = 6
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
	FQLParserRULE_qos_stmt       = 4
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
		p.SetState(10)
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
		p.SetState(12)
		p.Statement()
	}
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FQLParserSEMICOLON {
		{
			p.SetState(13)
			p.Match(FQLParserSEMICOLON)
		}
		{
			p.SetState(14)
			p.Statement()
		}

		p.SetState(19)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Set_stmt()
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

func (s *Set_stmtContext) SET() antlr.TerminalNode {
	return s.GetToken(FQLParserSET, 0)
}

func (s *Set_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterSet_stmt(s)
	}
}

func (s *Set_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitSet_stmt(s)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Match(FQLParserSET)
	}

	return localctx
}

// IQos_stmtContext is an interface to support dynamic dispatch.
type IQos_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQos_stmtContext differentiates from other interfaces.
	IsQos_stmtContext()
}

type Qos_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQos_stmtContext() *Qos_stmtContext {
	var p = new(Qos_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_qos_stmt
	return p
}

func (*Qos_stmtContext) IsQos_stmtContext() {}

func NewQos_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Qos_stmtContext {
	var p = new(Qos_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_qos_stmt

	return p
}

func (s *Qos_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Qos_stmtContext) QOS() antlr.TerminalNode {
	return s.GetToken(FQLParserQOS, 0)
}

func (s *Qos_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Qos_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Qos_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterQos_stmt(s)
	}
}

func (s *Qos_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitQos_stmt(s)
	}
}

func (p *FQLParser) Qos_stmt() (localctx IQos_stmtContext) {
	localctx = NewQos_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FQLParserRULE_qos_stmt)

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
		p.SetState(24)
		p.Match(FQLParserQOS)
	}

	return localctx
}
