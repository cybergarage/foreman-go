// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	server := NewServer()

	err := server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	err = server.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestNewServerEquals(t *testing.T) {
	server := NewServer()
	remoteNode := NewRemoteNodeWithDiscoveryNode(server)

	if !NodeEqual(server, remoteNode) {
		t.Errorf("%s != %s", server.GetName(), remoteNode.GetName())
	}
}
