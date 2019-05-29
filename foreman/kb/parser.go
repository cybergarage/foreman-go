// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

// Parser represents an parser interface for FQL.
type Parser interface {
	ParseString(factory Factory, ruleString string) (Rule, error)
}
