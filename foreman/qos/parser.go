// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

// Parsering represents an abstract interface for the parser.
type Parsering interface {
	CreateVariable(id string) (Variable, error)
}

// A Parser represents a parser.
type Parser struct {
	Parsering
}

// NewParserWithInterface returns a new parser with the specified interface.
func NewParserWithInterface(parsering Parsering) *Parser {
	parser := &Parser{
		Parsering: parsering,
	}
	return parser
}

// ParseString returns a rule when the specified string is valid, otherwise an error
func (parser *Parser) ParseString(ruleString string) (*Rule, error) {
	rule := NewRule()
	err := rule.ParseString(ruleString)
	if err != nil {
		return nil, err
	}
	return rule, nil
}