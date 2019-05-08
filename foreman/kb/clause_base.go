// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

import "bytes"

// BaseClause represents a clause.
type BaseClause struct {
	Clause
	Formulas []Formula
}

// NewClause returns a new clause.
func NewClause() *BaseClause {
	clause := &BaseClause{
		Formulas: make([]Formula, 0),
	}
	return clause
}

// AddFormula adds a new formula.
func (clause *BaseClause) AddFormula(formula Formula) error {
	clause.Formulas = append(clause.Formulas, formula)
	return nil
}

// GetFormulas returns all formula list.
func (clause *BaseClause) GetFormulas() []Formula {
	return clause.Formulas
}

// IsSatisfied returns whether the all formulas in the clause are satisfied.
func (clause *BaseClause) IsSatisfied() (bool, error) {
	var lastErr error
	for _, formula := range clause.Formulas {
		ok, err := formula.IsSatisfied()

		// Unknown error is considered to be satisfied
		// ex. the operand variable is not found
		if err != nil {
			lastErr = err
			continue
		}

		if !ok {
			return false, lastErr
		}
	}
	return true, lastErr
}

// String returns a string description of the instance
func (clause *BaseClause) String() string {
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
