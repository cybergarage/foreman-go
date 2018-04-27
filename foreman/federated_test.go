// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"testing"

	"github.com/cybergarage/foreman-go/foreman/discovery"
)

func federatedMetricsTest(t *testing.T, client *Client, nodes []Node) {
}

func setupFederatedNodes(t *testing.T) []discovery.Node {
	nodes := make([]discovery.Node, 0)
	return nodes
}

func TestFederatedMetricsWithStaticFinder(t *testing.T) {
	nodes := setupFederatedNodes(t)
	finder := discovery.NewStaticFinderWithNodes(nodes)

	client := NewClient()
	client.AddFinder(finder)

	//federatedMetricsTest(t, client, nodes)
}
