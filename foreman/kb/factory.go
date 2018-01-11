// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// Factory represents an abstract interface for the parser.
type Factory interface {
	// CreateVariable return a singleton value.
	CreateVariable(variable interface{}) (Variable, error)
	// CreateOperator return an operater instance.
	CreateOperator(operator interface{}) (Operator, error)
	// CreateObjective return an objective instance.
	CreateObjective(object interface{}) (Objective, error)
}
