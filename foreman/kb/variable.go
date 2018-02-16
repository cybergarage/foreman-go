// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// Variable represents an variable interface in formulas.
type Variable interface {
	GetName() string
	SetValue(interface{}) error
	GetValue() (interface{}, error)
}
