// Code generated from FQL.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 208,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 57,
	10, 3, 12, 3, 14, 3, 60, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 5, 4, 70, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 5, 6, 81, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 101,
	10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 107, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 5, 9, 114, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 120, 10, 10,
	3, 10, 3, 10, 5, 10, 124, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5,
	11, 131, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 139,
	10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 147, 10, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 154, 10, 13, 3, 14, 3, 14, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 16, 7, 16, 163, 10, 16, 12, 16, 14, 16, 166, 11,
	16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 7, 18, 173, 10, 18, 12, 18, 14,
	18, 176, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 7, 20,
	185, 10, 20, 12, 20, 14, 20, 188, 11, 20, 3, 21, 3, 21, 3, 21, 7, 21, 193,
	10, 21, 12, 21, 14, 21, 196, 11, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23,
	3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 2, 2, 26, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	2, 5, 3, 2, 27, 28, 3, 2, 30, 33, 3, 2, 19, 24, 2, 205, 2, 50, 3, 2, 2,
	2, 4, 52, 3, 2, 2, 2, 6, 69, 3, 2, 2, 2, 8, 71, 3, 2, 2, 2, 10, 73, 3,
	2, 2, 2, 12, 87, 3, 2, 2, 2, 14, 94, 3, 2, 2, 2, 16, 108, 3, 2, 2, 2, 18,
	115, 3, 2, 2, 2, 20, 125, 3, 2, 2, 2, 22, 132, 3, 2, 2, 2, 24, 140, 3,
	2, 2, 2, 26, 155, 3, 2, 2, 2, 28, 157, 3, 2, 2, 2, 30, 159, 3, 2, 2, 2,
	32, 167, 3, 2, 2, 2, 34, 169, 3, 2, 2, 2, 36, 177, 3, 2, 2, 2, 38, 181,
	3, 2, 2, 2, 40, 189, 3, 2, 2, 2, 42, 197, 3, 2, 2, 2, 44, 201, 3, 2, 2,
	2, 46, 203, 3, 2, 2, 2, 48, 205, 3, 2, 2, 2, 50, 51, 5, 4, 3, 2, 51, 3,
	3, 2, 2, 2, 52, 58, 5, 6, 4, 2, 53, 54, 5, 8, 5, 2, 54, 55, 5, 6, 4, 2,
	55, 57, 3, 2, 2, 2, 56, 53, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3,
	2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 5, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61,
	70, 5, 10, 6, 2, 62, 70, 5, 12, 7, 2, 63, 70, 5, 14, 8, 2, 64, 70, 5, 16,
	9, 2, 65, 70, 5, 18, 10, 2, 66, 70, 5, 20, 11, 2, 67, 70, 5, 22, 12, 2,
	68, 70, 5, 24, 13, 2, 69, 61, 3, 2, 2, 2, 69, 62, 3, 2, 2, 2, 69, 63, 3,
	2, 2, 2, 69, 64, 3, 2, 2, 2, 69, 65, 3, 2, 2, 2, 69, 66, 3, 2, 2, 2, 69,
	67, 3, 2, 2, 2, 69, 68, 3, 2, 2, 2, 70, 7, 3, 2, 2, 2, 71, 72, 9, 2, 2,
	2, 72, 9, 3, 2, 2, 2, 73, 74, 7, 6, 2, 2, 74, 75, 7, 13, 2, 2, 75, 80,
	5, 26, 14, 2, 76, 77, 7, 3, 2, 2, 77, 78, 5, 34, 18, 2, 78, 79, 7, 4, 2,
	2, 79, 81, 3, 2, 2, 2, 80, 76, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 82,
	3, 2, 2, 2, 82, 83, 7, 15, 2, 2, 83, 84, 7, 3, 2, 2, 84, 85, 5, 30, 16,
	2, 85, 86, 7, 4, 2, 2, 86, 11, 3, 2, 2, 2, 87, 88, 7, 8, 2, 2, 88, 89,
	7, 3, 2, 2, 89, 90, 5, 30, 16, 2, 90, 91, 7, 4, 2, 2, 91, 92, 7, 13, 2,
	2, 92, 93, 5, 26, 14, 2, 93, 13, 3, 2, 2, 2, 94, 100, 7, 5, 2, 2, 95, 101,
	7, 17, 2, 2, 96, 97, 7, 3, 2, 2, 97, 98, 5, 34, 18, 2, 98, 99, 7, 4, 2,
	2, 99, 101, 3, 2, 2, 2, 100, 95, 3, 2, 2, 2, 100, 96, 3, 2, 2, 2, 101,
	102, 3, 2, 2, 2, 102, 103, 7, 14, 2, 2, 103, 106, 5, 26, 14, 2, 104, 105,
	7, 16, 2, 2, 105, 107, 5, 40, 21, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3,
	2, 2, 2, 107, 15, 3, 2, 2, 2, 108, 109, 7, 11, 2, 2, 109, 110, 7, 14, 2,
	2, 110, 113, 5, 26, 14, 2, 111, 112, 7, 16, 2, 2, 112, 114, 5, 40, 21,
	2, 113, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 17, 3, 2, 2, 2, 115,
	116, 7, 9, 2, 2, 116, 119, 5, 26, 14, 2, 117, 118, 7, 8, 2, 2, 118, 120,
	5, 38, 20, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 123, 3,
	2, 2, 2, 121, 122, 7, 16, 2, 2, 122, 124, 5, 40, 21, 2, 123, 121, 3, 2,
	2, 2, 123, 124, 3, 2, 2, 2, 124, 19, 3, 2, 2, 2, 125, 126, 7, 7, 2, 2,
	126, 127, 7, 14, 2, 2, 127, 130, 5, 26, 14, 2, 128, 129, 7, 16, 2, 2, 129,
	131, 5, 40, 21, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 21,
	3, 2, 2, 2, 132, 133, 7, 10, 2, 2, 133, 134, 5, 32, 17, 2, 134, 135, 7,
	14, 2, 2, 135, 138, 5, 26, 14, 2, 136, 137, 7, 16, 2, 2, 137, 139, 5, 40,
	21, 2, 138, 136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 23, 3, 2, 2, 2,
	140, 141, 7, 12, 2, 2, 141, 146, 5, 26, 14, 2, 142, 143, 7, 3, 2, 2, 143,
	144, 5, 34, 18, 2, 144, 145, 7, 4, 2, 2, 145, 147, 3, 2, 2, 2, 146, 142,
	3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 153, 3, 2, 2, 2, 148, 149, 7, 15,
	2, 2, 149, 150, 7, 3, 2, 2, 150, 151, 5, 30, 16, 2, 151, 152, 7, 4, 2,
	2, 152, 154, 3, 2, 2, 2, 153, 148, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154,
	25, 3, 2, 2, 2, 155, 156, 7, 30, 2, 2, 156, 27, 3, 2, 2, 2, 157, 158, 9,
	3, 2, 2, 158, 29, 3, 2, 2, 2, 159, 164, 5, 28, 15, 2, 160, 161, 7, 27,
	2, 2, 161, 163, 5, 28, 15, 2, 162, 160, 3, 2, 2, 2, 163, 166, 3, 2, 2,
	2, 164, 162, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 31, 3, 2, 2, 2, 166,
	164, 3, 2, 2, 2, 167, 168, 7, 30, 2, 2, 168, 33, 3, 2, 2, 2, 169, 174,
	5, 32, 17, 2, 170, 171, 7, 27, 2, 2, 171, 173, 5, 32, 17, 2, 172, 170,
	3, 2, 2, 2, 173, 176, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 174, 175, 3, 2,
	2, 2, 175, 35, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 177, 178, 5, 32, 17, 2,
	178, 179, 7, 18, 2, 2, 179, 180, 5, 28, 15, 2, 180, 37, 3, 2, 2, 2, 181,
	186, 5, 36, 19, 2, 182, 183, 7, 27, 2, 2, 183, 185, 5, 36, 19, 2, 184,
	182, 3, 2, 2, 2, 185, 188, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 186, 187,
	3, 2, 2, 2, 187, 39, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 189, 194, 5, 42,
	22, 2, 190, 191, 7, 25, 2, 2, 191, 193, 5, 42, 22, 2, 192, 190, 3, 2, 2,
	2, 193, 196, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195,
	41, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 197, 198, 5, 46, 24, 2, 198, 199,
	5, 44, 23, 2, 199, 200, 5, 48, 25, 2, 200, 43, 3, 2, 2, 2, 201, 202, 9,
	4, 2, 2, 202, 45, 3, 2, 2, 2, 203, 204, 7, 30, 2, 2, 204, 47, 3, 2, 2,
	2, 205, 206, 9, 3, 2, 2, 206, 49, 3, 2, 2, 2, 18, 58, 69, 80, 100, 106,
	113, 119, 123, 130, 138, 146, 153, 164, 174, 186, 194,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "", "", "", "", "", "", "", "", "", "", "", "", "'*'",
	"'='", "'=='", "'<'", "'<='", "'>'", "'>='", "'!='", "", "", "','", "';'",
}
var symbolicNames = []string{
	"", "", "", "SELECT", "INSERT", "DELETE", "SET", "UPDATE", "ANALYZE", "EXPORT",
	"EXECUTE", "INTO", "FROM", "VALUES", "WHERE", "ASTERISK", "SINGLE_EQ",
	"DOUBLE_EQ", "OP_LT", "LE", "GT", "GE", "NOTEQ", "AND", "OR", "COMMA",
	"SEMICOLON", "WS", "IDENTIFIER", "STRING", "NUMBER", "REAL",
}

var ruleNames = []string{
	"fql", "queryList", "query", "querySep", "insertQuery", "setQuery", "selectQuery",
	"exportQuery", "updateQuery", "deleteQuery", "analyzeQuery", "executeQuery",
	"target", "value", "values", "column", "columns", "columnset", "columnsets",
	"conditions", "condition", "operator", "leftOperand", "rightOperand",
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
	FQLParserUPDATE     = 7
	FQLParserANALYZE    = 8
	FQLParserEXPORT     = 9
	FQLParserEXECUTE    = 10
	FQLParserINTO       = 11
	FQLParserFROM       = 12
	FQLParserVALUES     = 13
	FQLParserWHERE      = 14
	FQLParserASTERISK   = 15
	FQLParserSINGLE_EQ  = 16
	FQLParserDOUBLE_EQ  = 17
	FQLParserOP_LT      = 18
	FQLParserLE         = 19
	FQLParserGT         = 20
	FQLParserGE         = 21
	FQLParserNOTEQ      = 22
	FQLParserAND        = 23
	FQLParserOR         = 24
	FQLParserCOMMA      = 25
	FQLParserSEMICOLON  = 26
	FQLParserWS         = 27
	FQLParserIDENTIFIER = 28
	FQLParserSTRING     = 29
	FQLParserNUMBER     = 30
	FQLParserREAL       = 31
)

// FQLParser rules.
const (
	FQLParserRULE_fql          = 0
	FQLParserRULE_queryList    = 1
	FQLParserRULE_query        = 2
	FQLParserRULE_querySep     = 3
	FQLParserRULE_insertQuery  = 4
	FQLParserRULE_setQuery     = 5
	FQLParserRULE_selectQuery  = 6
	FQLParserRULE_exportQuery  = 7
	FQLParserRULE_updateQuery  = 8
	FQLParserRULE_deleteQuery  = 9
	FQLParserRULE_analyzeQuery = 10
	FQLParserRULE_executeQuery = 11
	FQLParserRULE_target       = 12
	FQLParserRULE_value        = 13
	FQLParserRULE_values       = 14
	FQLParserRULE_column       = 15
	FQLParserRULE_columns      = 16
	FQLParserRULE_columnset    = 17
	FQLParserRULE_columnsets   = 18
	FQLParserRULE_conditions   = 19
	FQLParserRULE_condition    = 20
	FQLParserRULE_operator     = 21
	FQLParserRULE_leftOperand  = 22
	FQLParserRULE_rightOperand = 23
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

func (s *FqlContext) QueryList() IQueryListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryListContext)
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
		p.SetState(48)
		p.QueryList()
	}

	return localctx
}

// IQueryListContext is an interface to support dynamic dispatch.
type IQueryListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryListContext differentiates from other interfaces.
	IsQueryListContext()
}

type QueryListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryListContext() *QueryListContext {
	var p = new(QueryListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_queryList
	return p
}

func (*QueryListContext) IsQueryListContext() {}

func NewQueryListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryListContext {
	var p = new(QueryListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_queryList

	return p
}

func (s *QueryListContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryListContext) AllQuery() []IQueryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQueryContext)(nil)).Elem())
	var tst = make([]IQueryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQueryContext)
		}
	}

	return tst
}

func (s *QueryListContext) Query(i int) IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *QueryListContext) AllQuerySep() []IQuerySepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuerySepContext)(nil)).Elem())
	var tst = make([]IQuerySepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuerySepContext)
		}
	}

	return tst
}

func (s *QueryListContext) QuerySep(i int) IQuerySepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuerySepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuerySepContext)
}

func (s *QueryListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterQueryList(s)
	}
}

func (s *QueryListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitQueryList(s)
	}
}

func (p *FQLParser) QueryList() (localctx IQueryListContext) {
	localctx = NewQueryListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FQLParserRULE_queryList)
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
		p.Query()
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FQLParserCOMMA || _la == FQLParserSEMICOLON {
		{
			p.SetState(51)
			p.QuerySep()
		}
		{
			p.SetState(52)
			p.Query()
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) InsertQuery() IInsertQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsertQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsertQueryContext)
}

func (s *QueryContext) SetQuery() ISetQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetQueryContext)
}

func (s *QueryContext) SelectQuery() ISelectQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectQueryContext)
}

func (s *QueryContext) ExportQuery() IExportQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportQueryContext)
}

func (s *QueryContext) UpdateQuery() IUpdateQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdateQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdateQueryContext)
}

func (s *QueryContext) DeleteQuery() IDeleteQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeleteQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeleteQueryContext)
}

func (s *QueryContext) AnalyzeQuery() IAnalyzeQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnalyzeQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnalyzeQueryContext)
}

func (s *QueryContext) ExecuteQuery() IExecuteQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecuteQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExecuteQueryContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (p *FQLParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FQLParserRULE_query)

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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FQLParserINSERT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.InsertQuery()
		}

	case FQLParserSET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.SetQuery()
		}

	case FQLParserSELECT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.SelectQuery()
		}

	case FQLParserEXPORT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(62)
			p.ExportQuery()
		}

	case FQLParserUPDATE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(63)
			p.UpdateQuery()
		}

	case FQLParserDELETE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(64)
			p.DeleteQuery()
		}

	case FQLParserANALYZE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(65)
			p.AnalyzeQuery()
		}

	case FQLParserEXECUTE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(66)
			p.ExecuteQuery()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQuerySepContext is an interface to support dynamic dispatch.
type IQuerySepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuerySepContext differentiates from other interfaces.
	IsQuerySepContext()
}

type QuerySepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuerySepContext() *QuerySepContext {
	var p = new(QuerySepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_querySep
	return p
}

func (*QuerySepContext) IsQuerySepContext() {}

func NewQuerySepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuerySepContext {
	var p = new(QuerySepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_querySep

	return p
}

func (s *QuerySepContext) GetParser() antlr.Parser { return s.parser }

func (s *QuerySepContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FQLParserSEMICOLON, 0)
}

func (s *QuerySepContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FQLParserCOMMA, 0)
}

func (s *QuerySepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuerySepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuerySepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterQuerySep(s)
	}
}

func (s *QuerySepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitQuerySep(s)
	}
}

func (p *FQLParser) QuerySep() (localctx IQuerySepContext) {
	localctx = NewQuerySepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FQLParserRULE_querySep)
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
		p.SetState(69)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FQLParserCOMMA || _la == FQLParserSEMICOLON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IInsertQueryContext is an interface to support dynamic dispatch.
type IInsertQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsertQueryContext differentiates from other interfaces.
	IsInsertQueryContext()
}

type InsertQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertQueryContext() *InsertQueryContext {
	var p = new(InsertQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_insertQuery
	return p
}

func (*InsertQueryContext) IsInsertQueryContext() {}

func NewInsertQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertQueryContext {
	var p = new(InsertQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_insertQuery

	return p
}

func (s *InsertQueryContext) GetParser() antlr.Parser { return s.parser }

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

func (s *InsertQueryContext) Columns() IColumnsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnsContext)
}

func (s *InsertQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *FQLParser) InsertQuery() (localctx IInsertQueryContext) {
	localctx = NewInsertQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FQLParserRULE_insertQuery)
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
		p.SetState(71)
		p.Match(FQLParserINSERT)
	}
	{
		p.SetState(72)
		p.Match(FQLParserINTO)
	}
	{
		p.SetState(73)
		p.Target()
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FQLParserT__0 {
		{
			p.SetState(74)
			p.Match(FQLParserT__0)
		}
		{
			p.SetState(75)
			p.Columns()
		}
		{
			p.SetState(76)
			p.Match(FQLParserT__1)
		}

	}
	{
		p.SetState(80)
		p.Match(FQLParserVALUES)
	}
	{
		p.SetState(81)
		p.Match(FQLParserT__0)
	}
	{
		p.SetState(82)
		p.Values()
	}
	{
		p.SetState(83)
		p.Match(FQLParserT__1)
	}

	return localctx
}

// ISetQueryContext is an interface to support dynamic dispatch.
type ISetQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetQueryContext differentiates from other interfaces.
	IsSetQueryContext()
}

type SetQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetQueryContext() *SetQueryContext {
	var p = new(SetQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_setQuery
	return p
}

func (*SetQueryContext) IsSetQueryContext() {}

func NewSetQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetQueryContext {
	var p = new(SetQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_setQuery

	return p
}

func (s *SetQueryContext) GetParser() antlr.Parser { return s.parser }

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

func (s *SetQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *FQLParser) SetQuery() (localctx ISetQueryContext) {
	localctx = NewSetQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FQLParserRULE_setQuery)

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
		p.SetState(85)
		p.Match(FQLParserSET)
	}
	{
		p.SetState(86)
		p.Match(FQLParserT__0)
	}
	{
		p.SetState(87)
		p.Values()
	}
	{
		p.SetState(88)
		p.Match(FQLParserT__1)
	}
	{
		p.SetState(89)
		p.Match(FQLParserINTO)
	}
	{
		p.SetState(90)
		p.Target()
	}

	return localctx
}

// ISelectQueryContext is an interface to support dynamic dispatch.
type ISelectQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectQueryContext differentiates from other interfaces.
	IsSelectQueryContext()
}

type SelectQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectQueryContext() *SelectQueryContext {
	var p = new(SelectQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_selectQuery
	return p
}

func (*SelectQueryContext) IsSelectQueryContext() {}

func NewSelectQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectQueryContext {
	var p = new(SelectQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_selectQuery

	return p
}

func (s *SelectQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectQueryContext) SELECT() antlr.TerminalNode {
	return s.GetToken(FQLParserSELECT, 0)
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

func (s *SelectQueryContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(FQLParserASTERISK, 0)
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

func (s *SelectQueryContext) Columns() IColumnsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnsContext)
}

func (s *SelectQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *FQLParser) SelectQuery() (localctx ISelectQueryContext) {
	localctx = NewSelectQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FQLParserRULE_selectQuery)
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
		p.SetState(92)
		p.Match(FQLParserSELECT)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FQLParserASTERISK:
		{
			p.SetState(93)
			p.Match(FQLParserASTERISK)
		}

	case FQLParserT__0:
		{
			p.SetState(94)
			p.Match(FQLParserT__0)
		}
		{
			p.SetState(95)
			p.Columns()
		}
		{
			p.SetState(96)
			p.Match(FQLParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(100)
		p.Match(FQLParserFROM)
	}
	{
		p.SetState(101)
		p.Target()
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FQLParserWHERE {
		{
			p.SetState(102)
			p.Match(FQLParserWHERE)
		}
		{
			p.SetState(103)
			p.Conditions()
		}

	}

	return localctx
}

// IExportQueryContext is an interface to support dynamic dispatch.
type IExportQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExportQueryContext differentiates from other interfaces.
	IsExportQueryContext()
}

type ExportQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExportQueryContext() *ExportQueryContext {
	var p = new(ExportQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_exportQuery
	return p
}

func (*ExportQueryContext) IsExportQueryContext() {}

func NewExportQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportQueryContext {
	var p = new(ExportQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_exportQuery

	return p
}

func (s *ExportQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportQueryContext) EXPORT() antlr.TerminalNode {
	return s.GetToken(FQLParserEXPORT, 0)
}

func (s *ExportQueryContext) FROM() antlr.TerminalNode {
	return s.GetToken(FQLParserFROM, 0)
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

func (s *ExportQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *FQLParser) ExportQuery() (localctx IExportQueryContext) {
	localctx = NewExportQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FQLParserRULE_exportQuery)
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
		p.SetState(106)
		p.Match(FQLParserEXPORT)
	}
	{
		p.SetState(107)
		p.Match(FQLParserFROM)
	}
	{
		p.SetState(108)
		p.Target()
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FQLParserWHERE {
		{
			p.SetState(109)
			p.Match(FQLParserWHERE)
		}
		{
			p.SetState(110)
			p.Conditions()
		}

	}

	return localctx
}

// IUpdateQueryContext is an interface to support dynamic dispatch.
type IUpdateQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdateQueryContext differentiates from other interfaces.
	IsUpdateQueryContext()
}

type UpdateQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateQueryContext() *UpdateQueryContext {
	var p = new(UpdateQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_updateQuery
	return p
}

func (*UpdateQueryContext) IsUpdateQueryContext() {}

func NewUpdateQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateQueryContext {
	var p = new(UpdateQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_updateQuery

	return p
}

func (s *UpdateQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateQueryContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(FQLParserUPDATE, 0)
}

func (s *UpdateQueryContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *UpdateQueryContext) SET() antlr.TerminalNode {
	return s.GetToken(FQLParserSET, 0)
}

func (s *UpdateQueryContext) Columnsets() IColumnsetsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnsetsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnsetsContext)
}

func (s *UpdateQueryContext) WHERE() antlr.TerminalNode {
	return s.GetToken(FQLParserWHERE, 0)
}

func (s *UpdateQueryContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *UpdateQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterUpdateQuery(s)
	}
}

func (s *UpdateQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitUpdateQuery(s)
	}
}

func (p *FQLParser) UpdateQuery() (localctx IUpdateQueryContext) {
	localctx = NewUpdateQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FQLParserRULE_updateQuery)
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
		p.SetState(113)
		p.Match(FQLParserUPDATE)
	}
	{
		p.SetState(114)
		p.Target()
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FQLParserSET {
		{
			p.SetState(115)
			p.Match(FQLParserSET)
		}
		{
			p.SetState(116)
			p.Columnsets()
		}

	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FQLParserWHERE {
		{
			p.SetState(119)
			p.Match(FQLParserWHERE)
		}
		{
			p.SetState(120)
			p.Conditions()
		}

	}

	return localctx
}

// IDeleteQueryContext is an interface to support dynamic dispatch.
type IDeleteQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeleteQueryContext differentiates from other interfaces.
	IsDeleteQueryContext()
}

type DeleteQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteQueryContext() *DeleteQueryContext {
	var p = new(DeleteQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_deleteQuery
	return p
}

func (*DeleteQueryContext) IsDeleteQueryContext() {}

func NewDeleteQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteQueryContext {
	var p = new(DeleteQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_deleteQuery

	return p
}

func (s *DeleteQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteQueryContext) DELETE() antlr.TerminalNode {
	return s.GetToken(FQLParserDELETE, 0)
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

func (s *DeleteQueryContext) WHERE() antlr.TerminalNode {
	return s.GetToken(FQLParserWHERE, 0)
}

func (s *DeleteQueryContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *DeleteQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *FQLParser) DeleteQuery() (localctx IDeleteQueryContext) {
	localctx = NewDeleteQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FQLParserRULE_deleteQuery)
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
		p.SetState(123)
		p.Match(FQLParserDELETE)
	}
	{
		p.SetState(124)
		p.Match(FQLParserFROM)
	}
	{
		p.SetState(125)
		p.Target()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FQLParserWHERE {
		{
			p.SetState(126)
			p.Match(FQLParserWHERE)
		}
		{
			p.SetState(127)
			p.Conditions()
		}

	}

	return localctx
}

// IAnalyzeQueryContext is an interface to support dynamic dispatch.
type IAnalyzeQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnalyzeQueryContext differentiates from other interfaces.
	IsAnalyzeQueryContext()
}

type AnalyzeQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnalyzeQueryContext() *AnalyzeQueryContext {
	var p = new(AnalyzeQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_analyzeQuery
	return p
}

func (*AnalyzeQueryContext) IsAnalyzeQueryContext() {}

func NewAnalyzeQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnalyzeQueryContext {
	var p = new(AnalyzeQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_analyzeQuery

	return p
}

func (s *AnalyzeQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *AnalyzeQueryContext) ANALYZE() antlr.TerminalNode {
	return s.GetToken(FQLParserANALYZE, 0)
}

func (s *AnalyzeQueryContext) Column() IColumnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnContext)
}

func (s *AnalyzeQueryContext) FROM() antlr.TerminalNode {
	return s.GetToken(FQLParserFROM, 0)
}

func (s *AnalyzeQueryContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *AnalyzeQueryContext) WHERE() antlr.TerminalNode {
	return s.GetToken(FQLParserWHERE, 0)
}

func (s *AnalyzeQueryContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *AnalyzeQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnalyzeQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnalyzeQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterAnalyzeQuery(s)
	}
}

func (s *AnalyzeQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitAnalyzeQuery(s)
	}
}

func (p *FQLParser) AnalyzeQuery() (localctx IAnalyzeQueryContext) {
	localctx = NewAnalyzeQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FQLParserRULE_analyzeQuery)
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
		p.SetState(130)
		p.Match(FQLParserANALYZE)
	}
	{
		p.SetState(131)
		p.Column()
	}
	{
		p.SetState(132)
		p.Match(FQLParserFROM)
	}
	{
		p.SetState(133)
		p.Target()
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FQLParserWHERE {
		{
			p.SetState(134)
			p.Match(FQLParserWHERE)
		}
		{
			p.SetState(135)
			p.Conditions()
		}

	}

	return localctx
}

// IExecuteQueryContext is an interface to support dynamic dispatch.
type IExecuteQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExecuteQueryContext differentiates from other interfaces.
	IsExecuteQueryContext()
}

type ExecuteQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecuteQueryContext() *ExecuteQueryContext {
	var p = new(ExecuteQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_executeQuery
	return p
}

func (*ExecuteQueryContext) IsExecuteQueryContext() {}

func NewExecuteQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecuteQueryContext {
	var p = new(ExecuteQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_executeQuery

	return p
}

func (s *ExecuteQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecuteQueryContext) EXECUTE() antlr.TerminalNode {
	return s.GetToken(FQLParserEXECUTE, 0)
}

func (s *ExecuteQueryContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *ExecuteQueryContext) Columns() IColumnsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnsContext)
}

func (s *ExecuteQueryContext) VALUES() antlr.TerminalNode {
	return s.GetToken(FQLParserVALUES, 0)
}

func (s *ExecuteQueryContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *ExecuteQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecuteQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecuteQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterExecuteQuery(s)
	}
}

func (s *ExecuteQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitExecuteQuery(s)
	}
}

func (p *FQLParser) ExecuteQuery() (localctx IExecuteQueryContext) {
	localctx = NewExecuteQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FQLParserRULE_executeQuery)
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
		p.SetState(138)
		p.Match(FQLParserEXECUTE)
	}
	{
		p.SetState(139)
		p.Target()
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FQLParserT__0 {
		{
			p.SetState(140)
			p.Match(FQLParserT__0)
		}
		{
			p.SetState(141)
			p.Columns()
		}
		{
			p.SetState(142)
			p.Match(FQLParserT__1)
		}

	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FQLParserVALUES {
		{
			p.SetState(146)
			p.Match(FQLParserVALUES)
		}
		{
			p.SetState(147)
			p.Match(FQLParserT__0)
		}
		{
			p.SetState(148)
			p.Values()
		}
		{
			p.SetState(149)
			p.Match(FQLParserT__1)
		}

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
	p.EnterRule(localctx, 24, FQLParserRULE_target)

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
		p.SetState(153)
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
	p.EnterRule(localctx, 26, FQLParserRULE_value)
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
		p.SetState(155)
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

func (s *ValuesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FQLParserCOMMA)
}

func (s *ValuesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FQLParserCOMMA, i)
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
	p.EnterRule(localctx, 28, FQLParserRULE_values)
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
		p.SetState(157)
		p.Value()
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FQLParserCOMMA {
		{
			p.SetState(158)
			p.Match(FQLParserCOMMA)
		}
		{
			p.SetState(159)
			p.Value()
		}

		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IColumnContext is an interface to support dynamic dispatch.
type IColumnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnContext differentiates from other interfaces.
	IsColumnContext()
}

type ColumnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnContext() *ColumnContext {
	var p = new(ColumnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_column
	return p
}

func (*ColumnContext) IsColumnContext() {}

func NewColumnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnContext {
	var p = new(ColumnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_column

	return p
}

func (s *ColumnContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FQLParserIDENTIFIER, 0)
}

func (s *ColumnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterColumn(s)
	}
}

func (s *ColumnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitColumn(s)
	}
}

func (p *FQLParser) Column() (localctx IColumnContext) {
	localctx = NewColumnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FQLParserRULE_column)

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
		p.SetState(165)
		p.Match(FQLParserIDENTIFIER)
	}

	return localctx
}

// IColumnsContext is an interface to support dynamic dispatch.
type IColumnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnsContext differentiates from other interfaces.
	IsColumnsContext()
}

type ColumnsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnsContext() *ColumnsContext {
	var p = new(ColumnsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_columns
	return p
}

func (*ColumnsContext) IsColumnsContext() {}

func NewColumnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnsContext {
	var p = new(ColumnsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_columns

	return p
}

func (s *ColumnsContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnsContext) AllColumn() []IColumnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumnContext)(nil)).Elem())
	var tst = make([]IColumnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumnContext)
		}
	}

	return tst
}

func (s *ColumnsContext) Column(i int) IColumnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumnContext)
}

func (s *ColumnsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FQLParserCOMMA)
}

func (s *ColumnsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FQLParserCOMMA, i)
}

func (s *ColumnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterColumns(s)
	}
}

func (s *ColumnsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitColumns(s)
	}
}

func (p *FQLParser) Columns() (localctx IColumnsContext) {
	localctx = NewColumnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FQLParserRULE_columns)
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
		p.SetState(167)
		p.Column()
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FQLParserCOMMA {
		{
			p.SetState(168)
			p.Match(FQLParserCOMMA)
		}
		{
			p.SetState(169)
			p.Column()
		}

		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IColumnsetContext is an interface to support dynamic dispatch.
type IColumnsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnsetContext differentiates from other interfaces.
	IsColumnsetContext()
}

type ColumnsetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnsetContext() *ColumnsetContext {
	var p = new(ColumnsetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_columnset
	return p
}

func (*ColumnsetContext) IsColumnsetContext() {}

func NewColumnsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnsetContext {
	var p = new(ColumnsetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_columnset

	return p
}

func (s *ColumnsetContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnsetContext) Column() IColumnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnContext)
}

func (s *ColumnsetContext) SINGLE_EQ() antlr.TerminalNode {
	return s.GetToken(FQLParserSINGLE_EQ, 0)
}

func (s *ColumnsetContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ColumnsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterColumnset(s)
	}
}

func (s *ColumnsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitColumnset(s)
	}
}

func (p *FQLParser) Columnset() (localctx IColumnsetContext) {
	localctx = NewColumnsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FQLParserRULE_columnset)

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
		p.SetState(175)
		p.Column()
	}
	{
		p.SetState(176)
		p.Match(FQLParserSINGLE_EQ)
	}
	{
		p.SetState(177)
		p.Value()
	}

	return localctx
}

// IColumnsetsContext is an interface to support dynamic dispatch.
type IColumnsetsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnsetsContext differentiates from other interfaces.
	IsColumnsetsContext()
}

type ColumnsetsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnsetsContext() *ColumnsetsContext {
	var p = new(ColumnsetsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FQLParserRULE_columnsets
	return p
}

func (*ColumnsetsContext) IsColumnsetsContext() {}

func NewColumnsetsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnsetsContext {
	var p = new(ColumnsetsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FQLParserRULE_columnsets

	return p
}

func (s *ColumnsetsContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnsetsContext) AllColumnset() []IColumnsetContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumnsetContext)(nil)).Elem())
	var tst = make([]IColumnsetContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumnsetContext)
		}
	}

	return tst
}

func (s *ColumnsetsContext) Columnset(i int) IColumnsetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnsetContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumnsetContext)
}

func (s *ColumnsetsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FQLParserCOMMA)
}

func (s *ColumnsetsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FQLParserCOMMA, i)
}

func (s *ColumnsetsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnsetsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnsetsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.EnterColumnsets(s)
	}
}

func (s *ColumnsetsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FQLListener); ok {
		listenerT.ExitColumnsets(s)
	}
}

func (p *FQLParser) Columnsets() (localctx IColumnsetsContext) {
	localctx = NewColumnsetsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FQLParserRULE_columnsets)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Columnset()
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(180)
				p.Match(FQLParserCOMMA)
			}
			{
				p.SetState(181)
				p.Columnset()
			}

		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 38, FQLParserRULE_conditions)
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
		p.SetState(187)
		p.Condition()
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FQLParserAND {
		{
			p.SetState(188)
			p.Match(FQLParserAND)
		}
		{
			p.SetState(189)
			p.Condition()
		}

		p.SetState(194)
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
	p.EnterRule(localctx, 40, FQLParserRULE_condition)

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
		p.SetState(195)
		p.LeftOperand()
	}
	{
		p.SetState(196)
		p.Operator()
	}
	{
		p.SetState(197)
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
	p.EnterRule(localctx, 42, FQLParserRULE_operator)
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
		p.SetState(199)
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
	p.EnterRule(localctx, 44, FQLParserRULE_leftOperand)

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
		p.SetState(201)
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

func (s *RightOperandContext) STRING() antlr.TerminalNode {
	return s.GetToken(FQLParserSTRING, 0)
}

func (s *RightOperandContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FQLParserNUMBER, 0)
}

func (s *RightOperandContext) REAL() antlr.TerminalNode {
	return s.GetToken(FQLParserREAL, 0)
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
	p.EnterRule(localctx, 46, FQLParserRULE_rightOperand)
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
		p.SetState(203)
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
