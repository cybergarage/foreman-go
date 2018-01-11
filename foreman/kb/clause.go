// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// A Clause represents a clause in a QoS rule.
type Clause struct {
	Formulas []*Formula
}

// NewClause returns a new clause.
func NewClause() *Clause {
	clause := &Clause{
		Formulas: make([]*Formula, 0),
	}
	return clause
}

// AddFormula adds a new formula.
func (clause *Clause) AddFormula(formula *Formula) error {
	clause.Formulas = append(clause.Formulas, formula)
	return nil
}
