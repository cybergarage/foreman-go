// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"github.com/cybergarage/foreman-go/foreman/discovery/echonet"
)

const (
	errorEchonetFinderIsNotRunning = "Finder is not running"
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
	log.SetSharedLogger(log.NewStdoutLogger(log.LoggerLevelTrace))

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
	}

	finder := NewEchonetFinder()

	err = finder.Start()
	if err != nil {
		t.Error(err)
		return
	}

	err = finder.Search()
	if err != nil {
		t.Error(err)
		return
	}

	finderTest(t, finder)

	time.Sleep(time.Minute * time.Duration(len(nodes)))

	err = finder.Stop()
	if err != nil {
		t.Error(err)
	}

	for _, node := range nodes {
		err = node.Stop()
		if err != nil {
			t.Error(err)
			return
		}
	}
}
*/
