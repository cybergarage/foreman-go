// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

type antlrParserListener struct {
	*BaseknowledgebaseListener
}

func newANTLRParserListener() *antlrParserListener {
	l := &antlrParserListener{
		BaseknowledgebaseListener: &BaseknowledgebaseListener{},
	}
	return l
}
