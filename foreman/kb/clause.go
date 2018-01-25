// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// Clause represents a clause interface.
type Clause interface {
	AddFormula(formula Formula) error
	GetFormulas() []Formula
	IsSatisfied() (bool, error)
	String() string
}
