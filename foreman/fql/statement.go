// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// Statement represents a statement interface.
type Statement interface {
	GetName() string
	GetType() StatementType
	GetParameters() Parameters
	GetParameter(string) (Parameter, bool)
}
