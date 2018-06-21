// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"testing"

	"github.com/cybergarage/foreman-go/foreman/node"
)

var testRegexTestCases = [][]string{
	// Graphite Wildcards (Only prefix wildcard)
	// See : http://graphite.readthedocs.io/en/latest/render_api.html
	[]string{"node01", "*", "node01"},
	[]string{"node01", "*.metrics01", "node01.metrics01"},
	[]string{"node01", "*.service.metrics01", "node01.service.metrics01"},
	// Graphite Wildcards (With metrics wildcards)
	[]string{"node01", "*.*", "node01.*"},
	[]string{"node01", "*.metrics01.*", "node01.metrics01.*"},
	[]string{"node01", "*.service.metrics01.*", "node01.service.metrics01.*"},
	[]string{"node01", "*.service.*.metrics01.*", "node01.service.*.metrics01.*"},
	// Graphite Wildcards (No prefix wildcard)
	[]string{"node01", "node01.*", "node01.*"},
	[]string{"node01", "node01.metrics01.*", "node01.metrics01.*"},
	[]string{"node01", "node01.service.metrics01.*", "node01.service.metrics01.*"},
	[]string{"node01", "node01.service.*.metrics01.*", "node01.service.*.metrics01.*"},
}

func TestRegexpGraphite(t *testing.T) {
	for n, testCase := range testRegexTestCases {
		name := testCase[0]
		expr := testCase[1]
		expand := testCase[2]

		re := NewRegexp()
		err := re.CompileGraphite(expr)
		if err != nil {
			t.Errorf("[%d] %s : %s", n, expr, err)
			continue
		}

		node := node.NewBaseNode().(*node.BaseNode)
		node.Name = name

		ok := re.MatchNode(node)
		if !ok {
			t.Errorf("[%d] %s != %s", n, expr, node.GetName())
			continue
		}

		expandedName, ok := re.ExpandNode(node)
		if expand != expandedName {
			t.Errorf("[%d] %s != %s", n, expand, expandedName)
			continue
		}
	}
}
