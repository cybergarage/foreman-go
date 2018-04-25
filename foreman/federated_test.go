// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"testing"

	"github.com/cybergarage/foreman-go/foreman/discovery"
)

func TestFederatedMetrics(t *testing.T) {
	nodes := make([]discovery.Node, 0)
	finder := discovery.NewStaticFinderWithNodes(nodes)
	finder.SearchAll()

	client := NewClient()
	client.AddFinder(finder)
}
