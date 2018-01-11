// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"github.com/cybergarage/foreman-go/foreman/kb"
)

// A Parser represents a parser.
type Parser struct {
	*kb.Parser
}

// NewParserWithFactory returns a new parser with the specified interface.
func NewParserWithFactory(factory kb.Factory) *Parser {
	parser := &Parser{
		Parser: kb.NewParserWithFactory(factory),
	}
	return parser
}
