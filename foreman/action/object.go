// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// Object represents an abstract object in the package.
type Object interface {
	GetName() string
	String() string
}
