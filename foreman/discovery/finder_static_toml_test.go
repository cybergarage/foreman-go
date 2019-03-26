// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"testing"
)

func TestStaticTOMLFinder(t *testing.T) {
	nodes := setupTestFinderNodes()
	cluster := "test cluster"

	for _, finder := range [2]Finder{
		NewStaticTOMLFinderWithNodes(nodes),
		NewStaticTOMLFinderWithHosts(testFinderNodeNames, cluster),
	} {
		err := finder.Start()
		if err != nil {
			t.Error(err)
			return
		}

		err = finderTest(finder)
		if err != nil {
			t.Error(err)
		}

		err = finder.Stop()
		if err != nil {
			t.Error(err)
		}
	}

}
