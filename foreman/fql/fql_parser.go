// Code generated from FQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 30, 116,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 7, 3, 40, 10, 3, 12, 3, 14, 3, 43, 11, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 50, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 5, 7, 73, 10, 7, 3, 8, 3, 8, 3, 8, 7, 8, 78, 10, 8,
	12, 8, 14, 8, 81, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 97, 10, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	17, 7, 17, 111, 10, 17, 12, 17, 14, 17, 114, 11, 17, 3, 17, 2, 2, 18, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 2, 4, 3, 2, 15,
	20, 3, 2, 26, 29, 2, 108, 2, 34, 3, 2, 2, 2, 4, 36, 3, 2, 2, 2, 6, 49,
	3, 2, 2, 2, 8, 51, 3, 2, 2, 2, 10, 59, 3, 2, 2, 2, 12, 66, 3, 2, 2, 2,
	14, 74, 3, 2, 2, 2, 16, 82, 3, 2, 2, 2, 18, 86, 3, 2, 2, 2, 20, 88, 3,
	2, 2, 2, 22, 90, 3, 2, 2, 2, 24, 92, 3, 2, 2, 2, 26, 98, 3, 2, 2, 2, 28,
	103, 3, 2, 2, 2, 30, 105, 3, 2, 2, 2, 32, 107, 3, 2, 2, 2, 34, 35, 5, 4,
	3, 2, 35, 3, 3, 2, 2, 2, 36, 41, 5, 6, 4, 2, 37, 38, 7, 24, 2, 2, 38, 40,
	5, 6, 4, 2, 39, 37, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2,
	41, 42, 3, 2, 2, 2, 42, 5, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 50, 5, 8,
	5, 2, 45, 50, 5, 10, 6, 2, 46, 50, 5, 12, 7, 2, 47, 50, 5, 24, 13, 2, 48,
	50, 5, 26, 14, 2, 49, 44, 3, 2, 2, 2, 49, 45, 3, 2, 2, 2, 49, 46, 3, 2,
	2, 2, 49, 47, 3, 2, 2, 2, 49, 48, 3, 2, 2, 2, 50, 7, 3, 2, 2, 2, 51, 52,
	7, 6, 2, 2, 52, 53, 7, 10, 2, 2, 53, 54, 5, 28, 15, 2, 54, 55, 7, 12, 2,
	2, 55, 56, 7, 3, 2, 2, 56, 57, 5, 32, 17, 2, 57, 58, 7, 4, 2, 2, 58, 9,
	3, 2, 2, 2, 59, 60, 7, 8, 2, 2, 60, 61, 7, 3, 2, 2, 61, 62, 5, 32, 17,
	2, 62, 63, 7, 4, 2, 2, 63, 64, 7, 10, 2, 2, 64, 65, 5, 28, 15, 2, 65, 11,
	3, 2, 2, 2, 66, 67, 7, 5, 2, 2, 67, 68, 7, 13, 2, 2, 68, 69, 7, 11, 2,
	2, 69, 72, 5, 28, 15, 2, 70, 71, 7, 30, 2, 2, 71, 73, 5, 14, 8, 2, 72,
	70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 13, 3, 2, 2, 2, 74, 79, 5, 16,
	9, 2, 75, 76, 7, 21, 2, 2, 76, 78, 5, 16, 9, 2, 77, 75, 3, 2, 2, 2, 78,
	81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 15, 3, 2, 2,
	2, 81, 79, 3, 2, 2, 2, 82, 83, 5, 20, 11, 2, 83, 84, 5, 18, 10, 2, 84,
	85, 5, 22, 12, 2, 85, 17, 3, 2, 2, 2, 86, 87, 9, 2, 2, 2, 87, 19, 3, 2,
	2, 2, 88, 89, 7, 26, 2, 2, 89, 21, 3, 2, 2, 2, 90, 91, 7, 26, 2, 2, 91,
	23, 3, 2, 2, 2, 92, 93, 7, 9, 2, 2, 93, 96, 5, 28, 15, 2, 94, 95, 7, 30,
	2, 2, 95, 97, 5, 14, 8, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97,
	25, 3, 2, 2, 2, 98, 99, 7, 7, 2, 2, 99, 100, 5, 30, 16, 2, 100, 101, 7,
	11, 2, 2, 101, 102, 5, 28, 15, 2, 102, 27, 3, 2, 2, 2, 103, 104, 7, 26,
	2, 2, 104, 29, 3, 2, 2, 2, 105, 106, 9, 3, 2, 2, 106, 31, 3, 2, 2, 2, 107,
	112, 5, 30, 16, 2, 108, 109, 7, 23, 2, 2, 109, 111, 5, 30, 16, 2, 110,
	108, 3, 2, 2, 2, 111, 114, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113,
	3, 2, 2, 2, 113, 33, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 8, 41, 49, 72, 79,
	96, 112,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "", "", "", "", "", "", "", "", "'*'", "'='", "'=='",
	"'<'", "'<='", "'>'", "'>='", "'!='", "", "", "','", "';'",
}
var symbolicNames = []string{
	"", "", "", "SELECT", "INSERT", "DELETE", "SET", "EXPORT", "INTO", "FROM",
	"VALUES", "ASTERISK", "SINGLE_EQ", "DOUBLE_EQ", "OP_LT", "LE", "GT", "GE",
	"NOTEQ", "AND", "OR", "COMMA", "SEMICOLON", "WS", "IDENTIFIER", "STRING",
	"NUMBER", "REAL", "WHERE",
}

var ruleNames = []string{
	"fql", "statement_list", "statement", "insert_stmt", "set_stmt", "select_stmt",
	"conditions", "condition", "operator", "leftOperand", "rightOperand", "export_stmt",
	"del_stmt", "target", "value", "values",
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
	FQLParserT__0       = 1
	FQLParserT__1       = 2
	FQLParserSELECT     = 3
	FQLParserINSERT     = 4
	FQLParserDELETE     = 5
	FQLParserSET        = 6
	FQLParserEXPORT     = 7
	FQLParserINTO       = 8
	FQLParserFROM       = 9
	FQLParserVALUES     = 10
	FQLParserASTERISK   = 11
	FQLParserSINGLE_EQ  = 12
	FQLParserDOUBLE_EQ  = 13
	FQLParserOP_LT      = 14
	FQLParserLE         = 15
	FQLParserGT         = 16
	FQLParserGE         = 17
	FQLParserNOTEQ      = 18
	FQLParserAND        = 19
	FQLParserOR         = 20
	FQLParserCOMMA      = 21
	FQLParserSEMICOLON  = 22
	FQLParserWS         = 23
	FQLParserIDENTIFIER = 24
	FQLParserSTRING     = 25
	FQLParserNUMBER     = 26
	FQLParserREAL       = 27
	FQLParserWHERE      = 28
)

// FQLParser rules.
const (
	FQLParserRULE_fql            = 0
	FQLParserRULE_statement_list = 1
	FQLParserRULE_statement      = 2
	FQLParserRULE_insert_stmt    = 3
	FQLParserRULE_set_stmt       = 4
	FQLParserRULE_select_stmt    = 5
	FQLParserRULE_conditions     = 6
	FQLParserRULE_condition      = 7
	FQLParserRULE_operator       = 8
	FQLParserRULE_leftOperand    = 9
	FQLParserRULE_rightOperand   = 10
	FQLParserRULE_export_stmt    = 11
	FQLParserRULE_del_stmt       = 12
	FQLParserRULE_target         = 13
	FQLParserRULE_value          = 14
	FQLParserRULE_values         = 15
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
		p.SetState(32)
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
		p.SetState(34)
		p.Statement()
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FQLParserSEMICOLON {
		{
			p.SetState(35)
			p.Match(FQLParserSEMICOLON)
		}
		{
			p.SetState(36)
			p.Statement()
		}

		p.SetState(41)
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

func (s *StatementContext) Insert_stmt() IInsert_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsert_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsert_stmtContext)
}

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

func (s *StatementContext) Export_stmt() IExport_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExport_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExport_stmtContext)
}

func (s *StatementContext) Del_stmt() IDel_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDel_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDel_stmtContext)
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

	p.SetState(47)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FQLParserINSERT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Insert_stmt()
		}

	case FQLParserSET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Set_stmt()
		}

	case FQLParserSELECT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(44)
			p.Select_stmt()
		}

	case FQLParserEXPORT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(45)
			p.Export_stmt()
		}

	case FQLParserDELETE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(46)
			p.Del_stmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInsert_stmtContext is an interface to support dynamic dispatch.
type IInsert_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsert_stmtContext differentiates from other interfaces.
	IsInsert_stmtContext()
}

type Insert_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsert_stmtContext() *Insert_stmtContext {
	var p = new(Insert_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_insert_stmt
	return p
}

func (*Insert_stmtContext) IsInsert_stmtContext() {}

func NewInsert_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Insert_stmtContext {
	var p = new(Insert_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_insert_stmt

	return p
}

func (s *Insert_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Insert_stmtContext) CopyFrom(ctx *Insert_stmtContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Insert_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Insert_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InsertQueryContext struct {
	*Insert_stmtContext
}

func NewInsertQueryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InsertQueryContext {
	var p = new(InsertQueryContext)

	p.Insert_stmtContext = NewEmptyInsert_stmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Insert_stmtContext))

	return p
}

func (s *InsertQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertQueryContext) INSERT() antlr.TerminalNode {
	return s.GetToken(FQLParserINSERT, 0)
}

func (s *InsertQueryContext) INTO() antlr.TerminalNode {
	return s.GetToken(FQLParserINTO, 0)
}

func (s *InsertQueryContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *InsertQueryContext) VALUES() antlr.TerminalNode {
	return s.GetToken(FQLParserVALUES, 0)
}

func (s *InsertQueryContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *InsertQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterInsertQuery(s)
	}
}

func (s *InsertQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitInsertQuery(s)
	}
}

func (p *FQLParser) Insert_stmt() (localctx IInsert_stmtContext) {
	localctx = NewInsert_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FQLParserRULE_insert_stmt)

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

	localctx = NewInsertQueryContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(FQLParserINSERT)
	}
	{
		p.SetState(50)
		p.Match(FQLParserINTO)
	}
	{
		p.SetState(51)
		p.Target()
	}
	{
		p.SetState(52)
		p.Match(FQLParserVALUES)
	}
	{
		p.SetState(53)
		p.Match(FQLParserT__0)
	}
	{
		p.SetState(54)
		p.Values()
	}
	{
		p.SetState(55)
		p.Match(FQLParserT__1)
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

type SetQueryContext struct {
	*Set_stmtContext
}

func NewSetQueryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetQueryContext {
	var p = new(SetQueryContext)

	p.Set_stmtContext = NewEmptySet_stmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Set_stmtContext))

	return p
}

func (s *SetQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetQueryContext) SET() antlr.TerminalNode {
	return s.GetToken(FQLParserSET, 0)
}

func (s *SetQueryContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *SetQueryContext) INTO() antlr.TerminalNode {
	return s.GetToken(FQLParserINTO, 0)
}

func (s *SetQueryContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *SetQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterSetQuery(s)
	}
}

func (s *SetQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitSetQuery(s)
	}
}

func (p *FQLParser) Set_stmt() (localctx ISet_stmtContext) {
	localctx = NewSet_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FQLParserRULE_set_stmt)

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

	localctx = NewSetQueryContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(FQLParserSET)
	}
	{
		p.SetState(58)
		p.Match(FQLParserT__0)
	}
	{
		p.SetState(59)
		p.Values()
	}
	{
		p.SetState(60)
		p.Match(FQLParserT__1)
	}
	{
		p.SetState(61)
		p.Match(FQLParserINTO)
	}
	{
		p.SetState(62)
		p.Target()
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

type SelectQueryContext struct {
	*Select_stmtContext
}

func NewSelectQueryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectQueryContext {
	var p = new(SelectQueryContext)

	p.Select_stmtContext = NewEmptySelect_stmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Select_stmtContext))

	return p
}

func (s *SelectQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectQueryContext) SELECT() antlr.TerminalNode {
	return s.GetToken(FQLParserSELECT, 0)
}

func (s *SelectQueryContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(FQLParserASTERISK, 0)
}

func (s *SelectQueryContext) FROM() antlr.TerminalNode {
	return s.GetToken(FQLParserFROM, 0)
}

func (s *SelectQueryContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *SelectQueryContext) WHERE() antlr.TerminalNode {
	return s.GetToken(FQLParserWHERE, 0)
}

func (s *SelectQueryContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *SelectQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterSelectQuery(s)
	}
}

func (s *SelectQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitSelectQuery(s)
	}
}

func (p *FQLParser) Select_stmt() (localctx ISelect_stmtContext) {
	localctx = NewSelect_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FQLParserRULE_select_stmt)
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

	localctx = NewSelectQueryContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(FQLParserSELECT)
	}
	{
		p.SetState(65)
		p.Match(FQLParserASTERISK)
	}
	{
		p.SetState(66)
		p.Match(FQLParserFROM)
	}
	{
		p.SetState(67)
		p.Target()
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FQLParserWHERE {
		{
			p.SetState(68)
			p.Match(FQLParserWHERE)
		}
		{
			p.SetState(69)
			p.Conditions()
		}

	}

	return localctx
}

// IConditionsContext is an interface to support dynamic dispatch.
type IConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionsContext differentiates from other interfaces.
	IsConditionsContext()
}

type ConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsContext() *ConditionsContext {
	var p = new(ConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_conditions
	return p
}

func (*ConditionsContext) IsConditionsContext() {}

func NewConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsContext {
	var p = new(ConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_conditions

	return p
}

func (s *ConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsContext) AllCondition() []IConditionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionContext)(nil)).Elem())
	var tst = make([]IConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionContext)
		}
	}

	return tst
}

func (s *ConditionsContext) Condition(i int) IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionsContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(FQLParserAND)
}

func (s *ConditionsContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(FQLParserAND, i)
}

func (s *ConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterConditions(s)
	}
}

func (s *ConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitConditions(s)
	}
}

func (p *FQLParser) Conditions() (localctx IConditionsContext) {
	localctx = NewConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FQLParserRULE_conditions)
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
		p.SetState(72)
		p.Condition()
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FQLParserAND {
		{
			p.SetState(73)
			p.Match(FQLParserAND)
		}
		{
			p.SetState(74)
			p.Condition()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) LeftOperand() ILeftOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILeftOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILeftOperandContext)
}

func (s *ConditionContext) Operator() IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *ConditionContext) RightOperand() IRightOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRightOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRightOperandContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *FQLParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FQLParserRULE_condition)

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
		p.LeftOperand()
	}
	{
		p.SetState(81)
		p.Operator()
	}
	{
		p.SetState(82)
		p.RightOperand()
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
	p.RuleIndex = FQLParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) DOUBLE_EQ() antlr.TerminalNode {
	return s.GetToken(FQLParserDOUBLE_EQ, 0)
}

func (s *OperatorContext) OP_LT() antlr.TerminalNode {
	return s.GetToken(FQLParserOP_LT, 0)
}

func (s *OperatorContext) LE() antlr.TerminalNode {
	return s.GetToken(FQLParserLE, 0)
}

func (s *OperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(FQLParserGT, 0)
}

func (s *OperatorContext) GE() antlr.TerminalNode {
	return s.GetToken(FQLParserGE, 0)
}

func (s *OperatorContext) NOTEQ() antlr.TerminalNode {
	return s.GetToken(FQLParserNOTEQ, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *FQLParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FQLParserRULE_operator)
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
		p.SetState(84)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FQLParserDOUBLE_EQ)|(1<<FQLParserOP_LT)|(1<<FQLParserLE)|(1<<FQLParserGT)|(1<<FQLParserGE)|(1<<FQLParserNOTEQ))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	p.RuleIndex = FQLParserRULE_leftOperand
	return p
}

func (*LeftOperandContext) IsLeftOperandContext() {}

func NewLeftOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftOperandContext {
	var p = new(LeftOperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_leftOperand

	return p
}

func (s *LeftOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *LeftOperandContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FQLParserIDENTIFIER, 0)
}

func (s *LeftOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeftOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeftOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterLeftOperand(s)
	}
}

func (s *LeftOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitLeftOperand(s)
	}
}

func (p *FQLParser) LeftOperand() (localctx ILeftOperandContext) {
	localctx = NewLeftOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FQLParserRULE_leftOperand)

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
		p.SetState(86)
		p.Match(FQLParserIDENTIFIER)
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
	p.RuleIndex = FQLParserRULE_rightOperand
	return p
}

func (*RightOperandContext) IsRightOperandContext() {}

func NewRightOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RightOperandContext {
	var p = new(RightOperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_rightOperand

	return p
}

func (s *RightOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *RightOperandContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FQLParserIDENTIFIER, 0)
}

func (s *RightOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RightOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RightOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterRightOperand(s)
	}
}

func (s *RightOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitRightOperand(s)
	}
}

func (p *FQLParser) RightOperand() (localctx IRightOperandContext) {
	localctx = NewRightOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FQLParserRULE_rightOperand)

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
		p.SetState(88)
		p.Match(FQLParserIDENTIFIER)
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

type ExportQueryContext struct {
	*Export_stmtContext
}

func NewExportQueryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExportQueryContext {
	var p = new(ExportQueryContext)

	p.Export_stmtContext = NewEmptyExport_stmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Export_stmtContext))

	return p
}

func (s *ExportQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportQueryContext) EXPORT() antlr.TerminalNode {
	return s.GetToken(FQLParserEXPORT, 0)
}

func (s *ExportQueryContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *ExportQueryContext) WHERE() antlr.TerminalNode {
	return s.GetToken(FQLParserWHERE, 0)
}

func (s *ExportQueryContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *ExportQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterExportQuery(s)
	}
}

func (s *ExportQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitExportQuery(s)
	}
}

func (p *FQLParser) Export_stmt() (localctx IExport_stmtContext) {
	localctx = NewExport_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FQLParserRULE_export_stmt)
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

	localctx = NewExportQueryContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(FQLParserEXPORT)
	}
	{
		p.SetState(91)
		p.Target()
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FQLParserWHERE {
		{
			p.SetState(92)
			p.Match(FQLParserWHERE)
		}
		{
			p.SetState(93)
			p.Conditions()
		}

	}

	return localctx
}

// IDel_stmtContext is an interface to support dynamic dispatch.
type IDel_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDel_stmtContext differentiates from other interfaces.
	IsDel_stmtContext()
}

type Del_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDel_stmtContext() *Del_stmtContext {
	var p = new(Del_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_del_stmt
	return p
}

func (*Del_stmtContext) IsDel_stmtContext() {}

func NewDel_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Del_stmtContext {
	var p = new(Del_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_del_stmt

	return p
}

func (s *Del_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Del_stmtContext) CopyFrom(ctx *Del_stmtContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Del_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Del_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeleteQueryContext struct {
	*Del_stmtContext
}

func NewDeleteQueryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeleteQueryContext {
	var p = new(DeleteQueryContext)

	p.Del_stmtContext = NewEmptyDel_stmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Del_stmtContext))

	return p
}

func (s *DeleteQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteQueryContext) DELETE() antlr.TerminalNode {
	return s.GetToken(FQLParserDELETE, 0)
}

func (s *DeleteQueryContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DeleteQueryContext) FROM() antlr.TerminalNode {
	return s.GetToken(FQLParserFROM, 0)
}

func (s *DeleteQueryContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *DeleteQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterDeleteQuery(s)
	}
}

func (s *DeleteQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitDeleteQuery(s)
	}
}

func (p *FQLParser) Del_stmt() (localctx IDel_stmtContext) {
	localctx = NewDel_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FQLParserRULE_del_stmt)

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

	localctx = NewDeleteQueryContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(FQLParserDELETE)
	}
	{
		p.SetState(97)
		p.Value()
	}
	{
		p.SetState(98)
		p.Match(FQLParserFROM)
	}
	{
		p.SetState(99)
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
	p.EnterRule(localctx, 26, FQLParserRULE_target)

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
		p.SetState(101)
		p.Match(FQLParserIDENTIFIER)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FQLParserIDENTIFIER, 0)
}

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(FQLParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FQLParserNUMBER, 0)
}

func (s *ValueContext) REAL() antlr.TerminalNode {
	return s.GetToken(FQLParserREAL, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *FQLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FQLParserRULE_value)
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
		p.SetState(103)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FQLParserIDENTIFIER)|(1<<FQLParserSTRING)|(1<<FQLParserNUMBER)|(1<<FQLParserREAL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IValuesContext is an interface to support dynamic dispatch.
type IValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValuesContext differentiates from other interfaces.
	IsValuesContext()
}

type ValuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuesContext() *ValuesContext {
	var p = new(ValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_values
	return p
}

func (*ValuesContext) IsValuesContext() {}

func NewValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuesContext {
	var p = new(ValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_values

	return p
}

func (s *ValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuesContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ValuesContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterValues(s)
	}
}

func (s *ValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitValues(s)
	}
}

func (p *FQLParser) Values() (localctx IValuesContext) {
	localctx = NewValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FQLParserRULE_values)
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
		p.SetState(105)
		p.Value()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FQLParserCOMMA {
		{
			p.SetState(106)
			p.Match(FQLParserCOMMA)
		}
		{
			p.SetState(107)
			p.Value()
		}

		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
