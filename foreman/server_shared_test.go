// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"testing"
)

const (
	testSharedNodeCont = 3
)

func sharedRegistryTest(t *testing.T, client *Client, nodes []*Server) {
	//testNodeCount := len(nodes)

	for _, node := range nodes {
		client.SetHTTPPort(node.GetHTTPPort())
		client.SetCarbonPort(node.GetCarbonPort())
		client.SetRenderPort(node.GetRenderPort())
	}
}

func TestStandaloneSharedRegistryWithStaticFinder(t *testing.T) {
}

func TestMultiNodeSharedRegistryWithStaticFinder(t *testing.T) {
}
