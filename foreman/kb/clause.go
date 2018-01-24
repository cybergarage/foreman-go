// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import "bytes"

// A Clause represents a clause in a QoS rule.
type Clause struct {
	Formulas []Formula
}

// NewClause returns a new clause.
func NewClause() *Clause {
	clause := &Clause{
		Formulas: make([]Formula, 0),
	}
	return clause
}

// AddFormula adds a new formula.
func (clause *Clause) AddFormula(formula Formula) error {
	clause.Formulas = append(clause.Formulas, formula)
	return nil
}

// String returns a string description of the instance
func (clause *Clause) String() string {
	if len(clause.Formulas) == 0 {
		return ""
	}

	if len(clause.Formulas) == 1 {
		return clause.Formulas[0].String()
	}

	var b bytes.Buffer
	b.Write([]byte(StartBracket))
	for n, formula := range clause.Formulas {
		if 0 < n {
			b.Write([]byte(SpaceSeparator + FormulaSeparator + SpaceSeparator))
		}
		b.Write([]byte(formula.String()))
	}
	b.Write([]byte(EndBracket))
	return b.String()
}
