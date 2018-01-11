// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// Operator represents an operator interface in qualities.
type Operator interface {
	// IsSatisfied checks whether ('value1' op 'value2') is valid
	IsSatisfied(value1 interface{}, value2 interface{}) (bool, error)
}
