// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// antlrParser represents a knowledgebase parser based on ANTLR.
type antlrParser struct {
	Parser
}

// NewParser returns a new parse.
func NewParser() Parser {
	parser := &antlrParser{}
	return parser
}

// ParseString parses a specified FQL string.
func (parser *antlrParser) ParseString(factory Factory, ruleString string) (Rule, error) {
	if len(ruleString) <= 0 {
		return nil, fmt.Errorf(errorParserEmptyRule)
	}

	rule := factory.CreateRule()

	input := antlr.NewInputStream(ruleString)
	lexer := NewKnowledgeLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewKnowledgeParser(stream)
	el := newANTLRParserErrorListener()
	p.AddErrorListener(el)
	p.BuildParseTrees = true
	tree := p.Knowledge()
	pl := newANTLRParserListener(rule, factory)
	antlr.ParseTreeWalkerDefault.Walk(pl, tree)
	if !el.IsSuccess() {
		return nil, fmt.Errorf("%s (%s)", ruleString, el.GetError().Error())
	}

	return rule, nil
}
