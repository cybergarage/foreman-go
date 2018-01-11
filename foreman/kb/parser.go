// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// A Parser represents a parser.
type Parser struct {
}

// NewParser returns a new parser.
func NewParser() *Parser {
	parser := &Parser{}
	return parser
}

// ParseString returns a rule when the specified string is valid, otherwise an error
func (parser *Parser) ParseString(factory Factory, ruleString string) (*Rule, error) {
	rule := NewRule()
	err := rule.ParseString(factory, ruleString)
	if err != nil {
		return nil, err
	}
	return rule, nil
}
