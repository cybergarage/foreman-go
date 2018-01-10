// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

// A Parser represents a parser.
type Parser struct {
	Factory
}

// NewParserWithFactory returns a new parser with the specified interface.
func NewParserWithFactory(factory Factory) *Parser {
	parser := &Parser{
		Factory: factory,
	}
	return parser
}

// ParseString returns a rule when the specified string is valid, otherwise an error
func (parser *Parser) ParseString(ruleString string) (*Rule, error) {
	rule := NewRule()
	err := rule.ParseString(parser, ruleString)
	if err != nil {
		return nil, err
	}
	return rule, nil
}
