// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"testing"
)

func TestSharedFinder(t *testing.T) {
	nodes := setupTestFinderNodes()

	finder := NewSharedFinder().(*SharedFinder)
	for _, node := range nodes {
		finder.addNode(node)
	}

	err := finder.Start()
	if err != nil {
		t.Error(err)
		return
	}

	finderTest(t, finder)

	err = finder.Stop()
	if err != nil {
		t.Error(err)
	}
}
