// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"testing"
)

func finderTest(t *testing.T, finder Finder) {
}

func TestNewSharedFinder(t *testing.T) {
	finder := NewSharedFinder()
	finderTest(t, finder)
}

func TestNewStaticFinder(t *testing.T) {
	finder := NewStaticFinderWithNodes(make([]Node, 0))
	finderTest(t, finder)
}
