// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"github.com/cybergarage/foreman-go/foreman/discovery/echonet"
)

func setupTestEchonetFinderNodes() ([]*echonet.EchonetNode, error) {
	nodes := setupTestFinderNodes()
	echonetNodes := make([]*echonet.EchonetNode, len(nodes))
	for n, node := range nodes {
		echonetNode, err := echonet.NewEchonetNodeWithNode(node)
		if err != nil {
			return nil, err
		}
		echonetNodes[n] = echonetNode
	}
	return echonetNodes, nil
}

/*
FIXME : Enable TestEchonetFinder
func TestEchonetFinder(t *testing.T) {
	nodes, err := setupTestEchonetFinderNodes()
	if err != nil {
		t.Error(err)
		return
	}

	for _, node := range nodes {
		err = node.Start()
		if err != nil {
			t.Error(err)
			return
		}
		defer node.Stop()
	}

	finder := NewEchonetFinder()

	err = finder.Start()
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
*/
