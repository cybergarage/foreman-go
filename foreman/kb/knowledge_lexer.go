// Code generated from Knowledge.g4 by ANTLR 4.7.2. DO NOT EDIT.

package kb

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 259,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 7, 16, 132,
	10, 16, 12, 16, 14, 16, 135, 11, 16, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17,
	141, 10, 17, 12, 17, 14, 17, 144, 11, 17, 3, 17, 3, 17, 3, 18, 5, 18, 149,
	10, 18, 3, 18, 6, 18, 152, 10, 18, 13, 18, 14, 18, 153, 3, 19, 5, 19, 157,
	10, 19, 3, 19, 6, 19, 160, 10, 19, 13, 19, 14, 19, 161, 3, 19, 3, 19, 7,
	19, 166, 10, 19, 12, 19, 14, 19, 169, 11, 19, 3, 19, 5, 19, 172, 10, 19,
	3, 19, 5, 19, 175, 10, 19, 3, 19, 3, 19, 6, 19, 179, 10, 19, 13, 19, 14,
	19, 180, 3, 19, 5, 19, 184, 10, 19, 3, 19, 5, 19, 187, 10, 19, 3, 19, 6,
	19, 190, 10, 19, 13, 19, 14, 19, 191, 3, 19, 5, 19, 195, 10, 19, 3, 20,
	3, 20, 5, 20, 199, 10, 20, 3, 20, 6, 20, 202, 10, 20, 13, 20, 14, 20, 203,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31,
	3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3,
	36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41,
	3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3,
	47, 3, 47, 2, 2, 48, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10,
	19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19,
	37, 20, 39, 2, 41, 2, 43, 2, 45, 2, 47, 2, 49, 2, 51, 2, 53, 2, 55, 2,
	57, 2, 59, 2, 61, 2, 63, 2, 65, 2, 67, 2, 69, 2, 71, 2, 73, 2, 75, 2, 77,
	2, 79, 2, 81, 2, 83, 2, 85, 2, 87, 2, 89, 2, 91, 2, 93, 2, 3, 2, 34, 5,
	2, 11, 12, 15, 15, 34, 34, 4, 2, 67, 92, 99, 124, 8, 2, 44, 44, 47, 48,
	50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 36, 36, 4, 2, 45, 45, 47, 47, 4,
	2, 71, 71, 103, 103, 5, 2, 50, 59, 67, 72, 99, 104, 4, 2, 67, 67, 99, 99,
	4, 2, 68, 68, 100, 100, 4, 2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 102,
	4, 2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 106,
	4, 2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 109,
	4, 2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 112,
	4, 2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 115,
	4, 2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 118,
	4, 2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 121,
	4, 2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 124,
	2, 248, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 3, 95, 3, 2, 2,
	2, 5, 97, 3, 2, 2, 2, 7, 99, 3, 2, 2, 2, 9, 101, 3, 2, 2, 2, 11, 104, 3,
	2, 2, 2, 13, 107, 3, 2, 2, 2, 15, 109, 3, 2, 2, 2, 17, 112, 3, 2, 2, 2,
	19, 114, 3, 2, 2, 2, 21, 117, 3, 2, 2, 2, 23, 119, 3, 2, 2, 2, 25, 121,
	3, 2, 2, 2, 27, 123, 3, 2, 2, 2, 29, 125, 3, 2, 2, 2, 31, 129, 3, 2, 2,
	2, 33, 136, 3, 2, 2, 2, 35, 148, 3, 2, 2, 2, 37, 194, 3, 2, 2, 2, 39, 196,
	3, 2, 2, 2, 41, 205, 3, 2, 2, 2, 43, 207, 3, 2, 2, 2, 45, 209, 3, 2, 2,
	2, 47, 211, 3, 2, 2, 2, 49, 213, 3, 2, 2, 2, 51, 215, 3, 2, 2, 2, 53, 217,
	3, 2, 2, 2, 55, 219, 3, 2, 2, 2, 57, 221, 3, 2, 2, 2, 59, 223, 3, 2, 2,
	2, 61, 225, 3, 2, 2, 2, 63, 227, 3, 2, 2, 2, 65, 229, 3, 2, 2, 2, 67, 231,
	3, 2, 2, 2, 69, 233, 3, 2, 2, 2, 71, 235, 3, 2, 2, 2, 73, 237, 3, 2, 2,
	2, 75, 239, 3, 2, 2, 2, 77, 241, 3, 2, 2, 2, 79, 243, 3, 2, 2, 2, 81, 245,
	3, 2, 2, 2, 83, 247, 3, 2, 2, 2, 85, 249, 3, 2, 2, 2, 87, 251, 3, 2, 2,
	2, 89, 253, 3, 2, 2, 2, 91, 255, 3, 2, 2, 2, 93, 257, 3, 2, 2, 2, 95, 96,
	7, 44, 2, 2, 96, 4, 3, 2, 2, 2, 97, 98, 7, 42, 2, 2, 98, 6, 3, 2, 2, 2,
	99, 100, 7, 43, 2, 2, 100, 8, 3, 2, 2, 2, 101, 102, 7, 63, 2, 2, 102, 103,
	7, 63, 2, 2, 103, 10, 3, 2, 2, 2, 104, 105, 7, 35, 2, 2, 105, 106, 7, 63,
	2, 2, 106, 12, 3, 2, 2, 2, 107, 108, 7, 62, 2, 2, 108, 14, 3, 2, 2, 2,
	109, 110, 7, 62, 2, 2, 110, 111, 7, 63, 2, 2, 111, 16, 3, 2, 2, 2, 112,
	113, 7, 64, 2, 2, 113, 18, 3, 2, 2, 2, 114, 115, 7, 64, 2, 2, 115, 116,
	7, 63, 2, 2, 116, 20, 3, 2, 2, 2, 117, 118, 7, 40, 2, 2, 118, 22, 3, 2,
	2, 2, 119, 120, 7, 126, 2, 2, 120, 24, 3, 2, 2, 2, 121, 122, 7, 46, 2,
	2, 122, 26, 3, 2, 2, 2, 123, 124, 7, 61, 2, 2, 124, 28, 3, 2, 2, 2, 125,
	126, 9, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 8, 15, 2, 2, 128, 30,
	3, 2, 2, 2, 129, 133, 9, 3, 2, 2, 130, 132, 9, 4, 2, 2, 131, 130, 3, 2,
	2, 2, 132, 135, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2,
	134, 32, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 136, 142, 7, 36, 2, 2, 137,
	138, 7, 36, 2, 2, 138, 141, 7, 36, 2, 2, 139, 141, 10, 5, 2, 2, 140, 137,
	3, 2, 2, 2, 140, 139, 3, 2, 2, 2, 141, 144, 3, 2, 2, 2, 142, 140, 3, 2,
	2, 2, 142, 143, 3, 2, 2, 2, 143, 145, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2,
	145, 146, 7, 36, 2, 2, 146, 34, 3, 2, 2, 2, 147, 149, 9, 6, 2, 2, 148,
	147, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 151, 3, 2, 2, 2, 150, 152,
	4, 50, 59, 2, 151, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 151, 3,
	2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 36, 3, 2, 2, 2, 155, 157, 9, 6, 2,
	2, 156, 155, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 159, 3, 2, 2, 2, 158,
	160, 4, 50, 59, 2, 159, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 159,
	3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 167, 7, 48,
	2, 2, 164, 166, 4, 50, 59, 2, 165, 164, 3, 2, 2, 2, 166, 169, 3, 2, 2,
	2, 167, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 171, 3, 2, 2, 2, 169,
	167, 3, 2, 2, 2, 170, 172, 5, 39, 20, 2, 171, 170, 3, 2, 2, 2, 171, 172,
	3, 2, 2, 2, 172, 195, 3, 2, 2, 2, 173, 175, 9, 6, 2, 2, 174, 173, 3, 2,
	2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 178, 7, 48, 2, 2,
	177, 179, 4, 50, 59, 2, 178, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180,
	178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 183, 3, 2, 2, 2, 182, 184,
	5, 39, 20, 2, 183, 182, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 195, 3,
	2, 2, 2, 185, 187, 9, 6, 2, 2, 186, 185, 3, 2, 2, 2, 186, 187, 3, 2, 2,
	2, 187, 189, 3, 2, 2, 2, 188, 190, 4, 50, 59, 2, 189, 188, 3, 2, 2, 2,
	190, 191, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192,
	193, 3, 2, 2, 2, 193, 195, 5, 39, 20, 2, 194, 156, 3, 2, 2, 2, 194, 174,
	3, 2, 2, 2, 194, 186, 3, 2, 2, 2, 195, 38, 3, 2, 2, 2, 196, 198, 9, 7,
	2, 2, 197, 199, 9, 6, 2, 2, 198, 197, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2,
	199, 201, 3, 2, 2, 2, 200, 202, 4, 50, 59, 2, 201, 200, 3, 2, 2, 2, 202,
	203, 3, 2, 2, 2, 203, 201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 40, 3,
	2, 2, 2, 205, 206, 9, 8, 2, 2, 206, 42, 3, 2, 2, 2, 207, 208, 9, 9, 2,
	2, 208, 44, 3, 2, 2, 2, 209, 210, 9, 10, 2, 2, 210, 46, 3, 2, 2, 2, 211,
	212, 9, 11, 2, 2, 212, 48, 3, 2, 2, 2, 213, 214, 9, 12, 2, 2, 214, 50,
	3, 2, 2, 2, 215, 216, 9, 7, 2, 2, 216, 52, 3, 2, 2, 2, 217, 218, 9, 13,
	2, 2, 218, 54, 3, 2, 2, 2, 219, 220, 9, 14, 2, 2, 220, 56, 3, 2, 2, 2,
	221, 222, 9, 15, 2, 2, 222, 58, 3, 2, 2, 2, 223, 224, 9, 16, 2, 2, 224,
	60, 3, 2, 2, 2, 225, 226, 9, 17, 2, 2, 226, 62, 3, 2, 2, 2, 227, 228, 9,
	18, 2, 2, 228, 64, 3, 2, 2, 2, 229, 230, 9, 19, 2, 2, 230, 66, 3, 2, 2,
	2, 231, 232, 9, 20, 2, 2, 232, 68, 3, 2, 2, 2, 233, 234, 9, 21, 2, 2, 234,
	70, 3, 2, 2, 2, 235, 236, 9, 22, 2, 2, 236, 72, 3, 2, 2, 2, 237, 238, 9,
	23, 2, 2, 238, 74, 3, 2, 2, 2, 239, 240, 9, 24, 2, 2, 240, 76, 3, 2, 2,
	2, 241, 242, 9, 25, 2, 2, 242, 78, 3, 2, 2, 2, 243, 244, 9, 26, 2, 2, 244,
	80, 3, 2, 2, 2, 245, 246, 9, 27, 2, 2, 246, 82, 3, 2, 2, 2, 247, 248, 9,
	28, 2, 2, 248, 84, 3, 2, 2, 2, 249, 250, 9, 29, 2, 2, 250, 86, 3, 2, 2,
	2, 251, 252, 9, 30, 2, 2, 252, 88, 3, 2, 2, 2, 253, 254, 9, 31, 2, 2, 254,
	90, 3, 2, 2, 2, 255, 256, 9, 32, 2, 2, 256, 92, 3, 2, 2, 2, 257, 258, 9,
	33, 2, 2, 258, 94, 3, 2, 2, 2, 20, 2, 133, 140, 142, 148, 153, 156, 161,
	167, 171, 174, 180, 183, 186, 191, 194, 198, 203, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'*'", "'('", "')'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='",
	"'&'", "'|'", "','", "';'",
}

var lexerSymbolicNames = []string{
	"", "ASTERISK", "BS", "BE", "DEQ", "NEQ", "LT", "LE", "GT", "GE", "AND",
	"OR", "COMMA", "SEMICOLON", "WS", "IDENTIFIER", "STRING", "NUMBER", "REAL",
}

var lexerRuleNames = []string{
	"ASTERISK", "BS", "BE", "DEQ", "NEQ", "LT", "LE", "GT", "GE", "AND", "OR",
	"COMMA", "SEMICOLON", "WS", "IDENTIFIER", "STRING", "NUMBER", "REAL", "EXPONENT",
	"HEX_DIGIT", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
	"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

type KnowledgeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewKnowledgeLexer(input antlr.CharStream) *KnowledgeLexer {

	l := new(KnowledgeLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Knowledge.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// KnowledgeLexer tokens.
const (
	KnowledgeLexerASTERISK   = 1
	KnowledgeLexerBS         = 2
	KnowledgeLexerBE         = 3
	KnowledgeLexerDEQ        = 4
	KnowledgeLexerNEQ        = 5
	KnowledgeLexerLT         = 6
	KnowledgeLexerLE         = 7
	KnowledgeLexerGT         = 8
	KnowledgeLexerGE         = 9
	KnowledgeLexerAND        = 10
	KnowledgeLexerOR         = 11
	KnowledgeLexerCOMMA      = 12
	KnowledgeLexerSEMICOLON  = 13
	KnowledgeLexerWS         = 14
	KnowledgeLexerIDENTIFIER = 15
	KnowledgeLexerSTRING     = 16
	KnowledgeLexerNUMBER     = 17
	KnowledgeLexerREAL       = 18
)