// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// Object represents an operand object which has a name.
type Object interface {
	Operand
	// GetName returns a operand name.
	GetName() string
}
