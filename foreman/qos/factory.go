// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

// Factory represents an abstract interface for the parser.
type Factory interface {
	CreateVariable(id string) (Variable, error)
}
