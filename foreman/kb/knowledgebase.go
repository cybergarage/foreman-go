// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kb

//KnowledgeBase includes all knowledge rules.
type KnowledgeBase struct {
	Rules     []Rule
	variables map[string]Variable
}

// NewKnowledgeBase returns a new knowledge base instance.
func NewKnowledgeBase() *KnowledgeBase {
	kb := &KnowledgeBase{
		Rules:     make([]Rule, 0),
		variables: make(map[string]Variable),
	}
	return kb
}
