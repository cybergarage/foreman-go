// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// ANTLRParser represents a FQL parser based on ANTLR.
type ANTLRParser struct {
	Parser
}

// NewParser returns a new parse.
func NewParser() Parser {
	parser := &ANTLRParser{}
	return parser
}

// ParseString parses a specified FQL string.
func (parser *ANTLRParser) ParseString(fqlString string) (Statements, error) {
	input := antlr.NewInputStream(fqlString)
	lexer := NewFQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewFQLParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Fql()
	l := newFQLTreeWalkListener()
	antlr.ParseTreeWalkerDefault.Walk(l, tree)
	return nil, nil
}
