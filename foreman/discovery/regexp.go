// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"regexp"
	"strings"
)

const (
	// See : http://graphite.readthedocs.io/en/latest/render_api.html
	exprAsteriskGraphite = "*"
	exprAsteriskGo       = "(.*)"
)

// Regexp represents a regexp for the finder.
type Regexp struct {
	expr     string
	goRegexp *regexp.Regexp
}

// NewRegexp returns a new regexp.
func NewRegexp() *Regexp {
	regexp := &Regexp{
		goRegexp: nil,
	}
	return regexp
}

// Compile parses a regular expression
func (re *Regexp) Compile(expr string) error {
	var err error
	re.goRegexp, err = regexp.Compile(expr)
	if err != nil {
		return err
	}
	re.expr = expr
	return nil
}

// CompileGraphite parses a regular expression to Graphite
// See : http://graphite.readthedocs.io/en/latest/render_api.html
func (re *Regexp) CompileGraphite(expr string) error {
	// FIXME : Only replacing the prefix expression string
	if strings.HasPrefix(expr, exprAsteriskGraphite) {
		expr = strings.Replace(expr, exprAsteriskGraphite, exprAsteriskGo, 1)
	}

	return re.Compile(expr)
}

// matchNodeString reports whether the Regexp matches the string.
func (re *Regexp) matchNodeString(nodeStr string) bool {
	if len(nodeStr) <= 0 {
		return false
	}

	// FIXME : Only checking the prefix expression string
	if strings.HasPrefix(re.expr, exprAsteriskGo) {
		return true
	}
	if strings.HasPrefix(re.expr, nodeStr) {
		return true
	}

	return re.goRegexp.MatchString(nodeStr)
}

// MatchNode reports whether the Regexp matches the node.
func (re *Regexp) MatchNode(node Node) bool {
	ok := re.matchNodeString(node.GetName())
	if ok {
		return true
	}

	ok = re.matchNodeString(node.GetAddress())
	if ok {
		return true
	}

	return false
}

// expandNodeString replaces expression to node returns the result;
func (re *Regexp) expandNodeString(nodeStr string) (string, bool) {
	if len(nodeStr) <= 0 {
		return "", false
	}

	// FIXME : Only replacing the prefix expression string
	if strings.HasPrefix(re.expr, exprAsteriskGo) {
		return strings.Replace(re.expr, exprAsteriskGo, nodeStr, 1), true
	}

	return re.expr, true
}

// ExpandNode replaces expression to node returns the result;
func (re *Regexp) ExpandNode(node Node) (string, bool) {
	result, ok := re.expandNodeString(node.GetName())
	if ok {
		return result, true
	}

	result, ok = re.expandNodeString(node.GetAddress())
	if ok {
		return result, true
	}

	return "", false
}
