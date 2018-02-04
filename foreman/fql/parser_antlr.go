// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// antlrParser represents a FQL parser based on ANTLR.
type antlrParser struct {
	Parser
}

// NewParser returns a new parse.
func NewParser() Parser {
	parser := &antlrParser{}
	return parser
}

// ParseString parses a specified FQL string.
func (parser *antlrParser) ParseString(queryString string) (Queries, error) {
	if len(queryString) <= 0 {
		return nil, fmt.Errorf(errorParserEmptyQuery)
	}

	input := antlr.NewInputStream(queryString)
	lexer := NewFQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewFQLParser(stream)
	el := newANTLRParserErrorListener()
	p.AddErrorListener(el)
	p.BuildParseTrees = true
	tree := p.Fql()
	pl := newANTLRParserListener()
	antlr.ParseTreeWalkerDefault.Walk(pl, tree)
	if !el.IsSuccess() {
		return nil, fmt.Errorf("%s (%s)", queryString, el.GetError().Error())
	}

	return pl.Queries, nil
}
