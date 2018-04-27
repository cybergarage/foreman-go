// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"regexp"
	"testing"
)

const (
	testNodeName      = "foreman001.cybergarage.org"
	testMatchingError = "Matching Error (%s) : %s"
)

var testNodeRegExpStrings = []string{
	testNodeName,
	".*",
}

func nodeMachingTest(t *testing.T, node Node) {
	name := node.GetName()
	for _, reString := range testNodeRegExpStrings {
		re := regexp.MustCompile(reString)
		if !re.MatchString(name) {
			t.Errorf(testMatchingError, reString, name)
		}
	}
}

func TestNewBaseNode(t *testing.T) {
	node := newBaseNode().(*baseNode)
	node.Name = testNodeName
	nodeMachingTest(t, node)
}
