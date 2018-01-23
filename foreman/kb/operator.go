// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// Operator represents an operator interface in formulas.
type Operator interface {
	// IsSatisfied checks whether ('leftValue' op 'rightValue') is valid
	IsSatisfied(leftValue interface{}, rightValue interface{}) (bool, error)
	// String returns a string description of the instance
	String() string
}
